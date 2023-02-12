# tiktok
字节青训营后端项目

# 项目结构

```
tiktok
├── conf # 配置文件
│   └── app.ini # 基础应用配置
├── controller
│   ├── comment.go
│   ├── common.go
│   ├── demo_data.go
│   ├── favorite.go
│   ├── feed.go
│   ├── message.go
│   ├── publish.go
│   ├── relation.go
│   └── user.go
├── docs
│   ├── docs.go
│   ├── swagger.json
│   └── swagger.yaml
├── middleware # 应用中间件
│   ├── jwt
│       └── jwt.go
├── models # 应用数据库模型
│   ├── article.go
│   ├── auth.go
│   ├── models.go # models 的初始化
│   └── tag.go
├── pkg # 第三方包
│   ├── e # API 错误码包
│   │   ├── code.go
│   │   └── msg.go
│   ├── logging # 日志
│   │   ├── file.go
│   │   └── log.go
│   ├── setting
│   │   └── setting.go # 读取配置文件并初始化
│   └── util # 工具包
│       └── pagination.go # 获取分页页码
├── routers # 路由逻辑处理
│   ├── api
│   │   ├── auth.go
│   │   ├── upload.go
│   │   └── v1
│   │       ├── article.go
│   │       └── tag.go
│   └── router.go
├── runtime # 应用运行时数据
├── service
├── sql
├── test
├── go.mod
├── main.go # 启动文件
└── READEME.md
```