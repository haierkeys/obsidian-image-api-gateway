[中文文档](readme-zh.md) / [English Document](README.md)
# Obsidian Image API Gateway

<p align="center">
    <img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
    <img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway" alt="license">
</p>

该项目为 [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian 插件提供图片上传、存储与云同步服务。

## 功能清单

- [x] 支持图片上传
- [x] 支持令牌授权，提升 API 安全性
- [x] 支持图片 HTTP 访问（基础功能，建议使用 Nginx 替代）
- [x] 存储支持：
  - [x] 同时保存至本地或云存储，方便后续迁移
  - [x] 本地存储支持（为 NAS 准备，功能已测试）
  - [x] 支持阿里云 OSS 云存储（功能已实现，尚未测试）
  - [x] 支持 Cloudflare R2 云存储（功能已实现，已测试）
  - [x] 支持 Amazon S3（功能已实现，已测试）
  - [x] 增加 MinIO 存储支持（v1.5+）
  - [ ] 支持 Google ECS（待开发）
- [x] 提供 Docker 安装支持，便于在家庭 NAS 或远程服务器上使用
- [x] 提供公共服务 API && Web界面，方便提供公共服务 <a href="#userapi">用户公共接口 & Web 界面</a>

## 更新日志

查看完整的更新内容，请访问 [Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases)。

## 价格

本软件是开源且免费的。如果您想表示感谢或帮助支持继续开发，可以通过以下方式为我提供支持：

[<img src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" alt="BuyMeACoffee" width="100">](https://ko-fi.com/haierkeys)

## 快速开始

### 容器化安装（Docker 方式）

假设您的服务器图片保存路径为 _/data/storage/uploads_，依次执行以下命令：

```bash
# 拉取最新的容器镜像
docker pull haierkeys/obsidian-image-api-gateway:latest

# 创建项目所需的目录
mkdir -p /data/image-api
cd /data/image-api

mkdir -p ./config && mkdir -p ./storage/logs && mkdir -p ./storage/uploads

# 下载默认配置文件到配置目录
wget  -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml

# 创建并启动容器
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest
```

### 二进制安装

从 [Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases) 下载最新版本，解压后执行：

```bash
./image-api run -c config/config.yaml
```

### 配置

默认的配置文件名为 _config.yaml_，请将其放置在 _根目录_ 或 _config_ 目录下。

更多配置详情请参考：

- [配置文件 - 中文注释](config/config.yaml)
- [配置文件 - 英文注释](config/config-en.yaml)

### 开放服务 - 用户公共接口 & Web 界面
<span id="lable"></span>

![Image](https://github.com/user-attachments/assets/39c798de-b243-42c1-a75a-cd179913fc49)

- Web 界面：[http://IP:[[config:http-port](config/config.yaml#http-port)]]
- 配置设置：[config:database](config/config.yaml#database) 和 [config:user](config/config.yaml#user)

## 其他资源

- [Obsidian Auto Image Remote Uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)