#### 组件

这部分组件相对与其他 Helper 工具类相比，会更复杂一些，并且依赖其他组件，潜在变动的风险会更高，因此对具体的实现作独立的封装（详细实现参见具体目录）。

* **缓存** Cache
* **内容生成** Content
* **令牌** Token
* **验证码** Verifier

#### 其他 Helper

* **Redis Helper**: 连接、关闭、Set、Get 等 NoSQL 的通用接口
* **Database Helper**: 数据库连接（仅仅为连接的接口，当前只实现了 MySQL 的连接）
* **Hash Helper**: HashCode 生成器 & MD5 加密