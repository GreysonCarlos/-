# Topic-web
- 展示话题(标题，文字描述)和回帖列表
- 暂不考虑前端页面实现，仅仅实现一个本地web服务
- 话题和回帖数据用文件存储
## 需求用例
用户浏览页面两要素：
1. 可以浏览话题topic
    - 定义结构体（根据ER图来进行定义）：
        - id、title、content、create_time
2. 可以浏览话题的回帖post
    - 定义结构体（根据ER图来进行定义）：
        - id、topic_id、content、create_time

## 分层结构
- 数据层Repository,数据的增删改查
- 逻辑层Service,处理核心业务逻辑输出
- 视图层Controller,处理和外部的交互逻辑

## 组件工具
- Gin 高性能go web框架
    - https://github.com/gin-gonic/gin#installation
- Go Mod