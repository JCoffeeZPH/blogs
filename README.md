## 1、说明
前端基于：[aurora](https://github.com/linhaojun857/aurora)

参考文档：

[gozero快速构建微服务](https://github.com/zeromicro/zero-doc/blob/main/docs/zero/bookstore.md)
[gorm使用文档](https://gorm.io/docs/advanced_query.html)

## 2、说明
### 2.1 缓存
* 1、redis db0 -> session信息
* 2、redis db1 -> 访客、website相关

redis使用详情

| db  | key                             | data structure | description         |
|-----|---------------------------------|----------------|---------------------|
| 0   | blob.user_session.{sessionId}   | string         | 用户session信息         |
| 1   | blog.unique.visitors            | set            | 访客唯一标识，用于去除重复访问(UV) |
| 1   | blog.view.count                 | string         | blog访问量             |
| 1   | blog.visitors.area              | hash           | 地区访客数量              |
| 1   | blog.website.config.2           | string         | 默认website 配置        |