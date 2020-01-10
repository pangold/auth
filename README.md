#### Auth (Account)

##### 配置 -- 策略：

身份验证策略（注册的策略）：

1. 无需身份验证（会被恶意注册）
2. 验证码（恶意注册风险最低）-- 推荐
3. URL验证（也会被恶意注册，虽然账号是未激活状态）

其他：

1. 第三方验证（不受上面3种策略限制）
2. 锁定/解锁机制（当密码错误次数到达阀值，手机被盗、邮箱被盗时自主申请锁定 / 解锁手机或邮箱）
3. 解除绑定（手机或邮箱）
4. 其他安全检测机制（后续增加的特性 -- 一般系统运行的很成功，用户量很大时才需要这些）

1. 简单：
   注册：用户名 + 密码 + 密码确认，无需任何验证（其他选填）         SignUp
   登录：用户名 + 密码 （可选：验证码登录（绑定邮箱或手机后可用））   UserName 
   忘记：联系管理员（管理员忘记密码就联系运维或开发）
   注销：清除 token
   ---------------
   可以通过绑定邮箱或手机号，享受验证码登录、通过验证码修改密码
   
2. 邮件链接（因为这个策略，需要多做很多功夫啊）
   注册：邮箱 + 密码 + 密码确认，需要接收邮件激活账号（其他选填）    SignUpWithActivate
   登录：邮箱 + 密码（可选：验证码登录）
   忘记：email -> 发送链接，点击链接重置密码                      Forgot -> ResetPasswordByHashCode
   注销：清除 token
   ---------------
   可以通过绑定手机号，可额外享受：验证码登录、通过验证码修改密码
   
3. 验证码：
   注册：邮箱/手机 + 验证码 + 密码 + 密码确认（其他选填）          SignUpWithVCode
   登录：邮箱/手机 + 密码 / 验证码                              SignInWithVCode
   忘记：邮箱/手机 + 验证码 + 新密码 + 密码确认                   ResetPassword
   注销：清除 token
   
设计理念：只要有邮箱或手机号，都可以用验证码

通过配置使用不同的策略，不同策略下，使用相同的接口，后台自己判断是用户名还是邮箱还是手机号，这样后台的适应性更强，前端再根据不同的需求做不同的页面。

##### 注册

包括字段：用户名、邮箱、手机号、手机/邮箱验证码、密码、确认密码，组合分3种：

* 用户名 + 密码
* 用户名 + 邮箱 + 密码 -> 需要邮箱激活账号
* 用户名 + 邮箱 + 验证吗 + 密码
* 用户名 + 手机 + 验证码 + 密码
* 第一次第三方登录时注册

上面各种注册方案中提到的字段为必填，其余为选添

##### 登录

使用用户名、手机号、Email作为账号登录，对应不同的注册4种方式，不同的账号登录逻辑不同：

* 使用用户名 + 密码登录（适合1 2 3 4）
* 使用邮箱 + 密码登录（适合2 3，和绑定邮箱的1 4）
* 使用邮箱 + 验证码登录（适合2 3，和绑定邮箱的1 4）
* 使用手机 + 密码登录（适合4，和绑定手机的1 2 3）
* 使用手机 + 验证码登录（适合4，和绑定手机号吗的1 2 3）
* 使用第三方登录

UI 只提供

* 账号、密码登录（默认）-- 账号可以是：用户名、邮箱、手机号，自动识别类型（可以在 POST 之前由 JS 识别；在后台识别 -- 这样的话接口简单点）
* 账号、验证码登录     -- 账号只能是：邮箱、手机号，自动识别

邮箱格式很容易识别，全世界都统一格式，
手机号码：根据地区识别，中国大陆的手机号也固定，
手机服务需要第三方短信服务，

**第三方登录**

1. 微信登录
2. QQ 登录
3. 新浪微博登录
4. Github 登录
5. Twitter 登录
6. Facebook 登录

##### 忘记密码

* 没有绑定邮箱、手机号的
* 使用邮箱 -> 发送链接到邮箱 -> 点击链接打开修改页面 -> 修改密码完成（适合2 3，和绑定邮箱的1 4）
* 使用邮箱 + 邮箱验证码 + 新密码 -> 修改密码完成（适合2 3，和绑定邮箱的1 4）
* 使用手机 + 手机验证码 + 新密码 -> 修改密码完成（适合4，和绑定手机号的1 2 3）

第三方登录无需忘记密码

##### 修改密码

必须先登录

##### API

* **注册**：包括字段
* **登录**：包括字段（参照#####登录提到的）
* **注销**：token （token分类：jwt、自定义 HashCode）
* 忘记密码

##### 其他 API

* 重置密码 HashCode -> 邮件激活/邮件重置密码
* 是否存在：字段指定，
* 获取激活码
* 绑定手机号
* 绑定邮箱

