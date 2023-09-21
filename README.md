# MicroShop

MicroShop 是一个基于 Go 语言和一系列微服务技术构建的微型电子商务平台。该项目使用了以下技术栈：

- Go 语言
- go-micro 微服务框架
- Consul 注册中心
- Jaeger 链路追踪
- Prometheus 监控
- Docker 容器化
- Docker Compose 多容器管理
- ELK 堆栈（Elasticsearch、Logstash 和 Kibana）

## 项目概述

MicroShop 旨在提供一个灵活、可扩展的电子商务平台，包括以下主要模块：

- **用户模块**: 处理用户身份验证和管理。
- **产品模块**: 管理产品信息和库存。
- **付款模块**: 处理付款和订单支付。
- **订单模块**: 创建、管理和跟踪订单。
- **分类模块**: 管理产品分类信息。
- **购物车模块**: 允许用户管理购物车。
- **API 模块**: 包括购物车和付款的API接口。