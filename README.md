# lucy

lucy 是一个基于 [kratos](https://github.com/go-kratos/kratos) 框架开发的中心化交易所，名字的灵感来源于斯嘉丽约翰逊主演的电影《超体》。
我会在该项目中展示一个中心化交易所的基本构成。

## 注意事项
1. 该项目为一个试验性项目，切勿将其应用于生产环境；
2. 由于工作原因，该项目的更新时间不固定；
3. 该项目的目录结构会参考 [beer-shop](https://github.com/go-kratos/beer-shop)，可能会做一定的修改；
4. 该项目为大仓模式，服务公用同一个 mod。

## 技术选型


* 后端：Go
* 数据库：MySQL
* 缓存：Redis
* 搜索：ElasticSearch
* 服务框架: Kratos
* 网关：
* 消息系统：RabbitMQ
* 服务发现：
* 配置中心
* 服务监控：Prometheus，Grafana
* 服务治理：Hystrix


## 项目结构
### 服务简介
- 用户服务 user
- 钱包服务 wallet
- 订单服务 order
- 活动服务 market
- 交易服务 trade
- 撮合服务 engine
- 开放平台 open-platform


### 系统架构图

### 目录结构


