wolffy
============
一个简单的代码发布工具.

## Usage

[demo](http://123.57.75.209:9020/), 帐号密码 `admin:123456`

## Install

### master

发布工具部署的机器, master机器执行项目克隆, 将指定的版本打包通过HTTP发送到各节点并解压到指定目录.

> 需要有克隆代码库的权限

> 需要与node节点能够互相通过HTTP访问

> 如需执行额外的脚本, 确保有执行权限

> 依赖mysql数据库, mysql中先创建好配置的数据库.


```
wget https://github.com/l2x/wolffy/releases/download/v0.0.1/wolffy-v0.0.1.tar.gz 
tar xvf wolffy-v0.0.1.tar.gz
cd wolffy-v0.0.1/master
```

编辑 `config/config.ini` 中的数据库连接信息.

```
nohup ./wolffy-master > wolffy-master.log&
```

访问 `yourip:9020`, 初始帐号密码`admin:123456`

### node

代码部署的机器, 接收master提交的部署任务, 将代码部署到指定目录以及执行自定义脚本.

```
wget https://github.com/l2x/wolffy/releases/download/v0.0.1/wolffy-v0.0.1.tar.gz 
tar xvf wolffy-v0.0.1.tar.gz
cd wolffy-v0.0.1/agent
nohup ./wolffy-agent -pk=privateKey -master=masterip > wolffy-agent.log&
```

