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

管理员账号：

```text
用户名：由 ADMIN_USERNAME 决定，默认 admin
密码：由 ADMIN_PASSWORD 决定；留空时首次启动生成随机密码并保存到 ./data/admin-password.txt
```

账号、密码和数据库配置写在 `docker-compose.yml` 的 `environment` 中，无需创建 `.env`。`ADMIN_PASSWORD` 留空时首次启动自动生成随机密码；填写自定义密码后，每次启动都会按该值同步管理员账号。修改账号密码后必须重新创建容器：

```bash
docker compose up -d --force-recreate
```

仅执行 `docker restart` 不会读取新的 `environment`；在飞牛 fnOS 的 Docker Compose 项目里，修改保存后应执行「构建/重新创建」。`ADMIN_PASSWORD` 非空时它是管理员密码的唯一来源，后台修改的密码会在下次按 Compose 重启时被覆盖；留空时系统不会覆盖后台修改的密码。数据目录 `./data` 会保留。容器内前端会自动代理到同一容器的 `127.0.0.1:8080` 后端，登录接口不再依赖容器外部的 `PORT` 环境变量。`JWT_SECRET` 留空时自动生成并保存到 `./data/jwt-secret.txt`，重启后登录状态保持有效。

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
