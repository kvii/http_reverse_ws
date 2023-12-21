# reverse 支持 websocket 的示例

* client 是 websocket 客户端，ws 连接通过反向代理服务器连接到服务端。
* reverse 是反向代理服务器。
* server 是 websocket 服务器，实现了 echo 功能。

## 使用方式

1. 新开一个终端，执行 `go run ./reverse` 启动反向代理服务器。
2. 新开一个终端，执行 `go run ./server` 启动服务端。
3. 新开一个终端，执行 `go run ./client` 启动客户端。

启动客户端后，在控制台输入任意字符，可以在服务器和客户端中看到输入的内容被打印。输入 "exit" 停止客户端。