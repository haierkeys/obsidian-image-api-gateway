FROM woahbase/alpine-glibc:latest
ARG TARGETOS
ARG TARGETARCH
ENV TZ=Asia/Shanghai
ENV P_NAME=api
ENV P_BIN=image-api
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories
RUN apk --update add libstdc++ curl ca-certificates bash curl gcompat tzdata && \
    cp /usr/share/zoneinfo/${TZ} /etc/localtime && \
    echo ${TZ} > /etc/timezone && \
    rm -rf  /tmp/* /var/cache/apk/*

EXPOSE 9000 9001
RUN mkdir -p /${P_NAME}/
VOLUME /${P_NAME}/config
VOLUME /${P_NAME}/storage
COPY ./build/${TARGETOS}_${TARGETARCH}/${P_BIN} /${P_NAME}/

# 将脚本复制到容器中
COPY entrypoint.sh /entrypoint.sh

# 给脚本执行权限
RUN chmod +x /entrypoint.sh

# 使用 ENTRYPOINT 执行脚本
ENTRYPOINT ["/entrypoint.sh"]