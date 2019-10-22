## 令牌

#### 接口

~~~~
type Token interface {
	GenerateToken(userId, userName string) string
	ExplainToken(token string, userId, userName *string) error
}
~~~~

#### MyToken

使用自己的方式来实现 Token 的生成与验证，本人结合 Redis 来存储 Token，实现的原理如下：

##### 生成 Token 的流程：

1. 调用 utils 的 Hash 生成遗传 64 位的随机字符串（长度可以设置为任意值）
2. 调用 Redis 的 `SET` 命令，以 `cache.` 为前缀，紧接上面的 64 位的随机字符串为 Key，以 `UserId` 为 Value，保存到 Redis 中
3. 调用 Redis 的 `EXPIRE` 命令，健值的过期时间（即 Token 的过期时间）

##### 校验 Token 的流程：

1. 拼接 `cache.` 和 `token` 为 Key
2. 调用 Redis 的 `GET` 命令查询缓存中的数据（UserId），此时可能会存储两种情况：
	1. Key 存在，验证成功，返回数据 UserId 用作业务上的处理
	2. Key 不存在，验证失败，无效 Token 或 Token 已过期

> 不做 `HasToken` 这个接口，直接调用 `ExplainToken` 就好，效果，性能都是一样的，没必要多此一举。

#### JwtToken

相比上面 MyToken，使用 Jwt 实现 Token 则不需要存储的操作了，因为 Token 自身就包含了过期，UserId 等信息了，这样当然有好处也有坏处：

**优点**

* 不需要依赖外部的服务（用于存储的服务，数据库，NoSQL，其他分布式存储如 ZK，ETCD 等，统统不需要了）
* 部署方便
* 单点登录

**缺点**

* Token 一经设置，就不能更改失效时间
* 更改密码，旧 Token 仍然有效
* 注销后，Token 仍然有效



