package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"github.com/samber/lo"
)

import (
	"strings"
)

type User struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	NickName string `json:"nick_name"`
	Mobile   string `json:"mobile"`
}

type LDAPServer struct {
	ServerUrl  string
	BaseDN     string
	BindDN     string
	BindPass   string
	UserFilter string
	UserAttr   string
}

func (l *LDAPServer) Conn() (*ldap.Conn, error) {
	conn, err := ldap.DialURL(l.ServerUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to LDAP server: %v", err)
	}

	if err = conn.Bind(l.BindDN, l.BindPass); err != nil {
		conn.Close()
		return nil, fmt.Errorf("LDAP bind failed: %v", err)
	}
	return conn, nil
}

func (l *LDAPServer) SearchLDAPUsers() ([]User, error) {
	// 提前解析用户属性
	userAttr, err := l.ParseUserAttr()
	if err != nil {
		return nil, err
	}

	conn, err := l.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sr, err := l.search(conn, l.UserFilter, userAttr)
	if err != nil {
		return nil, err
	}

	var users []User
	for _, entry := range sr.Entries {
		users = append(users, User{
			Username: GetAttributeValueOrEmpty(entry.GetAttributeValue(userAttr.Username)),
			Email:    GetAttributeValueOrEmpty(entry.GetAttributeValue(userAttr.Email)),
			NickName: GetAttributeValueOrEmpty(entry.GetAttributeValue(userAttr.NickName)),
			Mobile:   GetAttributeValueOrEmpty(entry.GetAttributeValue(userAttr.Mobile)),
		})
	}

	return users, nil
}

func (l *LDAPServer) VerifyLDAPUser(username, password string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	userDN, err := l.findUserDN(conn, username)
	if err != nil {
		return err
	}

	if err = conn.Bind(userDN, password); err != nil {
		return fmt.Errorf("LDAP authentication failed: %v", err)
	}
	return nil
}

func (l *LDAPServer) search(conn *ldap.Conn, filter string, userAttr User) (*ldap.SearchResult, error) {
	attributes := []string{userAttr.Username, userAttr.NickName, userAttr.Email, userAttr.Mobile}

	searchRequest := ldap.NewSearchRequest(
		l.BaseDN,
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		filter,
		attributes,
		nil,
	)

	return conn.Search(searchRequest)
}

func (l *LDAPServer) findUserDN(conn *ldap.Conn, username string) (string, error) {
	filter := fmt.Sprintf("(uid=%s)", username)
	userAttr, err := l.ParseUserAttr()
	if err != nil {
		return "", err
	}

	sr, err := l.search(conn, filter, userAttr)
	if err != nil {
		return "", err
	}

	if len(sr.Entries) == 0 {
		return "", fmt.Errorf("user not found")
	}
	if len(sr.Entries) > 1 {
		return "", fmt.Errorf("multiple users found")
	}

	return sr.Entries[0].DN, nil
}

// ParseUserAttr 将传递过来的字段映射到ldap字段
// e: {"username":"sn","nick_name":"cn","email":"mail","mobile":"homePhone"}
func (l *LDAPServer) ParseUserAttr() (User, error) {
	var user User
	err := json.Unmarshal([]byte(l.UserAttr), &user)
	if err != nil {
		return User{}, fmt.Errorf("error parsing JSON: %v", err)
	}

	return user, nil
}

func (l *LDAPServer) DeleteLDAPUser(username string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 找到用户的 DN
	userDN, err := l.findUserDN(conn, username)
	if err != nil {
		return err
	}

	// 创建删除请求
	delRequest := ldap.NewDelRequest(userDN, nil)

	// 执行删除
	if err = conn.Del(delRequest); err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	fmt.Printf("Successfully deleted user: %s\n", username)
	return nil
}

// 添加用户到指定的LDAP组
func (l *LDAPServer) AddUserToGroup(username, groupName string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 查找用户DN
	userDN, err := l.findUserDN(conn, username)
	if err != nil {
		return err
	}

	// 构造组DN
	groupDN := fmt.Sprintf("cn=%s,cn=group,%s", groupName, l.SplitDn(l.BindDN))

	// 修改组添加成员
	modify := ldap.NewModifyRequest(groupDN, nil)
	modify.Add("uniqueMember", []string{userDN})

	if err := conn.Modify(modify); err != nil {
		return fmt.Errorf("[ldap] 添加用户到组失败: %v", err)
	}

	return nil
}

// 从指定的LDAP组中移除用户
func (l *LDAPServer) RemoveUserFromGroup(username, groupName string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	// 查找用户DN
	userDN, err := l.findUserDN(conn, username)
	if err != nil {
		return err
	}

	// 构造组DN
	groupDN := fmt.Sprintf("cn=%s,cn=group,%s", groupName, l.SplitDn(l.BindDN))

	// 修改组移除成员
	modify := ldap.NewModifyRequest(groupDN, nil)
	modify.Delete("uniqueMember", []string{userDN})

	if err := conn.Modify(modify); err != nil {
		return fmt.Errorf("[ldap] 从组中移除用户失败: %v", err)
	}

	return nil
}

// 获取用户所属的所有组
func (l *LDAPServer) GetUserGroups(username string) ([]string, error) {
	conn, err := l.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	userDN, err := l.findUserDN(conn, username)
	if err != nil {
		return nil, err
	}

	searchRequest := ldap.NewSearchRequest(
		fmt.Sprintf("cn=group,%s", l.SplitDn(l.BindDN)),
		//"cn=group,dc=firecloud,dc=com",
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf("(&(objectClass=groupOfUniqueNames)(uniqueMember=%s))", userDN),
		[]string{"cn"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("[ldap] 搜索用户组失败: %v", err)
	}

	groups := make([]string, 0, len(sr.Entries))
	for _, entry := range sr.Entries {
		groups = append(groups, entry.GetAttributeValue("cn"))
	}

	return groups, nil
}

func (l *LDAPServer) AddLDAPUser(username, nickName, password, phone, mail string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	userDN := fmt.Sprintf("uid=%s,%s", username, l.BaseDN)
	_, err = l.findUserDN(conn, username)
	if err == nil {
		return fmt.Errorf("[ldap] User already exists: %v", err)
	}

	userAttr, err := l.ParseUserAttr()
	fmt.Println(userAttr)
	if err != nil {
		return err
	}

	addReq := ldap.NewAddRequest(userDN, nil)
	addReq.Attribute("objectClass", []string{"top", "person", "organizationalPerson", "inetOrgPerson"})
	addReq.Attribute(userAttr.NickName, []string{nickName})
	addReq.Attribute(userAttr.Username, []string{username})
	addReq.Attribute("userPassword", []string{password}) // 这个字段没做映射
	addReq.Attribute(userAttr.Mobile, []string{phone})
	addReq.Attribute(userAttr.Email, []string{mail})

	//// 添加其他属性
	//for key, values := range attributes {
	//	addReq.Attribute(key, values)
	//}

	// 执行添加操作
	if err := conn.Add(addReq); err != nil {
		return fmt.Errorf("[ldap] failed to add user: %v", err)
	}

	return nil
}

func (l *LDAPServer) SplitDn(bindDN string) string {
	parts := strings.Split(bindDN, ",")
	if len(parts) >= 2 {
		result := parts[1:]
		resultStr := strings.Join(result, ",")
		return resultStr
	}
	return ""
}

func (l *LDAPServer) AddNonExistentMemberOfGroup(group string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	addRequest := ldap.NewAddRequest(fmt.Sprintf("cn=%s,cn=group,%s", group, l.SplitDn(l.BindDN)), nil)
	addRequest.Attribute("objectClass", []string{"groupOfUniqueNames"})
	addRequest.Attribute("uniqueMember", []string{""})
	addRequest.Attribute("cn", []string{group})

	if err := conn.Add(addRequest); err != nil {
		return fmt.Errorf("[ldap] 添加memberof属性组失败: %v", err)
	}

	return nil
}

func (l *LDAPServer) DeleteMemberOfGroup(group string) error {
	conn, err := l.Conn()
	if err != nil {
		return err
	}
	defer conn.Close()

	groupDN := fmt.Sprintf("cn=%s,cn=group,%s", group, l.SplitDn(l.BindDN))
	delRequest := ldap.NewDelRequest(groupDN, nil)
	if err = conn.Del(delRequest); err != nil {
		return fmt.Errorf("[ldap] 删除memberof属性组失败: %v", err)
	}

	return nil
}

// 查询指定memberof属性组下的所有用户
func (l *LDAPServer) GetUsersInMemberOfGroup(group string) ([]string, error) {
	conn, err := l.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	groupDN := fmt.Sprintf("cn=%s,cn=group,%s", group, l.SplitDn(l.BindDN))
	filter := fmt.Sprintf("(memberOf=%s)", groupDN)
	userAttr, err := l.ParseUserAttr()
	if err != nil {
		return nil, err
	}

	sr, err := l.search(conn, filter, userAttr)
	if err != nil {
		return nil, err
	}

	var users []string
	for _, entry := range sr.Entries {
		users = append(users, entry.DN)
	}

	return users, nil
}

func (l *LDAPServer) GetAllMemberOfGroups() ([]string, error) {
	conn, err := l.Conn()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// 搜索所有用户
	searchRequest := ldap.NewSearchRequest(
		l.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(objectClass=*)",
		[]string{"memberOf"},
		nil,
	)

	sr, err := conn.Search(searchRequest)
	if err != nil {
		return nil, fmt.Errorf("[ldap] failed to search ldap memberof group: %v", err)
	}

	// 使用map来去重
	groupsMap := make(map[string]bool)

	for _, entry := range sr.Entries {
		for _, memberOf := range entry.GetAttributeValues("memberOf") {
			groupsMap[memberOf] = true
		}
	}

	// 将map转换为slice
	groups := lo.MapToSlice(groupsMap, func(key string, value bool) string {
		parts := strings.Split(key, ",")
		if len(parts) >= 2 {
			result := parts[:1]
			resultStr := strings.Join(result, ",")
			key = strings.TrimPrefix(resultStr, "cn=")
		}
		return key
	})

	return groups, nil
}
