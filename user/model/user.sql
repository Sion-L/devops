CREATE TABLE `auth_users`  (
    `user_id` bigint unsigned NOT NULL COMMENT '用户ID',
    `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户登录名',
    `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户登录密码',
    `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
    `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT "system" COMMENT '角色名;管理员,普通用户',
    `source` varchar(30) NOT NULL DEFAULT 'local' COMMENT '用户类型1local 2ldap',
    `role_type` tinyint(1) NOT NULL DEFAULT 1 COMMENT '角色类型1管理 2用户',
    `email` varchar(190) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NULL DEFAULT '888' COMMENT '用户邮箱',
    `mobile` bigint(0) NULL DEFAULT NULL COMMENT '用户手机号',
    PRIMARY KEY (`user_id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TABLE `auth_menus` (
    `menu_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '菜单ID',
    `menu_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '菜单名称',
    `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单路径',
    `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
    `parent_id` int(11) NULL DEFAULT NULL COMMENT '父菜单ID',
    `order` int(11) NULL DEFAULT 0 COMMENT '菜单排序',
    PRIMARY KEY (`menu_id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TABLE `auth_role_menu` (
    `role_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT "system" COMMENT '角色名;管理员,普通用户',
    `menu_id` int(11) NOT NULL COMMENT '菜单ID',
    PRIMARY KEY (`menu_id`)
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

CREATE TABLE `settings` (
    `id` int NOT NULL AUTO_INCREMENT,
    `name` varchar(128) NOT NULL,
    `value` longtext NOT NULL,
    `category` varchar(128) NOT NULL COMMENT '设置类别',
    `encrypted` tinyint(1) NOT NULL COMMENT 'value是否加密',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=42 DEFAULT CHARSET=utf8;
