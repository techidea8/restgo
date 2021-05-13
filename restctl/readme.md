#安装
```
go install github.com/techidea8/restgo/restctl@latest
```

#配置文件说明
```yaml
database: dbprintx
table:  test
username: winlion
password: 1D007648b4f8
model: 
dstdir: ./
package: turingapp
port: 3306
host: iot.techidea8.com
```


#使用方法
```bash
restctl -h
flag needs an argument: -h
Usage of restctl:
  -a string
        mysql port (default "3306")
  -c string
        config file path (default "./restgo.yaml")
  -db string
        database name (default "test")
  -h string
        database host (default "127.0.0.1")
  -m string
        out model
  -o string
        dist dir (default "./")
  -p string
        password
  -pkg string
        application package (default "turinapp")
  -t string
        table name (default "test")
  -u string
        user name (default "root")
```

# 效果
系统将自动生成如下文件
```bash
restctl -c ./restgo.yaml -t biz_order -m order -db 127.0.0.1 -u root -p root -o ./code -pkg turingapp
```
```bash
├─server
│  ├─args
│  │      order.go
│  │
│  ├─ctrl
│  │      order.go
│  │
│  ├─model
│  │      order.go
│  │
│  └─service
│          order.go
│
└─ui
    ├─api
    │      order.js
    │
    └─view
        └─order
                list.vue

```