[中文文档](readme-zh.md) / [English Document](README.md)

 <h1 align="center">Obsidian Image API Gateway</h1>

<p align="center">
<img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
<img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway.svg" alt="license" >
</p>

为 [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian插件提供图片上传/存储/同步云存储服务

功能清单:

- [x] 支持图片上传
- [x] 支持授权令牌,增加 API 安全
- [x] 图片http访问 (基础功能,建议使用 nginx 替代)
- [x] 存储相关:
  - [x] 支持同时保存到本地或云存储,方便后续迁移
  - [x] 支持本地保存 ( 为 NAS 准备,功能支持,测试 OK )
  - [x] 支持 阿里云 OSS 云存储( 功能支持,尚未测试 )
  - [x] 支持 Cloudflare R2 云存储( 功能支持,测试 OK )
  - [x] 支持 Amazon S3 ( 功能支持,测试OK )
  - [x] 增加 MinIO 存储支持 ( v1.5+ )
  - [ ] 支持 Google ECS ( 待开发 )
- [x] Docker命令安装,方便大家使用在家庭NAS和远端服务器中
- [ ] 增加公共API,针对不方便架设 API 服务的用户

## 变更日志

[Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases)



## 价格

本软件是开源软件,免费提供给大家的使用，但如果您想表示感谢或帮助支持继续开发，请随时通过以下任一方式为我提供一点帮助：

[<img src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" alt="BuyMeACoffee" width="100">](https://ko-fi.com/haierkeys)

# 开始使用

## 容器化安装 (docker方式)

假设您的服务器图片保存路径为 _/data/storage/uploads_

依次执行以下命令

```bash

# 下载容器镜像
docker pull haierkeys/obsidian-image-api-gateway:latest

# 创建项目运行必要目录
mkdir -p /data/image-api
cd /data/image-api

mkdir -p ./config && mkdir -p ./storage/logs && mkdir -p ./storage/uploads

# 下载默认配置到配置文件目录内
wget  -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml



# 创建&启动容器
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest

```

## 二进制下载安装

https://github.com/haierkeys/obsidian-image-api-gateway/releases 下载最新版本

解压到相应的目录执行

```bash
./image-api run -c config/config.yaml
```

## 配置

配置文件默认文件名 _config.yaml_, 需要直到 _根目录_ 或者 _config_ 目录内

配置详情参阅:

[配置文件-中文注释](config/config.yaml)
[配置文件-英文注释](config/config-en.yaml)

## TODO

## 其他

Obsidian Auto Image Remote Uploader

https://github.com/haierkeys/obsidian-auto-image-remote-uploader
