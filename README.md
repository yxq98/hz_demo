# hz_demo
http框架hertz demo
[官方文档](https://www.cloudwego.io/zh/docs/hertz/overview/)

## 安装脚手架

```sh
go install github.com/cloudwego/hertz/cmd/hz@latest
```

## 初始化项目

```sh
mkdir demo
cd demo
hz new --module ${go_module_name} --idl ${idl_path} --service ${service_name}

go mod edit -replace github.com/apache/thrift=github.com/apache/thrift@v0.13.0   

go mod tidy
```

# 编译

```sh
sh build.sh
```