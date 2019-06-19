# 首先拉取基础镜像
FROM golang:alpine
# 将当前项目的代码放到镜像的指定位置
ADD ./src /go/src/web-go
# 将项目依赖的第三方库复制到镜像的指定位置，注意这里宿主机的位置应该在上下文之中
COPY ./vendor /go/src
# 使用go 编译项目
RUN go install web-go

# 拉取基础镜像
FROM alpine:latest
# 将编译后的二进制拷贝到当前目录
COPY --from=0 /go/bin/web-go .
# 暴露端口
ENV PORT 3001
# 运行命令
CMD ["./web-go"]