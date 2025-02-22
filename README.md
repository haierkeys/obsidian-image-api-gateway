

[中文文档](readme-zh.md) / [English Document](README.md)

# Obsidian Image API Gateway

<p align="center">
    <img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
    <img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway" alt="license">
</p>

This project provides image upload, storage, and cloud synchronization services for the [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian plugin.

## Features List

- [x] Supports image upload
- [x] Supports token-based authorization for enhanced API security
- [x] Supports image HTTP access (basic functionality; it is recommended to use Nginx as a substitute)
- [x] Storage support:
  - [x] Save images both locally and in cloud storage for easier migration
  - [x] Local storage support (tested for NAS setup)
  - [x] Supports Alibaba Cloud OSS storage (feature implemented, not yet tested)
  - [x] Supports Cloudflare R2 storage (feature implemented, tested)
  - [x] Supports Amazon S3 (feature implemented, tested)
  - [x] Added support for MinIO storage (v1.5+)
  - [ ] Google ECS support (to be developed)
- [x] Provides Docker installation support for easy use on home NAS or remote servers
- [x] Provides public service API & Web interface for easier public service provision <a href="#userapi">Public API & Web Interface</a>

## Changelog

For the full list of updates, please visit the [Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

## Pricing

This software is open-source and free. If you wish to express your gratitude or support continued development, you can contribute by:

[<img src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" alt="BuyMeACoffee" width="100">](https://ko-fi.com/haierkeys)

## Quick Start

### Containerized Installation (Docker Method)

Assuming your server's image storage path is _/data/storage/uploads_, execute the following commands:

```bash
# Pull the latest container image
docker pull haierkeys/obsidian-image-api-gateway:latest

# Create the necessary directories for the project
mkdir -p /data/image-api
cd /data/image-api

mkdir -p ./config && mkdir -p ./storage/logs && mkdir -p ./storage/uploads

# Download the default configuration file to the config directory
wget  -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml

# Create and start the container
docker run -tid --name image-api \
        -p 8000:8000 -p 8001:8001 \
        -v /data/image-api/storage/logs/:/api/storage/logs/ \
        -v /data/image-api/storage/uploads/:/api/storage/uploads/ \
        -v /data/image-api/config/:/api/config/ \
        haierkeys/obsidian-image-api-gateway:latest
```

### Binary Installation

Download the latest version from [Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases), extract it, and execute:

```bash
./image-api run -c config/config.yaml
```

### Configuration

The default configuration file is named _config.yaml_ and should be placed in the _root directory_ or the _config_ directory.

For more configuration details, refer to:

- [Configuration File - Chinese Comments](config/config.yaml)
- [Configuration File - English Comments](config/config-en.yaml)

### Open Service - Public User API & Web Interface
<span id="lable"></span>

![Image](https://github.com/user-attachments/assets/39c798de-b243-42c1-a75a-cd179913fc49)

- Web Interface: [http://IP:[[config:http-port](config/config.yaml#http-port)]]
- Configuration Settings: [config:database](config/config.yaml#database) and [config:user](config/config.yaml#user)

## Other Resources

- [Obsidian Auto Image Remote Uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)
