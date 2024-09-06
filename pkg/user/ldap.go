package user

import (
	"encoding/json"
	"fmt"
	"github.com/go-ldap/ldap/v3"
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
			Username: entry.GetAttributeValue(userAttr.Username),
			Email:    entry.GetAttributeValue(userAttr.Email),
			NickName: entry.GetAttributeValue(userAttr.NickName),
			Mobile:   entry.GetAttributeValue(userAttr.Mobile),
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

func (l *LDAPServer) ParseUserAttr() (User, error) {
	var user User
	err := json.Unmarshal([]byte(l.UserAttr), &user)
	if err != nil {
		return User{}, fmt.Errorf("error parsing JSON: %v", err)
	}

	return user, nil
}
