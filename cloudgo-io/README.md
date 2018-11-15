
# 编程 web 应用程序 cloudgo-io

## 项目结构
1. service/js.go 
  返回JSON数据结构
  
2. service/server.go
  新建服务程序，建立路由表，处理GET请求和POST请求
3. service/table.go
  对POST请求处理，渲染静态html页面，返回用户名和时间
4. service/unknown.go
  对不存在的路径返回unknown
5. main.go
  启动程序
6. templates
  存放静态html

## 支持静态文件服务
![](./assets/image/getIndex.PNG)
## 支持简单 js 访问
![](./assets/image/getjs.PNG)
## 提交表单，并输出一个表格
![](./assets/image/gettable.PNG)

##  对 /unknown 给出开发中的提示，返回码 5xx
![](./assets/image/getUnknown.PNG)
