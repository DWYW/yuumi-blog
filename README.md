# YUUMI BLOG

利用 go + vue + nuxt 构建自己的博客


## 目录结构

```
├─ admin                           管理后台
|   ├─ administrator               管理员服务
|   |   ├─ scripts
|   |   ├─ src
|   |   ├─ .env.development
|   |   └─ .env.production
|   ├─ article                     文章服务
|   |   ├─ scripts
|   |   ├─ src
|   |   ├─ .env.development
|   |   └─ .env.production
|   ├─ entry                       基座服务
|   |   ├─ scripts
|   |   ├─ src
|   |   ├─ .env.development
|   |   └─ .env.production
├─ client                          客户端
|   ├─ assets                      资源文件夹
|   ├─ components                  存放可导入到你的页面中的 Vue 组件或其他组件
|   ├─ composables                 约定composables目录下创建的组合逻辑文件将会被系统自动识别导入到应用程序，以供全局使用
|   ├─ pages                       Nuxt会自动集成vue-router，结合pages目录下的文件(夹)名来构建我们的项目
|   └─ nuxt.config.ts              配置文件
├─ server                          服务端
|   ├─ api                         API描述文件
|   ├─ app
|   |   ├─ init                    初始化内容
|   |   ├─ interface               外部服务
|   |   |   ├─ administrator
|   |   |   |   ├─ cmd             入口
|   |   |   |   ├─ config          配置信息
|   |   |   |   └─ internal        内部文件
|   |   |   ├─ article
|   |   |   |   ├─ cmd             入口
|   |   |   |   ├─ config          配置信息
|   |   |   |   └─ internal        内部文件
|   |   |   ├─ passport
|   |   |   |   ├─ cmd             入口
|   |   |   |   ├─ config          配置信息
|   |   |   |   └─ internal        内部文件
|   |   ├─ service                 内部服务
|   |   |   ├─ administrator
|   |   |   |   ├─ cmd             入口
|   |   |   |   ├─ config          配置信息
|   |   |   |   └─ internal        内部文件
|   |   |   ├─ article
|   |   |   |   ├─ cmd             入口
|   |   |   |   ├─ config          配置信息
|   |   |   |   └─ internal        内部文件
|   ├─ middleware                  服务中间件
|   ├─ pkg                         包文件
|   └─ third_party                 其他第三方文件
```
## 部署客户端

> 注意替换source_path

```bash
# 生成部署文件
make buildClient
# 上传服务器
scp -r dist/client/* user@ip:source_path/client
```

supervisor 配置服务

```
# cat /etc/supervisor/conf.d/blog.client.conf

[program:blog_client]
command = sudo PORT=8080 node source_path/client/server/index.mjs
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = root
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_client.log
```

nginx 配置

```
server {
    listen 80;
    server_name client.xxxx.com;

    location / {
        return 301 https://$host$request_uri;
    }
}

server {
    listen 443 ssl;
    server_name client.xxxx.com;
    ssl_certificate client.xxxx.com.crt;
    ssl_certificate_key client.xxxx.com.key;
    ssl_session_timeout 5m;
    #请按照以下协议配置
    ssl_protocols TLSv1.2 TLSv1.3;
    #请按照以下套件配置，配置加密套件，写法遵循 openssl 标准。
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;


    location / {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:8080$request_uri;
    }
}
```

## 部署服务端

> 注意替换source_path

```bash
# 生成部署文件
make buildAdministratorServer
make buildArticleServer
make buildAdministratorInterface
make buildArticleInterface
make buildPassportInterface

# 上传服务器
scp -r dist/server/* user@ip:source_path/server
```

supervisor 配置服务

```
# cat /etc/supervisor/conf.d/server.client.conf

[program: administrator_server]
command= source_path/server/administrator.server -config source_path/server/administrator.server.config.yaml
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = ubuntu
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_server.log

[program: article_server]
command= source_path/server/article.server -config source_path/server/article.server.config.yaml
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = ubuntu
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_server.log

[program: passport_interface]
command= source_path/server/passport.interface.server -config source_path/server/passport.interface.config.yaml
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = ubuntu
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_server.log

[program: administrator_interface]
command = source_path/server/administrator.interface.server -config source_path/server/administrator.interface.config.yaml
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = ubuntu
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_server.log

[program: article_interface]
command = source_path/server/article.interface.server -config source_path/server/article.interface.config.yaml
autostart = true
startsecs = 3
autorestart = true
startretries = 3
user = ubuntu
redirect_stderr = true
stdout_logfile_maxbytes = 50MB
stdout_logfile_backups = 20
stdout_logfile = /tmp/blog_server.log
```

nginx 配置

```
server {
    listen 80;
    server_name api.xxxx.com;
    return 301 https://$host$request_uri;
}

server {
    listen 443 ssl;
    server_name api.xxxx.com;
    ssl_certificate api.xxxx.com.crt;
    ssl_certificate_key api.xxxx.com.key;
    ssl_session_timeout 5m;
    #请按照以下协议配置
    ssl_protocols TLSv1.2 TLSv1.3;
    #请按照以下套件配置，配置加密套件，写法遵循 openssl 标准。
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;

    location ~ ^/passport-api/(.*)$ {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:9005/$1?$query_string;
    }

    location ~ ^/administrator-api/(.*)$ {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:9009/$1?$query_string;
    }

    location ~ ^/article-api/(.*)$ {
        proxy_set_header X-Real-IP $remote_addr;
        proxy_pass http://127.0.0.1:9007/$1?$query_string;
    }
}
```

## 部署管理端

> 注意替换source_path

```bash
# 生成部署文件
make buildAdminEntry
make buildAdmiAdministrator
make buildAdminArticle
# 上传服务器
scp -r dist/admin/* user@ip:source_path/admin
```

nginx 配置

```
server {
    listen 80;
    server_name admin.xxxx.com;

    location / {
        return 301 https://$host$request_uri;
    }
}


server {
    listen 443 ssl;
    server_name admin.xxx.com;
    ssl_certificate admin.xxx.com.crt;
    ssl_certificate_key admin.xxx.com.key;
    ssl_session_timeout 5m;
    #请按照以下协议配置
    ssl_protocols TLSv1.2 TLSv1.3;
    #请按照以下套件配置，配置加密套件，写法遵循 openssl 标准。
    ssl_ciphers ECDHE-RSA-AES128-GCM-SHA256:HIGH:!aNULL:!MD5:!RC4:!DHE;
    ssl_prefer_server_ciphers on;


    location / {
        root   source_path/admin/entry;
        index  index.html;
    }

    location ~ ^/administrator/ {
        root   source_path/admin;
        index  index.html;
    }

    location ~ ^/article/ {
        root   source_path/admin;
        index  index.html;
    }
}
```

