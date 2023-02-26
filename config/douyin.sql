/*
 Navicat Premium Data Transfer

 Source Server         : 腾讯云MySQL
 Source Server Type    : MySQL
 Source Server Version : 80032
 Source Host           : 1.15.97.114:3306
 Source Schema         : douyin

 Target Server Type    : MySQL
 Target Server Version : 80032
 File Encoding         : 65001

 Date: 26/02/2023 11:35:30
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
  `user_id` bigint NOT NULL COMMENT '评论发布用户id',
  `video_id` bigint NOT NULL COMMENT '评论视频id',
  `comment_text` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '评论内容',
  `create_date` datetime NOT NULL COMMENT '评论发布时间',
  `cancel` tinyint NOT NULL DEFAULT 0 COMMENT '默认评论发布为0，取消后为1',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `videoIdIdx`(`video_id`) USING BTREE COMMENT '评论列表使用视频id作为索引-方便查看视频下的评论列表'
) ENGINE = InnoDB AUTO_INCREMENT = 1215 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '评论表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1206, 17, 126, '好的', '2023-02-23 03:53:04', 1);
INSERT INTO `comments` VALUES (1207, 17, 126, '猜猜现在几点了', '2023-02-23 03:53:29', 1);
INSERT INTO `comments` VALUES (1208, 17, 126, '烦烦烦', '2023-02-23 03:57:42', 1);
INSERT INTO `comments` VALUES (1209, 13, 126, '乐', '2023-02-23 09:17:41', 1);
INSERT INTO `comments` VALUES (1210, 23, 127, '真好看', '2023-02-24 09:11:01', 1);
INSERT INTO `comments` VALUES (1211, 17, 127, 'hhhh', '2023-02-24 16:23:06', 1);
INSERT INTO `comments` VALUES (1212, 13, 121, '好好看', '2023-02-24 17:41:50', 1);
INSERT INTO `comments` VALUES (1213, 18, 121, '6666', '2023-02-24 18:44:35', 1);
INSERT INTO `comments` VALUES (1214, 17, 121, '嗨嗨嗨', '2023-02-24 19:55:46', 1);

-- ----------------------------
-- Table structure for follows
-- ----------------------------
DROP TABLE IF EXISTS `follows`;
CREATE TABLE `follows`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '用户id',
  `follower_id` bigint NOT NULL COMMENT '关注的用户',
  `cancel` tinyint NOT NULL DEFAULT 0 COMMENT '默认关注为0，取消关注为1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `userIdToFollowerIdIdx`(`user_id`, `follower_id`) USING BTREE,
  INDEX `FollowerIdIdx`(`follower_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '关注表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of follows
-- ----------------------------

-- ----------------------------
-- Table structure for likes
-- ----------------------------
DROP TABLE IF EXISTS `likes`;
CREATE TABLE `likes`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键',
  `user_id` bigint NOT NULL COMMENT '点赞用户id',
  `video_id` bigint NOT NULL COMMENT '被点赞的视频id',
  `cancel` tinyint NOT NULL DEFAULT 0 COMMENT '默认点赞为0，取消赞为1',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `userIdtoVideoIdIdx`(`user_id`, `video_id`) USING BTREE,
  INDEX `userIdIdx`(`user_id`) USING BTREE,
  INDEX `videoIdx`(`video_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1241 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '点赞表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of likes
-- ----------------------------
INSERT INTO `likes` VALUES (1229, 13, 122, 2);
INSERT INTO `likes` VALUES (1230, 13, 121, 2);
INSERT INTO `likes` VALUES (1231, 17, 123, 1);
INSERT INTO `likes` VALUES (1232, 17, 122, 1);
INSERT INTO `likes` VALUES (1233, 17, 121, 1);
INSERT INTO `likes` VALUES (1234, 17, 124, 1);
INSERT INTO `likes` VALUES (1235, 13, 125, 1);
INSERT INTO `likes` VALUES (1236, 17, 126, 1);
INSERT INTO `likes` VALUES (1237, 13, 123, 1);
INSERT INTO `likes` VALUES (1238, 23, 127, 1);
INSERT INTO `likes` VALUES (1239, 17, 127, 1);
INSERT INTO `likes` VALUES (1240, 18, 131, 1);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '用户id，自增主键',
  `name` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户名',
  `password` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '用户密码',
  `avatar` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'http://1.15.97.114/user/danji.jpg' COMMENT '用户头像',
  `background_image` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT 'http://1.15.97.114/user/1.jpg' COMMENT '用户个人页顶部大图',
  `signature` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL DEFAULT '心有所向，无问西东' COMMENT '个人简介',
  `work_count` bigint NOT NULL COMMENT '作品数',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `name_password_idx`(`name`, `password`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '用户表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'fgdkgh', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (2, 'fgdkgh`', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (3, 'fgdk', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (4, 'fgdkgh123', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (5, 'fsdffsdf', '6c039fac2de4de868f90af9c7d9578fb0b0da4faef31b6ee78601d9776f1cd16', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (6, 'fsdffsdfsgfsd', '6c039fac2de4de868f90af9c7d9578fb0b0da4faef31b6ee78601d9776f1cd16', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (7, 'ergsdfg', '0b832e22d5ce23b65d2ef471d303ae5089875ca0e0bbdc6c49df3b08bf7bade5', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (8, '乔涛', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (9, '7', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (10, 'maria', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (11, 'angela', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (12, 'gdfgdfg', '1419b62a008ceee9a8f95112f854240b8519212bc120c71015c22ed1f3afb35d', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (13, 'admin', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (14, 'fjghfjhgjhf', '2dce67b1101d3b86bca52647c533c13a979e401398c8ebdd1c4e21d136eaf941', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (15, 'sgdfsg', '7adf65f510ea45d895d0a80d6adbb8fdedfcf68019507e55ce7bb36421d76ecb', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (16, 'yangming', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (17, 'hhh', '8d32cf11af0f9646fb25c4fd362287a08a65f34bd378703c306cb4d2c8f54a1a', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (18, 'hhhhh', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (19, 'michelle', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (20, 'ronald', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (21, 'david', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (22, 'richard', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);
INSERT INTO `users` VALUES (23, 'hhhhhh', 'b946ccc987465afcda7e45b1715219711a13518d1f1663b8c53b848cb0143441', 'http://1.15.97.114/user/danji.jpg', 'http://1.15.97.114/user/1.jpg', '心有所向，无问西东', 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `id` bigint NOT NULL AUTO_INCREMENT COMMENT '自增主键，视频唯一id',
  `author_id` bigint NOT NULL COMMENT '视频作者id',
  `play_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '播放url',
  `cover_url` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NOT NULL COMMENT '封面url',
  `publish_time` datetime NOT NULL COMMENT '发布时间戳',
  `title` varchar(255) CHARACTER SET utf8mb3 COLLATE utf8mb3_general_ci NULL DEFAULT NULL COMMENT '视频名称',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `time`(`publish_time`) USING BTREE,
  INDEX `author`(`author_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 133 CHARACTER SET = utf8mb3 COLLATE = utf8mb3_general_ci COMMENT = '\r\n视频表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (121, 13, 'http://1.15.97.114/videos/4bcfa46d-b509-44e7-8a49-54000f138eaf.mp4', 'http://1.15.97.114/images/4bcfa46d-b509-44e7-8a49-54000f138eaf.jpg', '2023-02-22 12:08:14', '哈哈哈哈哈');
INSERT INTO `videos` VALUES (122, 13, 'http://1.15.97.114/videos/66ecb3e7-4d46-4a8e-8870-41239020ccfb.mp4', 'http://1.15.97.114/images/66ecb3e7-4d46-4a8e-8870-41239020ccfb.jpg', '2023-02-22 12:36:40', 'hai');
INSERT INTO `videos` VALUES (123, 16, 'http://1.15.97.114/videos/c45c8bd1-b293-4222-9aa7-a0996aa2878a.mp4', 'http://1.15.97.114/images/c45c8bd1-b293-4222-9aa7-a0996aa2878a.jpg', '2023-02-22 16:33:52', 'kjsdhfkjahwbe');
INSERT INTO `videos` VALUES (124, 13, 'http://1.15.97.114/videos/23c95924-7c35-422a-b5ea-9a52125b1a5d.mp4', 'http://1.15.97.114/images/23c95924-7c35-422a-b5ea-9a52125b1a5d.jpg', '2023-02-22 20:13:19', 'ghfhgjcvgh');
INSERT INTO `videos` VALUES (125, 17, 'http://1.15.97.114/videos/ae806a2c-1293-4913-839b-19be68cb0109.mp4', 'http://1.15.97.114/images/ae806a2c-1293-4913-839b-19be68cb0109.jpg', '2023-02-23 03:34:45', '凌晨3.34了');
INSERT INTO `videos` VALUES (126, 17, 'http://1.15.97.114/videos/5bef3802-3883-4369-8d8f-6819c54212af.mp4', 'http://1.15.97.114/images/5bef3802-3883-4369-8d8f-6819c54212af.jpg', '2023-02-23 03:52:05', '3.51了现在');
INSERT INTO `videos` VALUES (127, 13, 'http://1.15.97.114/videos/abc45912-f0f7-41c2-9335-a39298e48bd4.mp4', 'http://1.15.97.114/images/abc45912-f0f7-41c2-9335-a39298e48bd4.jpg', '2023-02-24 08:55:10', 'hhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhhh');
INSERT INTO `videos` VALUES (128, 18, 'http://1.15.97.114/videos/1732da97-aa6e-4938-bb99-58d7772b69ab.mp4', 'http://1.15.97.114/images/1732da97-aa6e-4938-bb99-58d7772b69ab.jpg', '2023-02-24 12:12:02', 'hhhh');
INSERT INTO `videos` VALUES (129, 17, 'http://1.15.97.114/videos/709d29ae-275c-487f-9f3c-4bf44f26c8bf.mp4', 'http://1.15.97.114/images/709d29ae-275c-487f-9f3c-4bf44f26c8bf.jpg', '2023-02-24 16:24:43', 'ghjgjhghj');
INSERT INTO `videos` VALUES (130, 13, 'http://1.15.97.114/videos/fbfb66d2-8475-4bf7-9f5f-8bd38449bfc6.mp4', 'http://1.15.97.114/images/fbfb66d2-8475-4bf7-9f5f-8bd38449bfc6.jpg', '2023-02-24 17:42:58', '感冒吃药啊');
INSERT INTO `videos` VALUES (131, 23, 'http://1.15.97.114/videos/f917b8e6-61a4-4fb2-a34c-6e838feeae57.mp4', 'http://1.15.97.114/images/f917b8e6-61a4-4fb2-a34c-6e838feeae57.jpg', '2023-02-24 18:42:26', '药吃了吗？');
INSERT INTO `videos` VALUES (132, 17, 'http://1.15.97.114/videos/418005eb-89c9-426d-b5a2-e2ec75a55af4.mp4', 'http://1.15.97.114/images/418005eb-89c9-426d-b5a2-e2ec75a55af4.jpg', '2023-02-24 19:55:27', '你好');

-- ----------------------------
-- Procedure structure for addFollowRelation
-- ----------------------------
DROP PROCEDURE IF EXISTS `addFollowRelation`;
delimiter ;;
CREATE PROCEDURE `addFollowRelation`(IN user_id int, IN follower_id int)
BEGIN
    #Routine body goes here...
    # 声明记录个数变量。
    DECLARE cnt INT DEFAULT 0;
    # 获取记录个数变量。
    SELECT COUNT(1) FROM follows f where f.user_id = user_id AND f.follower_id = follower_id INTO cnt;
    # 判断是否已经存在该记录，并做出相应的插入关系、更新关系动作。
    # 插入操作。
    IF cnt = 0 THEN
        INSERT INTO follows(`user_id`, `follower_id`) VALUES (user_id, follower_id);
    END IF;
    # 更新操作
    IF cnt != 0 THEN
        UPDATE follows f SET f.cancel = 0 WHERE f.user_id = user_id AND f.follower_id = follower_id;
    END IF;
END
;;
delimiter ;

-- ----------------------------
-- Procedure structure for delFollowRelation
-- ----------------------------
DROP PROCEDURE IF EXISTS `delFollowRelation`;
delimiter ;;
CREATE PROCEDURE `delFollowRelation`(IN `user_id` int, IN `follower_id` int)
BEGIN
    #Routine body goes here...
    # 定义记录个数变量，记录是否存在此关系，默认没有关系。
    DECLARE cnt INT DEFAULT 0;
    # 查看是否之前有关系。
    SELECT COUNT(1) FROM follows f WHERE f.user_id = user_id AND f.follower_id = follower_id INTO cnt;
    # 有关系，则需要update cancel = 1，使其关系无效。
    IF cnt = 1 THEN
        UPDATE follows f SET f.cancel = 1 WHERE f.user_id = user_id AND f.follower_id = follower_id;
    END IF;
END
;;
delimiter ;

SET FOREIGN_KEY_CHECKS = 1;
