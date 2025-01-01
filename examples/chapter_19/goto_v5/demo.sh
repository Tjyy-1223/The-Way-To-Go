#!/bin/sh   声明使用的是 Bash Shell 解释器

# 在后台启动 `goto` 服务，使用 HTTP 协议监听 8081 端口，并启用 RPC 服务
# `-http=:8081` 表示在 8081 端口启动 HTTP 服务
# `-rpc=true` 启用 RPC 服务
go run . -http=:8081 -rpc=true &

# 获取主服务的进程 ID (PID)，并将其保存在变量 `master_pid` 中
master_pid=$!

# 等待 1 秒钟，确保主服务已经成功启动
sleep 1

# 启动 `goto` 的从服务，连接到主服务 (127.0.0.1:8081)
# `-master=127.0.0.1:8081` 表示从服务连接到主服务，地址为本地的 8081 端口
go run . -master=127.0.0.1:8081 &

# 获取从服务的进程 ID (PID)，并将其保存在变量 `slave_pid` 中
slave_pid=$!

# 输出信息，告知用户主服务和从服务的端口
echo "Running master on :8081, slave on :8080."
# 提示用户访问从服务的接口
echo "Visit: http://localhost:8080/add"
# 提示用户按回车键来关闭服务
echo "Press enter to shut down"

# 等待用户按下回车键
read

# 用户按下回车键后，使用 `kill` 命令结束主服务和从服务的进程
kill $master_pid
kill $slave_pid