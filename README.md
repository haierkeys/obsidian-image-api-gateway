# Obsidian Image API Gateway

<p align="center">
    <img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
    <img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway" alt="license">
</p>

This project provides image upload, storage, and cloud synchronization services for the [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian plugin.

## Features List

- [x] Supports image upload
- [x] Supports token authorization to enhance API security
- [x] Supports image HTTP access (basic feature, recommend using Nginx as an alternative)
- [x] Storage support:
  - [x] Save to local or cloud storage simultaneously for convenient migration
  - [x] Local storage support (for NAS, feature has been tested)
  - [x] Support for Alibaba Cloud OSS cloud storage (feature has been implemented, not yet tested)
  - [x] Support for Cloudflare R2 cloud storage (feature has been implemented, tested)
  - [x] Support for Amazon S3 (feature has been implemented, tested)
  - [x] Added MinIO storage support (v1.5+)
  - [x] Support for WebDAV storage (v2.5+)
- [x] Provides Docker installation support for easy use on home NAS or remote servers
- [x] Provides public service API and Web interface for convenient public service provision <a href="#userapi">User Public Interface & Web Interface</a>

## Update Log

For complete update content, please visit [Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

## Pricing

This software is open source and free. If you want to show your appreciation or help support the continuous development, you can provide support for me in the following ways:

[<img src="https://cdn.ko-fi.com/cdn/kofi3.png?v=3" alt="BuyMeACoffee" width="100">](https://ko-fi.com/haierkeys)

## Quick Start
### Installation

- Directory Setup

  ```bash
  # Create the directories required for the project
  mkdir -p /data/image-api
  cd /data/image-api

  mkdir -p ./config && mkdir -p ./storage/logs && mkdir -p ./storage/uploads
  ```

  When starting for the first time, if you don't download the configuration file, the program will automatically generate a default configuration to **config/config.yaml**

  If you want to download a default configuration from the network, use the following command to download:

  ```bash
  # Download the default configuration file from the open source library to the configuration directory
  wget -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml
  ```

- Binary Installation

  Download the latest version from [Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases), unzip and execute:

  ```bash
  ./image-api run -c config/config.yaml
  ```


- Containerized Installation (Docker Method)

  Docker Command:

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

  Docker Compose
  Use *containrrr/watchtower* to monitor the image to implement automatic project updates
  **docker-compose.yaml** content as follows

  ```yaml
  # docker-compose.yaml
  services:
    image-api:
      image: haierkeys/obsidian-image-api-gateway:latest  # Your application image
      container_name: image-api
      ports:
        - "9000:9000"  # Map port 9000
        - "9001:9001"  # Map port 9001
      volumes:
        - /data/image-api/storage/:/api/storage/  # Map storage directory
        - /data/image-api/config/:/api/config/    # Map configuration directory

    watchtower:
      image: containrrr/watchtower
      container_name: watchtower
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock  # Allow Watchtower to access Docker Daemon
      environment:
        - WATCHTOWER_SCHEDULE=0 0,30 * * * *  # Check for updates every half hour
        - WATCHTOWER_CLEANUP=true            # Delete old images, save space
      restart: unless-stopped
  ```

  Execute **docker compose**

  Register docker container as a service

  ```bash
  docker compose up -d
  ```

  Unregister and destroy docker container

  ```bash
  docker compose down
  ```



### Usage

- **Use Single Service Gateway**

  Supports `Local Storage`, `OSS`, `Cloudflare R2`, `Amazon S3`, `MinIO`, `WebDAV`

  Need to modify [config.yaml](config/config.yaml#http-port)

  Modify `http-port` and `auth-token` two options

  Start the gateway program

  API Gateway address is `http://{IP:PORT}/api/upload`

  API access token is `auth-token` content


- **Use Multi-User Open Gateway**

  Supports `Local Storage`, `OSS`, `Cloudflare R2`, `Amazon S3`, `MinIO` (v2.3+), `WebDAV` (v2.5+)

  Need to modify in [config.yaml](config/config.yaml#user)

  `http-port` and `database`

  Also modify `user.is-enable` and `user.register-is-enable` to `true`

  Start the gateway program

  Access the `WebGUI` address `http://{IP:PORT}` for user registration configuration

  ![Image](https://github.com/user-attachments/assets/39c798de-b243-42c1-a75a-cd179913fc49)

  API Gateway address is `http://{IP:PORT}/api/user/upload`

  Click in the `WebGUI` to copy API configuration to obtain configuration information



- **Storage Type Description**


  | Storage Type         | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                |
  |----------------------|------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
  | Server Local Storage | The default save path is: `/data/storage/uploads` Associated configuration item `config.local-fs.save-path` is `storage/uploads`,  <br />If using the gateway image resource access service, `config.local-fs.httpfs-is-enable` needs to be set to `true` <br />The corresponding `Access Address Prefix` is `http://{IP:PORT}`, use single-service gateway to set `config.app.upload-url-pre` <br />Recommended to use Nginx to implement resource access |



### Configuration Description

The default configuration file name is **config.yaml**, please place it in the **root directory** or **config** directory.

For more configuration details, please refer to:

- [config/config.yaml](config/config.yaml)


## Other Resources

- [Obsidian Auto Image Remote Uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)