wolffy
============
一个简单的代码发布工具, 将git版本库代码提交到各节点, 适合运维能力弱的小团队.

## Install

### master

发布工具部署的机器, master机器执行项目克隆, 将指定的版本打包通过HTTP发送到各节点.

> 需要有克隆代码库的权限

> 需要与node节点能够互相通过HTTP访问

```
wget https://github.com/l2x/wolffy/releases/download/v0.0.1/wolffy-v0.0.1.tar.gz 
tar xvf wolffy-v0.0.1.tar.gz
cd wolffy-v0.0.1/master
nohup ./wolffy-master > wolffy-master.log&
```

### node

代码部署的机器, 接收master提交的部署任务, 将代码部署到指定目录以及执行自定义脚本.

```
wget https://github.com/l2x/wolffy/releases/download/v0.0.1/wolffy-v0.0.1.tar.gz 
tar xvf wolffy-v0.0.1.tar.gz
cd wolffy-v0.0.1/agent
nohup ./wolffy-agent -pk=privateKey -master=masterip > wolffy-agent.log&
```

## Usage

[demo]()
