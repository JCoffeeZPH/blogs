CREATE TABLE `article_tab` (
                               `id` int NOT NULL AUTO_INCREMENT,
                               `user_id` int NOT NULL COMMENT '作者',
                               `category_id` int DEFAULT NULL COMMENT '文章分类',
                               `article_cover` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '文章缩略图',
                               `article_title` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标题',
                               `article_content` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '内容',
                               `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶 0否 1是',
                               `is_featured` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否推荐 0否 1是',
                               `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
                               `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密 3草稿',
                               `type` tinyint(1) NOT NULL DEFAULT '1' COMMENT '文章类型 1原创 2转载 3翻译',
                               `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '访问密码',
                               `original_url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '原文链接',
                               `ctime` int NOT NULL COMMENT '发表时间',
                               `mtime` int DEFAULT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `article_tag_tab` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `article_id` int NOT NULL COMMENT '文章id',
                                   `tag_id` int DEFAULT NULL,
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `category_tab` (
                                `id` int NOT NULL AUTO_INCREMENT,
                                `category_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '分类名',
                                `ctime` int NOT NULL COMMENT '创建时间',
                                `mtime` int DEFAULT NULL COMMENT '更新时间',
                                PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE `comment_tab` (
                               `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                               `user_id` int NOT NULL COMMENT '被评论用户Id',
                               `topic_id` int DEFAULT NULL COMMENT '评论主题id',
                               `comment_content` text CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '评论内容',
                               `reply_user_id` int DEFAULT NULL COMMENT '回复用户id',
                               `parent_id` int DEFAULT NULL COMMENT '父评论id',
                               `type` tinyint NOT NULL COMMENT '评论类型 1.文章 2.留言 3.关于我 4.友链 5.说说',
                               `is_delete` tinyint NOT NULL DEFAULT '0' COMMENT '是否删除  0否 1是',
                               `is_review` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否审核',
                               `ctime` int NOT NULL COMMENT '评论时间',
                               `mtime` int DEFAULT NULL COMMENT '更新时间',
                               PRIMARY KEY (`id`) USING BTREE,
                               KEY `fk_comment_user` (`user_id`) USING BTREE,
                               KEY `fk_comment_parent` (`parent_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `menu_tab` (
                            `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单名',
                            `path` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单路径',
                            `component` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '组件',
                            `icon` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '菜单icon',
                            `ctime` int NOT NULL COMMENT '创建时间',
                            `mtime` int DEFAULT NULL COMMENT '更新时间',
                            `order_num` tinyint(1) NOT NULL COMMENT '排序',
                            `parent_id` int DEFAULT NULL COMMENT '父id',
                            `is_hidden` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否隐藏  0否1是',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

INSERT INTO `menu_tab` VALUES (1, '首页', '/', '/home/Home.vue', 'el-icon-myshouye', 1680231036, 1680231036, 1, NULL, 0);
INSERT INTO `menu_tab` VALUES (2, '文章管理', '/article-submenu', 'Layout', 'el-icon-mywenzhang-copy', 1680231036, 1680231036, 2, NULL, 0);
INSERT INTO `menu_tab` VALUES (3, '消息管理', '/message-submenu', 'Layout', 'el-icon-myxiaoxi', 1680231036, 1680231036, 3, NULL, 0);
INSERT INTO `menu_tab` VALUES (4, '系统管理', '/system-submenu', 'Layout', 'el-icon-myshezhi', 1680231036, 1680231036, 5, NULL, 0);
INSERT INTO `menu_tab` VALUES (5, '个人中心', '/setting', '/setting/Setting.vue', 'el-icon-myuser', 1680231036, 1680231036, 7, NULL, 0);
INSERT INTO `menu_tab` VALUES (6, '发布文章', '/articles', '/article/Article.vue', 'el-icon-myfabiaowenzhang', 1680231036, 1680231036, 1, 2, 0);
INSERT INTO `menu_tab` VALUES (7, '修改文章', '/articles/*', '/article/Article.vue', 'el-icon-myfabiaowenzhang', 1680231036, 1680231036, 2, 2, 1);
INSERT INTO `menu_tab` VALUES (8, '文章列表', '/article-list', '/article/ArticleList.vue', 'el-icon-mywenzhangliebiao', 1680231036, 1680231036, 3, 2, 0);
INSERT INTO `menu_tab` VALUES (9, '分类管理', '/categories', '/category/Category.vue', 'el-icon-myfenlei', 1680231036, 1680231036, 4, 2, 0);
INSERT INTO `menu_tab` VALUES (10, '标签管理', '/tags', '/tag/Tag.vue', 'el-icon-myicontag', 1680231036, 1680231036, 5, 2, 0);
INSERT INTO `menu_tab` VALUES (11, '评论管理', '/comments', '/comment/Comment.vue', 'el-icon-mypinglunzu', 1680231036, 1680231036, 1, 3, 0);
INSERT INTO `menu_tab` VALUES (12, '用户列表', '/users', '/user/User.vue', 'el-icon-myyonghuliebiao', 1680231036, 1680231036, 1, 202, 0);
INSERT INTO `menu_tab` VALUES (13, '角色管理', '/roles', '/role/Role.vue', 'el-icon-myjiaoseliebiao', 1680231036, 1680231036, 2, 213, 0);
INSERT INTO `menu_tab` VALUES (14, '接口管理', '/resources', '/resource/Resource.vue', 'el-icon-myjiekouguanli', 1680231036, 1680231036, 2, 213, 0);
INSERT INTO `menu_tab` VALUES (15, '菜单管理', '/menus', '/menu/Menu.vue', 'el-icon-mycaidan', 1680231036, 1680231036, 2, 213, 0);
INSERT INTO `menu_tab` VALUES (16, '友链管理', '/links', '/friendLink/FriendLink.vue', 'el-icon-mydashujukeshihuaico-', 1680231036, 1680231036, 3, 4, 0);
INSERT INTO `menu_tab` VALUES (17, '关于我', '/about', '/about/About.vue', 'el-icon-myguanyuwo', 1680231036, 1680231036, 4, 4, 0);
INSERT INTO `menu_tab` VALUES (18, '日志管理', '/log-submenu', 'Layout', 'el-icon-myguanyuwo', 1680231036, 1680231036, 6, NULL, 0);
INSERT INTO `menu_tab` VALUES (19, '操作日志', '/operation/log', '/log/OperationLog.vue', 'el-icon-myguanyuwo', 1680231036, 1680231036, 1, 19, 0);
INSERT INTO `menu_tab` VALUES (20, '在线用户', '/online/users', '/user/Online.vue', 'el-icon-myyonghuliebiao', 1680231036, 1680231036, 7, 202, 0);
INSERT INTO `menu_tab` VALUES (21, '用户管理', '/users-submenu', 'Layout', 'el-icon-myyonghuliebiao', 1680231036, 1680231036, 4, NULL, 0);
INSERT INTO `menu_tab` VALUES (22, '相册管理', '/album-submenu', 'Layout', 'el-icon-myimage-fill', 1680231036, 1680231036, 5, NULL, 0);
INSERT INTO `menu_tab` VALUES (23, '相册列表', '/albums', '/album/Album.vue', 'el-icon-myzhaopian', 1680231036, 1680231036, 1, 205, 0);
INSERT INTO `menu_tab` VALUES (24, '照片管理', '/albums/:albumId', '/album/Photo.vue', 'el-icon-myzhaopian',1680231036, 1680231036, 1, 205, 1);
INSERT INTO `menu_tab` VALUES (25, '定时任务', '/quartz', '/quartz/Quartz.vue', 'el-icon-myyemianpeizhi', 1680231036, 1680231036, 2, 4, 0);
INSERT INTO `menu_tab` VALUES (26, '照片回收站', '/photos/delete', '/album/Delete.vue', 'el-icon-myhuishouzhan', 1680231036, NULL, 3, 205, 1);
INSERT INTO `menu_tab` VALUES (27, '权限管理', '/permission-submenu', 'Layout', 'el-icon-mydaohanglantubiao_quanxianguanli', 1680231036, 1680231036, 4, NULL, 0);
INSERT INTO `menu_tab` VALUES (28, '网站管理', '/website', '/website/Website.vue', 'el-icon-myxitong', 1680231036, NULL, 1, 4, 0);
INSERT INTO `menu_tab` VALUES (29, '定时任务日志', '/quartz/log/:quartzId', '/log/QuartzLog.vue', 'el-icon-myguanyuwo', 1680231036, 1680231036, 2, 19, 1);
INSERT INTO `menu_tab` VALUES (30, '说说管理', '/talk-submenu', 'Layout', 'el-icon-mypinglun', 1680231036, 1680231036, 3, NULL, 0);
INSERT INTO `menu_tab` VALUES (31, '说说列表', '/talk-list', '/talk/TalkList.vue', 'el-icon-myiconfontdongtaidianji', 1680231036, NULL, 1, 221, 0);
INSERT INTO `menu_tab` VALUES (32, '发布说说', '/talks', '/talk/Talk.vue', 'el-icon-myfabusekuai', 1680231036, 1680231036, 2, 221, 0);
INSERT INTO `menu_tab` VALUES (33, '修改说说', '/talks/:talkId', '/talk/Talk.vue', 'el-icon-myfabusekuai', 1680231036, 1680231036, 3, 221, 1);
INSERT INTO `menu_tab` VALUES (34, '异常日志', '/exception/log', '/log/ExceptionLog.vue', 'el-icon-myguanyuwo', 1680231036, 1680231036, 1, 19, 0);


CREATE TABLE `photo_album_tab` (
                                   `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                                   `album_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '相册名',
                                   `album_desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '相册描述',
                                   `album_cover` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '相册封面',
                                   `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
                                   `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态值 1公开 2私密',
                                   `ctime` int NOT NULL COMMENT '创建时间',
                                   `mtime` int DEFAULT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='相册';

CREATE TABLE `photo_tab` (
                             `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                             `album_id` int NOT NULL COMMENT '相册id',
                             `photo_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '照片名',
                             `photo_desc` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '照片描述',
                             `photo_src` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '照片地址',
                             `is_delete` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否删除',
                             `ctime` int NOT NULL COMMENT '创建时间',
                             `mtime` int DEFAULT NULL COMMENT '更新时间',
                             PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC COMMENT='照片';

CREATE TABLE `role_menu_tab` (
                                 `id` int NOT NULL AUTO_INCREMENT COMMENT '主键',
                                 `role_id` int DEFAULT NULL COMMENT '角色id',
                                 `menu_id` int DEFAULT NULL COMMENT '菜单id',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `role_tab` (
                            `id` int NOT NULL AUTO_INCREMENT COMMENT '主键id',
                            `role_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '角色名',
                            `is_disable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用  0否 1是',
                            `ctime` int NOT NULL COMMENT '创建时间',
                            `mtime` int DEFAULT NULL COMMENT '更新时间',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;


CREATE TABLE `tag_tab` (
                           `id` int NOT NULL AUTO_INCREMENT,
                           `tag_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '标签名',
                           `ctime` int NOT NULL COMMENT '创建时间',
                           `mtime` int DEFAULT NULL COMMENT '更新时间',
                           PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `user_auth_tab` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_info_id` int NOT NULL COMMENT '用户信息id',
                                 `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户名',
                                 `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '密码',
                                 `login_type` tinyint(1) NOT NULL COMMENT '登录类型(1: 邮箱, 2: qq)',
                                 `ip_address` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户登录ip',
                                 `ip_source` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT 'ip来源',
                                 `ctime` int NOT NULL COMMENT '创建时间',
                                 `mtime` int DEFAULT NULL COMMENT '更新时间',
                                 `last_login_time` int DEFAULT NULL COMMENT '上次登录时间',
                                 PRIMARY KEY (`id`) USING BTREE,
                                 UNIQUE KEY `username` (`username`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `user_info_tab` (
                                 `id` int NOT NULL AUTO_INCREMENT COMMENT '用户ID',
                                 `email` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '邮箱号',
                                 `nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '用户昵称',
                                 `avatar` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL DEFAULT '' COMMENT '用户头像',
                                 `intro` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '用户简介',
                                 `website` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '个人网站',
                                 `is_subscribe` tinyint(1) DEFAULT NULL COMMENT '是否订阅',
                                 `is_disable` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否禁用',
                                 `ctime` int NOT NULL COMMENT '创建时间',
                                 `mtime` int DEFAULT NULL COMMENT '更新时间',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `user_role_tab` (
                                 `id` int NOT NULL AUTO_INCREMENT,
                                 `user_id` int DEFAULT NULL COMMENT '用户id',
                                 `role_id` int DEFAULT NULL COMMENT '角色id',
                                 PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `website_config_tab` (
                                      `id` int NOT NULL AUTO_INCREMENT,
                                      `config` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '配置信息',
                                      `ctime` int NOT NULL COMMENT '创建时间',
                                      `mtime` int DEFAULT NULL COMMENT '更新时间',
                                      PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `talk_tab` (
                            `id` int NOT NULL AUTO_INCREMENT COMMENT '说说id',
                            `user_id` int NOT NULL COMMENT '用户id',
                            `content` varchar(2000) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '说说内容',
                            `images` varchar(2500) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '图片',
                            `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                            `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1.公开 2.私密',
                            `ctime` int NOT NULL COMMENT '创建时间',
                            `mtime` int DEFAULT NULL COMMENT '更新时间',
                            PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;

CREATE TABLE `unique_view_tab` (
                                   `id` int NOT NULL AUTO_INCREMENT,
                                   `views_count` int NOT NULL COMMENT '访问量',
                                   `ctime` int NOT NULL COMMENT '创建时间',
                                   `mtime` int DEFAULT NULL COMMENT '更新时间',
                                   PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci ROW_FORMAT=DYNAMIC;