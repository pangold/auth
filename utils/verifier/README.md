## 验证码

#### 验证码的生成

调用 utils.GenerateRandomNumber(4) 来生成4位的验证码

#### 调用 Content 生成内容

调用 content.GetVerificationText(code) 来生成 短信/邮件 内容 

> content 有点多余

#### SendEmail 发送邮件（通过邮件服务发送）

暂未选型

#### SendSMS 发送短信（通过第三方SMS服务发送）

暂未选型