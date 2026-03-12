# 本地开发指南

本文档介绍如何在本地运行 MatrixHub 的前端和后端服务。

## 前置要求

- Go 1.23+
- Node.js 18+
- pnpm 8+
- Docker & Docker Compose（推荐）

## 快速开始

### 本地开发

如果需要本地修改代码并调试，可以手动启动服务。

#### 1. 启动 MySQL 数据库

```bash
docker run -d \
  --name matrixhub-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=matrixhub \
  -e MYSQL_USER=matrixhub \
  -e MYSQL_PASSWORD=changeme \
  -p 3306:3306 \
  mysql:8.4
```

#### 2. 配置环境变量

```bash
export MATRIXHUB_DATABASE_DSN="matrixhub:changeme@tcp(127.0.0.1:3306)/matrixhub?charset=utf8mb4&multiStatements=true&parseTime=true"
```

#### 3. 启动后端服务

```bash
# 使用默认配置文件
go run ./cmd/matrixhub apiserver

# 或指定配置文件
go run ./cmd/matrixhub apiserver -c config/dev-config.yaml
```

#### 4. 启动前端服务

```bash
cd ui
pnpm install  # 首次运行需要安装依赖
pnpm dev
```

前端开发服务器将在 http://localhost:5173 启动，并自动代理 API 请求到后端。

## 使用 Makefile

项目提供了便捷的 Makefile 命令：

**⚠️ 重要**: 使用 `local-run` 和 `local-run-api` 命令前，**必须先启动 MySQL 数据库**。

```bash
# 1. 启动 MySQL 数据库（如果还没有启动）
docker run -d \
  --name matrixhub-mysql \
  -e MYSQL_ROOT_PASSWORD=password \
  -e MYSQL_DATABASE=matrixhub \
  -e MYSQL_USER=matrixhub \
  -e MYSQL_PASSWORD=changeme \
  -p 3306:3306 \
  mysql:8.4

# 2. 本地运行（前端 + 后端）
make local-run

# 或分别运行：
make local-run-api   # 只运行后端 API 服务器
make local-run-web   # 只运行前端
```

**提示**:
- `local-run-web` 不依赖 MySQL，可独立运行
- `local-run-api` 和 `local-run` 需要 MySQL 运行
- 如果没有配置环境变量，确保 `config/config.yaml` 中的数据库 DSN 正确

## 配置说明

### 环境变量

`config.yaml` 支持通过环境变量配置数据库连接：

```bash
export MATRIXHUB_DATABASE_DSN="matrixhub:changeme@tcp(127.0.0.1:3306)/matrixhub?charset=utf8mb4&multiStatements=true&parseTime=true"
```

### 前端代理配置

前端已配置 Vite 代理，开发环境下会自动将 `/api/*` 和 `/apis/*` 请求代理到后端服务器（http://127.0.0.1:3001），无需额外配置。

配置文件：`ui/vite.config.ts`

## 访问应用

- **前端 UI**: http://localhost:5173
- **后端 API**: http://localhost:3001
- **API 健康检查**: http://localhost:3001/health

## 开发提示

### 数据库

1. **首次启动**: 设置 `database.migrate: true` 会自动创建数据库表
2. **调试模式**: 设置 `debug: true` 可以看到详细的 SQL 日志
3. **数据持久化**: Docker Compose 使用命名卷，数据在容器重启后保留

### 前端

1. **热重载**: 修改代码会自动刷新浏览器
2. **API 代理**: 开发环境下无需处理 CORS
3. **TypeScript**: 使用 `pnpm typecheck` 进行类型检查

### 后端

1. **代码修改**: 需要手动重启服务
2. **依赖管理**: 使用 `go mod tidy` 整理依赖
3. **配置验证**: 启动时会自动验证配置文件

## 故障排查

### MySQL 连接失败

```bash
# 检查 MySQL 容器状态
docker ps | grep matrixhub-mysql

# 查看 MySQL 日志
docker logs matrixhub-mysql

# 重启 MySQL
docker restart matrixhub-mysql
```

### 端口冲突

如果端口被占用，可以修改：

**后端端口**: 

修改 `config.yaml` 中的 `apiServer.port`

**前端端口**:
```bash
cd ui
pnpm dev --port 3000
```

### 依赖问题

**Go 依赖**:
```bash
go mod tidy
go mod download
```

**前端依赖**:
```bash
cd ui
rm -rf node_modules pnpm-lock.yaml
pnpm install
```
