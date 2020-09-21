# 查找手机号归属
golang gin框架+github.com/xluohome/phonedata项目的封装

# 安装
```
go get -u github.com/gin-gonic/gin
go get github.com/kardianos/govendor

govendor init
govendor fetch github.com/gin-gonic/gin@v1.6.3
```
//创建main.go文件

//加载依赖
```
go mod init
```
//下载对应依赖
```
go mod vendor
//go run main.go

//win 
//指定 phone.dat 的文件路径
set PHONE_DATA_DIR='C:\Users\Administrator\go\src\phone'

//linux
//指定 phone.dat 的文件路径
export PHONE_DATA_DIR='/data/shelldir/'
setsid ./phone
```
