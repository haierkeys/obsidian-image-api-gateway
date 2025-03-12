[中文文档](readme-zh.md) / [English Document](README.md)
# Obsidian Image API Gateway

<p align="center">
    <img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
    <img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway" alt="license">
</p>

This project provides image uploading, storage, and cloud synchronization services for the [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian plugin.

## Feature List

- [x] Supports image uploading
- [x] Supports token authorization to enhance API security
- [x] Supports HTTP access to images (basic functionality, Nginx is recommended as an alternative)
- [x] Storage support:
  - [x] Save to both local and cloud storage for easy migration
  - [x] Local storage support (designed for NAS, functionality tested)
  - [x] Support for Alibaba Cloud OSS storage (implemented but untested)
  - [x] Support for Cloudflare R2 storage (implemented and tested)
  - [x] Support for Amazon S3 (implemented and tested)
  - [x] Added MinIO storage support (v1.5+)
  - [ ] Google ECS support (to be developed)
- [x] Provides Docker installation support for use on home NAS or remote servers
- [x] Provides public service API & web interface for public service usage <a href="#userapi">User Public Interface & Web Interface</a>

## Changelog

To view the complete update details, please visit [Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

## Pricing

This software is open-source and free. If you would like to show your appreciation or support continued development, you can do so via the following method:

[<img src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" alt="BuyMeACoffee" width="100">](https://ko-fi.com/haierkeys)

## Quick Start
### Installation

- Directory Setup

  ```bash
  # Create the required directories for the project
  mkdir -p /data/image-api
  cd /data/image-api

  mkdir -p ./config && mkdir -p ./storage/logs && mkdir -p ./storage/uploads
  ```

  On first startup, if the configuration file is not downloaded, the program will automatically generate a default configuration in `config/config.yaml`.

  If you want to download the default configuration from the network, use the following command:

  ```bash
  # Download the default configuration file from the open-source repository to the config directory
  wget -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml
  ```




- Containerized Installation (Docker Method)

  Assuming your server stores images at _/data/storage/uploads_, execute the following commands:

  ```bash
  # Pull the latest container image
  docker pull haierkeys/obsidian-image-api-gateway:latest

  # Create and start the container
  docker run -tid --name image-api \
          -p 9000:9000 -p 9001:9001 \
          -v /data/image-api/storage/:/api/storage/ \
          -v /data/image-api/config/:/api/config/ \
          haierkeys/obsidian-image-api-gateway:latest
  ```

- Binary Installation

  Download the latest version from [Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases), extract it, and execute:

  ```bash
  ./image-api run -c config/config.yaml
  ```

### Usage

- Using the Single-Service API

    Supports `Local Storage`, `OSS`, `Cloudflare R2`, `Amazon S3`, and `MinIO`.

    Modify [config.yaml](config/config.yaml#http-port)

    Change `http-port` and `auth-token` options.

    Start the gateway program.

    The API gateway address is `http://{IP:PORT}/api/upload`

    The API access token is the content of `auth-token`.

- Using the Multi-User Public Service API

    Supports `OSS`, `Cloudflare R2`, and `Amazon S3`.

    Modify [config.yaml](config/config.yaml#user)

    Change `http-port` and `database`.

    Also, set `user.is-enable` and `user.register-is-enable` to `true`.

    Start the gateway program.

    Access the `WebGUI` at `http://{IP:PORT}` to configure user registration.

    ![Image](https://github.com/user-attachments/assets/39c798de-b243-42c1-a75a-cd179913fc49)

    The API gateway address is `http://{IP:PORT}/api/user/upload`.

    Click on `WebGUI` to copy API configuration and obtain setup details.

### Configuration Instructions

The default configuration file is named _config.yaml_. Place it in the _root directory_ or the _config_ directory.

For more configuration details, refer to:

- [config/config.yaml](config/config.yaml)
## Other Resources

- [Obsidian Auto Image Remote Uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)

