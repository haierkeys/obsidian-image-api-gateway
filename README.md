[中文文档](readme-zh.md) / [English Document](README.md)
# Obsidian Image API Gateway

<p align="center">
    <img src="https://img.shields.io/github/release/haierkeys/obsidian-image-api-gateway" alt="version">
    <img src="https://img.shields.io/github/license/haierkeys/obsidian-image-api-gateway" alt="license">
</p>

This project provides image upload, storage, and cloud synchronization services for the [Custom Image Auto Uploader](https://github.com/haierkeys/obsidian-custom-image-auto-uploader) Obsidian plugin.

## Feature List

- [x] Supports image upload
- [x] Supports token authorization to enhance API security
- [x] Supports HTTP access to images (basic feature, recommended to use Nginx instead)
- [x] Storage support:
  - [x] Simultaneously save to local or cloud storage for easy migration
  - [x] Local storage support (prepared for NAS, functionality tested)
  - [x] Supports Alibaba Cloud OSS storage (implemented, not yet tested)
  - [x] Supports Cloudflare R2 storage (implemented and tested)
  - [x] Supports Amazon S3 (implemented and tested)
  - [x] Added MinIO storage support (v1.5+)
  - [ ] Supports Google ECS (to be developed)
- [x] Provides Docker installation support for easy use on home NAS or remote servers
- [x] Provides public service API && Web interface for convenient public service <a href="#userapi">User Public Interface & Web Interface</a>

## Changelog

For a complete list of updates, please visit [Changelog](https://github.com/haierkeys/obsidian-image-api-gateway/releases).

## Pricing

This software is open-source and free. If you’d like to show your appreciation or help support ongoing development, you can support me through the following method:

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

  The default local image storage path on the server is **/data/storage/uploads**

  If you don’t download the configuration file on the first startup, the program will automatically generate a default configuration at **config/config.yaml**

  If you want to download a default configuration from the web, use the following command:

  ```bash
  # Download the default configuration file from the open-source repository to the config directory
  wget -P ./config/ https://raw.githubusercontent.com/haierkeys/obsidian-image-api-gateway/main/config/config.yaml
  ```

- Binary Installation

  Download the latest version from [Releases](https://github.com/haierkeys/obsidian-image-api-gateway/releases), unzip it, and run:

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
  Use *containrrr/watchtower* to monitor the image for automatic project updates
  The content of **docker-compose.yaml** is as follows:

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
        - /data/image-api/config/:/api/config/    # Map config directory

    watchtower:
      image: containrrr/watchtower
      container_name: watchtower
      volumes:
        - /var/run/docker.sock:/var/run/docker.sock  # Allow Watchtower to access Docker Daemon
      environment:
        - WATCHTOWER_SCHEDULE=0 0,30 * * * *  # Check for updates every half hour
        - WATCHTOWER_CLEANUP=true            # Delete old images to save space
      restart: unless-stopped
  ```

  Execute **docker compose**

  Register the Docker container as a service:

  ```bash
  docker compose up -d
  ```

  Stop and destroy the Docker container:

  ```bash
  docker compose down
  ```

### Usage

- Using Single-Service Interface

  Supports **local storage**, **OSS**, **Cloudflare R2**, **Amazon S3**, **MinIO**

  You need to modify [config.yaml](config/config.yaml#http-port)

  Modify the `http-port` and `auth-token` options

  Start the gateway program

  The API gateway address is `http://{IP:PORT}/api/upload`

  The API access token is the content of `auth-token`

- Using **Multi-User** Public Service Interface

  Supports **OSS**, **Cloudflare R2**, **Amazon S3**

  You need to modify [config.yaml](config/config.yaml#user)

  Modify `http-port` and `database`

  Also set `user.is-enable` and `user.register-is-enable` to `true`

  Start the gateway program

  Visit the **WebGUI** address `http://{IP:PORT}` for user registration and configuration

  ![Image](https://github.com/user-attachments/assets/39c798de-b243-42c1-a75a-cd179913fc49)

  The API gateway address is `http://{IP:PORT}/api/user/upload`

  Click on **WebGUI** to copy the API configuration and retrieve the configuration information

### Configuration Details

The default configuration file is named **config.yaml**, please place it in the **root directory** or **config** directory.

For more configuration details, refer to:

- [config/config.yaml](config/config.yaml)

## Other Resources

- [Obsidian Auto Image Remote Uploader](https://github.com/haierkeys/obsidian-auto-image-remote-uploader)

