简介
这是一个用Golang实现用户注册登录功能的demo
功能要求
• 打开localhost:3000/signup，填写注册信息
• 用户名(验证要求：手机号码或邮箱)
• 密码(拥有大小写字母及数字，至少8位)
• 打开localhost:3000/signin，显示登录框(可填写用户名和密码)，登录框下方有按钮可以调到signup页面注册
• 若登录成功 -> 跳转localhost:3000/profile，显示用户名及Login Succeeded
• 若登录失败 -> 停留当前页面，显示Login Failed
• 登录状态保持在前端，未登录状态打开localhost:3000/profile，跳转至localhost:3000/signin
模块需求
• 任选一种web framework（https://github.com/mingrammer/go-web-framework-stars）
• 数据库使用postgres
• 登录状态缓存使用redis
项目开发
使用的技术和服务框架
使用Go语言，基于Beego框架进行开发，使用Session+Cookie进行登录状态认证，Session在服务器端使用Redis保存在内存中，账户数据通过Postgres进行持久化保存。
进度
目前功能要求都已经在本地实现，但是还存在一些细枝末节的小问题，如果clone到其他的电脑上，也无法直接运行，像数据库之类的设置都是前期手动创建调试的，接下来会对这方面进一步优化。
未完待续。。。