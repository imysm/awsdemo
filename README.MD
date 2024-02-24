# fearture awsdemo

    简单实现一个读写aws rds for mysql 、elastic cache for redis的接口

## 简介

    1. mysql和redis配置信息从secret manager 读取
    2. secret mange的登录校验信息从环境变量或者aws的默认配置文件读取
    3. 建立 rds for mysql 连接, 测试数据写入和读取
    4. 建立 alastic cache for redis 链接 测试写入和读取
    5. 简单验证, token 可在config中自定义 ,如未配置则使用默认tokn

## 测试说明

    curl "http://127.0.0.1/awsdemo/test/mysqlInsert?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJjb3JwIjoiTWljdnMiLCJpYXQiOjE3MDg2N&name=samyang&message=Imysm and is testing"
    curl "http://127.0.0.1/awsdemo/test/redisInsert?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9eyJjb3JwIjoiTWljdnMiLCJpYXQiOjE3MDg2N&name=samyang&message=Imysm and is testing"

## 构建说明

    1. 编译代码: CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-w"  -o target/awsdemo cmd/demo/main.go
    2. 构建镜像: docker build -t awsdemo:v0.1 -f deploy/Dockerfile .
    3. docker tag awsdemo:v0.1  远程仓库uri
    4. docker push 远程仓库uri
    5. 在aws ecs 部署 参考ecs 文档