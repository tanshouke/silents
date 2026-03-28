### 配置初始化
说明：初始化配置、全局变量、全局共享资源，提供统一的初始化入口（程序启动前调用）

```shell
# 1.root 指令

go run .\main.go

# 2.start 指令
go run .\main.go start -f "./etc/config.yaml"

# 3.test指令
go run .\main.go test  -t ttttttt -n nnnnnn

```