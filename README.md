# gogo-server

# 运转原理

1. 首先创建一个服务器，让服务器不断轮询读取连接传进来的值
2. 创建客户端，客户端建立连接，创建一个conn，
3. 使用Gob对conn进行包装，创建一个头部和body，传进gob包装的conn
4. 服务器那边接收到信息，对消息进行处理
5. 服务器这边也使用gob对conn进行包装，读取conn中传过来的值，读取解析header和body。
6. 对客户端发出响应
7. 客户端对头部和body进行解析，必须全部解析，不然会出现读取值时读取错的问题
   
