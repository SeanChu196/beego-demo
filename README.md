# [Beego-demo](https://github.com/SeanChu196/beego-demo)
##  项目简介
这是一个基于beego框架，用Golang实现用户注册登录功能的demo。使用Session+Cookie进行登录状态认证，Session在服务器端使用Redis保存在内存中，Cookie缓存在浏览器，用户信息使用Postgres进行持久化保存。

- [controllers] 包含beego控制器，根据HTTP请求的方法负责不同的逻辑处理

    - [signup.go] 用户注册页面的控制器，会调用[models]检测用户的输入是否合法并存取数据库
    - [signup.go] 用户登录界面控制器，不仅检测用户输入合法性，还会设置session信息
    - [profile.go] 用户简况页面控制器，会检查session信息。
    
- [models] 包含models.go，负责postgres的注册连接和读写，还提供字符串的正则匹配以验证用户和密码的格式完整性
- [views] 包含HTML模板
- [scripts] 包含初始化脚本以创建postgres数据库和开启Redis服务器
- [conf] 包含项目配置文件，存放了beego框架session和数据库相关参数

## 环境
Mac OS 10.15.5

go 1.14

redis 6.05

postgresql 12.3

beego 1.12.2

## 安装使用
- 使用`git clone https://github.com/SeanChu196/beego-demo.git` 拷贝项目到本地

- 运行`$ bash scripts/init.sh` 启动Redis服务并创建数据库

- 使用`go run .` 或者`go run main.go` 启动服务程序

- 访问`localhost:3000/signin` 进行注册或登录

## 功能说明

- 打开`localhost:3000/signup`，填写注册信息

    - 用户名(验证要求：手机号码或邮箱)
    - 密码(拥有大小写字母及数字，至少8位)
- 打开`localhost:3000/signin`，显示登录框(可填写用户名和密码)，登录框下方有按钮可以调到signup页面注册

- 若登录成功 -> 跳转`localhost:3000/profile`，显示用户名及Login Succeeded

- 若登录失败 -> 停留当前页面，显示Login Failed

- 登录状态保持在前端，未登录状态打开`localhost:3000/profile`，跳转至`localhost:3000/signin`

- 模块需求

    - 任选一种[web framework](https://github.com/mingrammer/go-web-framework-stars)

    - 数据库使用postgres

    - 登录状态缓存使用redis

## 项目状态
目前功能要求都已经实现，Redis开启服务器和Postgres数据库初始化都已经写成脚本，session管理进行了重写，各种参数都使用了配置文件，使用起来更加灵活方便。

## 配置参数和脚本说明
#### app.conf

- [appname] 应用名称

- [httpport] 服务器端口

- [runmode] 运行模式

- [sessionon] 是否启用session

- [sessionprovider] session引擎默认是memory，目前支持还有file、mysql、redis等

- [sessionname] cookies中的名字

- [sessiongcmaxlifetime] Session 过期的时间，默认值是 3600 秒

- [sessionproviderconfig] 对应file、mysql、redis引擎的保存路径或者链接地址，默认值是空

- [pgusr] postgres用户名

- [pgpassword] postgres密码

- [pgdbname] postgres数据库名称，必须和init.sh中创建的数据库名称一致，且不可以使用大写字母，因为脚本创建数据库时postgres会自动转换为小写字母作为数据库名

- [pghost] postgres连接地址

- [pgport] postgres连接端口

- [pgsslmode] 是否启用SSL协议

