<p align="center">
  <img width="800" src="https://github.com/pchaibo/admin/blob/master/public/imges/top.png">
  <img width="800" src="https://github.com/pchaibo/admin/blob/master/public/imges/user.png">
   <img width="800" src="https://github.com/pchaibo/admin/blob/master/public/imges/add.png">
</p>

## 使用golang语言基于Gin + gorm +mysql 的前后端分离权限管理系统
前端vue vue-element-admin

git clone https://github.com/pchaibo/admin.git
## 后端
遵循 RESTful API 设计规范，只需要配置文件中，修改数据库连接

基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）

JWT 认证

后端接口:

git clone https://github.com/pchaibo/shop.git

配置文件

conf/app.conf

导入数据库文件

gin.sql

安装好golang环境执行

go run main.go

编译文件

cd shop

go build main.go
