---
title: Docker 部署
description: 使用 Docker Compose 部署无限画布
---

# Docker 部署

如果你希望在自己的机器或服务器上运行项目，可以直接使用 Docker Compose。

## 使用发布镜像

```bash
git clone https://github.com/mingtian886/zou-canvas.git
cd zou-canvas
docker compose up -d
```

启动后访问：

```text
http://localhost:51873
```

默认管理员账号：

```text
用户名：admin
密码：infinite-canvas
```

账号、密码和数据库配置已直接写在 `docker-compose.yml` 的 `environment` 中，无需创建 `.env`；需要修改时直接编辑该文件，再执行 `docker compose up -d`。

## 本地构建镜像

如果需要基于当前源码构建镜像：

```bash
docker compose -f docker-compose.local.yml up -d --build
```

## 数据目录

`docker-compose.yml` 会把本地 `./data` 挂载到容器内 `/app/data`，用于保存 SQLite 数据库、提示词数据和上传素材。

容器内数据库路径已内置为：

```text
DATABASE_DSN=/app/data/infinite-canvas.db
```

如需调整，直接修改 `docker-compose.yml` 中的 `environment` 配置。
