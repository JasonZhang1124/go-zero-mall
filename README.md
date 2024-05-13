# 项目简介
#### go-zero实战：简单的微服务练习
#### 项目技术栈
* mysql
* gorm
* go-zero
* grpc
* etcd

#### 文件介绍
* common: 系统配置文件、中间件、工具等
* service: 系统的微服务：
  * user用户服务
  * product产品服务
  * order订单服务
  * pay支付服务

### 使用
* go语言环境
* MySQL数据库
* etcd 
* api测试工具

1.克隆项目到本地
```

```
2.安装相关依赖
```
go mod tidy
```
3.启动每一个系统微服务文件中api和rpc的[服务名].go文件
如启动user用户服务
```api
go-zero-mall\service\user\rpc> go run user.go
go-zero-mall\service\user\api> go run user.go
```
如上步骤继续启动：product产品服务、order订单服务、pay支付服务