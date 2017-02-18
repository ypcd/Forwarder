# Forwarder
Network data Forwarder


端口转发器

功能：
支持本机和远程ip端口tcp数据转发。
支持tcp短连接（http等）和长连接（基于tcp大容量（GiB以上）文件传输）。

工作流程简述：
转发器建立监听端口p1，接收到新连接请求后，建立新连接s1，再建立与目标地址端口p2新连接s2，然后s1、s2组合为一条新转发线路。最后开始转发数据。

基于go语言开发，使用go默认net库。Go 语言层面采用阻塞+多协程模式，进行网络通信。

网络模型：
因为go的默认net库，底层基于非阻塞+多路复用模型（windows iocp、linux epoll），
所以转发器实质模型为非阻塞+多路复用模型。保证了转发器的高性能。

性能：
与之前的python版本相比，性能提高200-300倍。因为python版本，网络模型采用阻塞+多线程轮询模式，并非python的高性能网络模型。这样的性能差异与编程语言的执行效率并无太大关联。如果python使用异步多路复用模型，两者的性能差应该在10-20倍之间。

基于http代理端口转发性能测试，python版本典型cpu占用为20%-30%，go版本典型cpu占用为0.05%-1%。

性能测试硬件平台：

Cpu  i3  2核心，4线程
内存 12GiB
硬盘 三星固态硬盘 120GiB


开源协议：
项目基于GPLv3协议开源。


The following is from the google translation.

Port forwarders

Features:
Support local and remote ip port tcp data forwarding.
Support tcp short connection (http, etc) and long connection (based on tcp large capacity (GiB above) file transfer).

Workflow Brief:
The transponder establishes the listening port p1, and after receiving the new connection request, establishes a new connection s1, and then establishes a new connection s2 with the destination address port p2, and then s1, s2 is combined into a new forwarding line. Finally start forwarding data.

Based on the go language development, use go default net library. Go language level using blocking + multi-link mode, the network communication.

Network model:
Because the default net library for go, the underlying layer based on non-blocking + multiplexing (windows iocp, linux epoll)
So the transponder substantive model is a non-blocking + multiplexing model. To ensure the high performance of the transponder.

performance:
Compared with the previous version of python, performance increased by 200-300 times. Because python version, the network model uses blocking + multi-threaded polling mode, not python's high-performance network model. This performance difference is not much related to the execution efficiency of the programming language. If python uses an asynchronous multiplexing model, the performance difference between the two should be between 10 and 20 times.

Based on the http proxy port forwarding performance test, python version of the typical cpu occupies 20% -30%, go version of the typical cpu occupies 0.05% -1%.

Performance Test Hardware Platform:

Cpu i3 2 core, 4 threads
Memory 12GiB
Hard disk Samsung solid state hard drive 120GiB


Open source agreement:
The project is based on the GPLv3 protocol.

