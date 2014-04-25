## 安装
```
    ./install
```

## 启动
```
    ./bin/web -conf=etc/conf.ini
```

## 访问
```
    localhost:8080/test
```

## 使用
1、路由
```
a) 见src/web/main.go，gear.AddRoute("/info", &InfoApp{})
b) 路由有两级，/module/action
c) 由InfoApp接管/info的路由，module为info，自动检测InfoApp内XxxAction函数，例：/info/detail会路由到InfoApp.DetailAction
d) InfoApp必须定义IndexAction，作为没有action时默认执行方法，/info == /info/index
```

2、gear/app.go
```
a) 继承自request，response，实现appInerface
b) 业务app需继承gear.App，见src/web/app/infoApp.go
```

3、gear/request.go
```
a) 封装 *http.Request
b) 实现 GetXxx/PostXxx/Ip/Ip2Long 等基于*http.Request的方
```

4、gear/response.go
```
a) 封装 http.ResponseWriter
b) 实现 Display/Json/Redirect 等基于http.ResponseWriter方法
```

5、gear/check.go
```
a) 实现在App业务执行之前做公共验证，ip/auth/login等等，定义接口CheckInterface
b) 实现此接口需要实现Check，Failed 两函数
c) app实现CheckInterface后，框架在执行action前会调用Check，false失败，执行Failed，true成功，执行相应路由action
d) 针对类似登录检测等，可实现一个公共CheckLogin，对需要做检测的app则继承CheckLogin
e) 针对单个app 特殊验证，直接在app内实现Check，Failed
```

6、gear/func.go
```
a) 公共无其他依赖的基础方法
```

7、gear/gear.go
```
a) 框架入口文件，用来启动服务
```

8、conf.ini
>addr=":8080"                   # 服务监听地址
>r_timeout=10                   # 读超时
>w_timeout=10                   # 写超时
>[static]                       # 静态文件section
>dir="/xxx/src/web/"            # 静态文件路径
>js="path/"                     # html 调用 /js/jquery.min.js  = dir + "path/jquery.min.js"
>css="css/"                     # 同上
>image="image/"                 # 同上
>favicon.ico="favicon.ico"      # 同上

>[section]                      #配置section
>key=val                        #配置内容

8、gear/ini.go 和 conf.ini
```
a) gear.IniParse.Parse(file) Ini    解析配置文件
b) Ini.[Bool/Int/Int64/Float/String/Strings][("section::key")/("key")]     获取etc/conf.ini中的配置信息

```

9、gear/log.go
```
a) 记录错误日志
b) 记录访问日志
```

## TODO
```
session
```
