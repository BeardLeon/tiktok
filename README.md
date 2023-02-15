# tiktok
字节青训营后端项目

# 项目结构

```
tiktok
├── conf # 配置文件
├── controller
├── docs
├── middleware # 应用中间件
│   └── jwt
├── models # 数据库模型
├── pkg # 第三方包
│   ├── app
│   ├── e # API 错误码包
│   ├── export
│   ├── file
│   ├── gredis
│   ├── logging # 日志
│   ├── setting # 配置信息
│   ├── upload
│   └── util # 工具包
├── routers # 路由逻辑处理
├── runtime # 应用运行时数据
├── service
├── sql # sql 文件
├── test
├── main.go # 启动文件
└── READEME.md
```

# 快速上手

1. 创建 tiktok 数据库

    ```mysql
    CREATE DATABASE IF NOT EXISTS tiktok DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
    ```

2. 执行 sql/tiktok.sql 文件

3. 修改 conf/app.ini 文件