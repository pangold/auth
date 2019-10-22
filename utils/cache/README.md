## 缓存

公共的工具类

#### 接口 

~~~~
type Cache interface {
	HasCacheKey(service, action, key string) bool
	GetCacheValue(service, action, key string, value *string) error
	SetCacheValue(service, action, key, value string, timeout int) error
}
~~~~

#### SimpleCache

简单的使用了 `map[string]interface{}` 来存储数据。优点是实现简单，非常适合在开发环境中使用。当然缺点也非常明显，以普通变量的方式存放在内存栈中的数据，当进程关闭后，数据将会丢失。

#### RedisCache

使用 Redis 来存储数据，没有 SimpleCache 那么简单的使用 map 来实现数据的缓存，适合在生产环境中使用，不存在数据丢失的问题。

其调用参数中 Key 值相关的有3个参数，分别为 service、action、key，可灵活的更换其他 Redis 命令进行功能的实现，如：

* 整合 service.action.key 作为 String 类型的 Key，使用 Set & Get 两个命令设置/获取 Key 的数据
* 使用 service 作为 Hash 类型的 Key，然后用 action.key 作为 Field，使用 HSet & HGet 两个命令来获取 Key 的数据

当然还可以衍生出很多种，但是以上已经足够实现当前的功能，无需搞的太复杂。

至于以上两种方案中，哪种更优？暂未测试，这里只强调扩展性。

> 使用 Redis 的话，其实2个参数足够了。


