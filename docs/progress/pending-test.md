---
title: 待测试
description: 当前版本已实现但仍需人工验证的变更项
---

# 待测试

- 确认首页「站点公告」栏目不再显示。
- 确认首页轮播大图仅保留 4 张无推广链接内容，且两张推广视频/图片不再显示。
- 确认模型配置中的推广「获取 API Key」链接已移除，TokenDance 登录功能仍可用。
- 确认 Docker 部署监听非常用端口 `51873`，访问 `http://localhost:51873` 正常。
- 确认顶部导航不再显示「音乐创作台」入口，避免进入未实现的 /music 页面。
- 确认顶部操作区与 3D 导演台内都不再显示 GitHub 图标及其跳转链接。
- 确认首页顶部「无限画布」品牌与导航菜单整体居中显示，右侧操作区仍靠右，其它页面顶部布局不变。
- 确认 `docker-compose.yml` 和 `docker-compose.local.yml` 通过 `env_file: .env` 读取管理员账号、密码和 JWT 配置，公开仓库只提交 `.env.example`。
- 确认 README、快速开始和 Docker 部署文档已说明 `cp .env.example .env`、编辑 `.env` 后启动，以及修改后需 `docker compose up -d --force-recreate`；`docker-compose.local.yml` 同样读取 `.env`。
- 确认 GitHub 仓库、GHCR 镜像和容器名统一改为 `zou-canvas`，GHCR 地址为 `ghcr.io/mingtian886/zou-canvas:latest`。
- 确认 README 标题、徽标和克隆地址已统一指向 `zou-canvas`，且 `web/bun.lock` 的包名已同步。
- 确认登录页仅保留账号密码登录，后台设置与用户列表不再显示已停用的第三方账号登录配置项。
- 确认页面顶部不再显示版本号及版本更新弹窗，也不再请求上游版本信息。
- 确认 Docker 容器内登录接口走 `127.0.0.1:8080` 后端代理，不再受前端 `PORT=51873` 影响。
- 确认 `.env` 被 `.gitignore` 忽略，公开仓库不包含真实管理员密码或 JWT 密钥；`ADMIN_PASSWORD` 留空时首次启动生成随机密码并写入 `./data/admin-password.txt`。
- 确认在 `.env` 中修改 `ADMIN_USERNAME` / `ADMIN_PASSWORD` 后执行 `docker compose up -d --force-recreate` 会同步生效，且 `./data` 数据保留。
- 确认 `.env` 中的 `ADMIN_PASSWORD` 仍填写 `infinite-canvas` 时容器拒绝启动，避免继续使用示例密码。
- 确认 `.env` 中的 `JWT_SECRET` 留空时自动生成并写入 `./data/jwt-secret.txt`，容器重启后旧登录 token 仍有效。
- 确认 NAS 部署已按 `.env` 重建，旧默认账号密码失效、新账号密码可登录，且数据目录保留。
