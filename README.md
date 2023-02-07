# 简介
这是福佑家和go脚手架架构，为机构后台、运营后台提供创作者管理能力。



## 架构更新或插件加载

```
go mod tidy
```

## 编译项目

```
go build main.go
```

## swagger生成文档

```
go get -u github.com/swaggo/swag/cmd/swag

swag init
```

## 执行项目

```
./main.go
```



## 数据库
```
   /*
Source Server         : 127.0.0.1
Source Server Type    : MySQL
Source Server Version : 50739
Source Host           : 127.0.0.1:3306
Source Schema         : test

Target Server Type    : MySQL
Target Server Version : 50739
File Encoding         : 65001

Date: 30/12/2022 14:20:29
*/
SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
`id` bigint(20) NOT NULL AUTO_INCREMENT,
`userid` bigint(20) NOT NULL,
`username` varchar(64) NOT NULL,
`password` varchar(64) NOT NULL,
`email` varchar(64) DEFAULT NULL,
`gender` tinyint(4) DEFAULT '0',
`create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
`update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
PRIMARY KEY (`id`),
UNIQUE KEY `idx_username` (`username`) USING BTREE,
UNIQUE KEY `idx_user_id` (`userid`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=14 DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
```