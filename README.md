# 简介
gofarm是服务端框架用Gin框架来实现，用restful设计风格和开发方式，为机构后台、运营后台提供创作者管理能力。



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
## 目录说明
```
├── Infrastructures 架构基础服务
│         ├── Init.go   服务初始化
│         ├── Redis Redis服务类
│         └── Rom   Mysql服务类
├── Interfaces  接口声明
│         └── User  用户接口声明
├── LICENSE
├── Logics  服务逻辑层
│         ├── Public    公共数据操作类
│         └── User  用户数据操作类
├── Middlewares 中间件服务类
│         └── Auth.go
├── Models  模型层
│         ├── Response  Swagger针对Service发布时Response接口模型
│         ├── Swagger   Swagger针对Service发布时请求接口模型
│         └── User  用户数据库模型层
├── Pkg 项目引入包
│         ├── HttpResponse  Response消息返回类
│         ├── Jwt   Jwt密钥类
│         └── Logger    日志服务类
├── README.md
├── Routes  路由配置目录
│         └── Routes.go 路由服务类
├── Services    服务接口类
│         ├── Public    公共服务类
│         └── User      用户接口类实现
├── Settings    架构配置目录
│         ├── config.yaml   配置文件
│         └── settings.go   配置文件解析类
├── docs    swagger生成的接口相关问文档
│         ├── docs.go
│         ├── swagger.json
│         └── swagger.yaml
├── go.mod
├── go.sum
├── main    编译文件
├── main.go 启动入口
└── web_app.log
```
