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
    a) 见main.go，gear.AddRoute("/test", &testApp{})
    b) 路由有两级，/module/action
    c) 由testApp接管/test的路由，module为test，自动检测testApp内XxxAction命名的函数注册为action，例如：/test/go 会自动路由到testApp.GoAction
    d) testApp必须定义IndexAction，作为没有action时默认执行方法，/test == /test/index
```

2、gear/app.go
```
    a) 继承自request，response，实现appInerface
    b) 业务app需继承gear.App
```

3、gear/check.go
```
    a) 定义接口CheckInterface，需求：在App业务执行之前做公共验证，ip/auth/login等等
    b) 实现此接口需要实现Check，Failed 两函数
    c) app实现了CheckInterface后，框架在执行action之前会调用Check，返回false验证失败，执行Failed，返回true验证成功，执行相应路由action
    d) 针对类似ip黑名单/登录检测等，可实现一个公共Verify（参考testApp），对需要做检测的app则继承Verify
    e) 针对单个app 特殊验证，直接在app内实现以上两方法即可
```

4、gear/func.go
```
    a) 公共无其他依赖的基础方法
```

5、gear/gear.go
```
    a) 框架入口文件，用来启动服务
```

6、gear/ini.go 和 conf.ini
```
    a) 内容示例
        ###########################################################
        addr=":8080"                # 服务监听地址
        r_timeout=10                # 读超时
        w_timeout=10                # 写超时

        [static]                    # 静态文件section
        dir="/xxx/src/web/"         # 静态文件路径
        js="path/"                  # html 调用 /js/jquery.min.js  = dir + "path/jquery.min.js"
        css="css/"                  # 同上
        image="image/"              # 同上
        favicon.ico="favicon.ico"   # 同上
        
        #example
        [section]                   #配置section
        key=val                     #配置内容
        ###########################################################
    b) gear.Ini.[Bool/Int/Int64/Float/String/Strings][("section::key")/("key")]     获取etc/conf.ini中的配置信息
    c) gear.IniParse.Parse(file)    解析配置文件
```

8、gear/log.go
```
    a) 记录错误日志
```

9、gear/request.go
```
    a) 封装 *http.Request
    b) 实现 GetXxx/PostXxx/Ip/Ip2Long 等基于*http.Request的方
```

10、gear/response.go
```
    a) 封装 http.ResponseWriter
    b) 实现 Display/Json/Redirect 等基于http.ResponseWriter方法
```

11、gear/route.go
```
    a) 路由
```

## TODO
```
session
```
