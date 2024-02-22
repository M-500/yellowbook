# 指定基础镜像
FROM ubuntu:20.04
# 把编译后的打包进来到这个镜像
COPY webook /app/webook
# 指定工作目录
WORKDIR /app

# 启动项目
ENTRYPOINT ["/app/webook"]