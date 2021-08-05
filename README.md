# sparrow-tools


# 启动参数
* --timeout=20s
* --listen-port=:10000
* --inter-path=/inter/timeout/simulate

server 监听到启动
http://本机IP:10000/inter/timeout/simulate

客户端调用server接口时，20s后服务端处理完成。

以上为：模拟服务端处理时间长的接口。
