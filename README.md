# webhook

> 一个极简的 git webhook，部署容易，非常简单就可以部署一个 webhook server。支持 **GitLab** 

> 本身是一个web服务执行 shell 目录下的sh文件

# 依赖

> go get github.com/Unknwon/goconfig

> go 1.12 开启mod 

# 安装

> go get github.com/qhzhui/webhook

> 修改conf.ini的配置，编译后用守护进程跑起来

> 自行编写shell目录下的sh文件，文件名和program参数一致 并加上可写权限

然后在浏览器中打开 `http://host:10086` 就可以看到下面的一些信息了（需要传递header的`X-Gitlab-Token`参数）
