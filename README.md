<p align="center">
  <img width="800" src="https://github.com/pchaibo/shop/blob/master/web/public/imges/top.png">
  <img width="800" src="https://github.com/pchaibo/shop/blob/master/web/public/imges/user.png">
   <img width="800" src="https://github.com/pchaibo/shop/blob/master/web/public/imges/add.png">
</p>

## 使用golang语言基于Gin + gorm +mysql 的前后端分离权限管理系统
前端vue vue-element-admin

## 后端
* 遵循 RESTful API 设计规范，只需要配置文件中，修改数据库连接
* 基于 GIN WEB API 框架，提供了丰富的中间件支持（用户认证、跨域、访问日志、追踪ID等）
* JWT 认证
* 下载代码:
```bash
git clone https://github.com/pchaibo/shop.git
```
## 配置文件
conf/app.conf
导入数据库文件
gin.sql
安装好golang环境执行
```bash
go run main.go Trc20Block.go
```
## 浏览器访问：
http://127.0.0.1:805/ui/

## 编译文件
```bash
cd shop
go build main.go Trc20Block.go
```
## 前端使用vue开发
*/web
## 开发

```bash
# 克隆项目

# 进入项目目录
cd web

# 安装依赖
npm install

# 建议不要直接使用 cnpm 安装依赖，会有各种诡异的 bug。可以通过如下操作解决 npm 下载速度慢的问题
npm install --registry=https://registry.npm.taobao.org

# 启动服务
cd web
npm run dev
```

浏览器访问 http://localhost:9000

## 发布

```bash
# 构建测试环境
npm run build:stage

# 构建生产环境
npm run build:prod
```

## 其它

```bash
# 预览发布环境效果
npm run preview

# 预览发布环境效果 + 静态资源分析
npm run preview -- --report

# 代码格式检查
npm run lint

# 代码格式检查并自动修复
npm run lint -- --fix
```
usdt 测试
http://127.0.0.1:805/ui/#/demo
## Donate
