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
cp .env.example .env
# 编辑 .env，至少设置 ADMIN_USERNAME、ADMIN_PASSWORD 和 JWT_SECRET
docker compose up -d
```

启动后访问：

```text
http://localhost:51873
```

管理员账号：

```text
用户名：由 .env 中的 ADMIN_USERNAME 决定
密码：由 .env 中的 ADMIN_PASSWORD 决定；留空时首次启动生成随机密码并保存到 ./data/admin-password.txt
```

Compose 通过 `env_file: .env` 把账号、密码、JWT 密钥和数据库配置注入容器；不要使用 `infinite-canvas` 作为密码，容器会拒绝启动。`.env` 已被 `.gitignore` 忽略，公开仓库只提交 `.env.example`。`ADMIN_PASSWORD` 留空时首次启动自动生成随机密码；填写自定义密码后，每次启动都会按 `.env` 同步管理员账号。修改 `.env` 后必须重新创建容器：

```bash
docker compose up -d --force-recreate
```

仅执行 `docker restart` 不会读取新的环境变量；在飞牛 fnOS 的 Docker Compose 项目里，修改保存后应执行「构建/重新创建」。`ADMIN_PASSWORD` 非空时它是管理员密码的唯一来源，后台修改的密码会在下次按 Compose 重启时被覆盖；留空时系统不会覆盖后台修改的密码。数据目录 `./data` 会保留。容器内前端会自动代理到同一容器的 `127.0.0.1:8080` 后端，登录接口不再依赖容器外部的 `PORT` 环境变量。`JWT_SECRET` 留空时自动生成并保存到 `./data/jwt-secret.txt`，重启后登录状态保持有效。

## 本地构建镜像

如果需要基于当前源码构建镜像，同样先准备 `.env`：

```bash
cp .env.example .env
# 编辑 .env
docker compose -f docker-compose.local.yml up -d --build
```

修改 `.env` 后同样需要重新创建容器：

```bash
docker compose -f docker-compose.local.yml up -d --force-recreate
```

## 数据目录

`docker-compose.yml` 会把本地 `./data` 挂载到容器内 `/app/data`，用于保存 SQLite 数据库、提示词数据和上传素材。

容器内数据库路径由 Compose 固定为：

```text
DATABASE_DSN=/app/data/infinite-canvas.db
```

`.env` 中的 `DATABASE_DSN=data/infinite-canvas.db` 主要用于本地 `go run .`；Docker Compose 会在容器 environment 中覆盖为上述绝对路径。如需调整，修改 `.env` 后重新创建容器；若同时保留 Compose 的 `environment.DATABASE_DSN`，以 Compose 中的值为准。
