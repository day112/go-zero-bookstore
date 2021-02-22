## 启动：
### **注意事项**： dockerfile 文件配置了 LOCAL_HOST 环境变量

    1、项目目录下执行 ./docker.sh  脚本生成 rpc服务docker镜像
       ./docker.sh

    2、docker-compose-db 创建 mysql redis etcd 容器
      执行命令：
      docker-compose -f docker-compose-db.yml up -d

    注意：etcd 使用 k8s 自带的，启动 k8s 之后，需要手动执行 命令：
        etcd --advertise-client-urls 'http://0.0.0.0:2379' --listen-client-urls 'http://0.0.0.0:2379'
        注释：--listen-client-urls：对外提供服务的地址：比如http://ip:2379,http://127.0.0.1:2379，客户端会连接到这里和 etcd 交互。
            --advertise-client-urls：对外公告的该节点客户端监听地址，这个值会告诉集群中其他节点。
        website： https://www.cnblogs.com/xishuai/p/docker-etcd.html

    3、docker-compose-prom 创建 prometheus grafana 容器
      执行命令：
      docker-compose -f docker-compose-prom.yml up -d
    
    注意：
        http://127.0.0.1:9090/ 普罗米修斯 web url
        http://localhost:3000/ Grafana web url
      
    4、docker-compose 创建rpc服务容器，  rpc 服务基于依赖于etcd 服务
       执行命令：
       docker-compose up -d      
       
### rpc 服务
    1、定义数据边界
    2、数据库相互隔离、通过 rpc 相互调用
    3、服务间相互调用, 例如：add 之前需要调用 check 服务 
        1. 配置 add.yaml 文件添加 check 配置项
        2. Config addRpc 服务下 添加配置项
        3. ServiceContext 添加服务上下文配置
        4. 在 AddLogic 添加逻辑中 调用 check 服务


### 编写 k8s 服务
## 方法一
# 注意事项：上 k8s 部署的时候需要注意 etcd 配置，etcd 在 k8s 集群内服务地址
# 如何查找 etcd 在 k8s 的内部地址， 步骤如下
# 首先可以通过终端执行命令：`docker ps | grep etcd` 查找到 k8s 的 etcd 容器 `k8s_etcd_etcd-docker-desktop_kube-system_c7cc6a3c3118f127f5fd469ef69477e0_2`
# 然后执行命令： `kubectl get pods -n kube-system` 查找到 etcd 的 pod 名称 `etcd-docker-desktop`
# 最后执行命令： `kubectl describe pod etcd-docker-desktop -n kube-system` 查找到 pod 的详情，
# `advertise-client-urls=https://192.168.65.3:2379` 可以找到 etcd 的内部 IP

## 方法二
# k8s dashboard 
# 执行 `kubectl proxy` 命令
# 在浏览器访问 `http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/#/login`
# 命名空间选择 `kube-system` 然后找到 etcd 的 pod， 点击进入 pod 就可以看到详情。可以找到 etcd 的内部 IP

## 启动 kubernetes-dashboard 命令


##  生成 k8s checkrpc.yaml 配置文件

    goctl kube deploy -name checkrpc -namespace discov -image checkrpc:0.0.1 -o checkrpc.yaml -port 8081

## k8s 部署服务 如果 checkrpc namespace 不存在的话，请先通过 kubectl create namespace checkrpc 创建

    kubectl create namespace discov

## 首先部署 etcd

    kubectl apply -f etcd.yaml 

## 部署

    kubectl apply -f checkrpc.yaml 

## 测试
    kubectl run -i --tty --rm cli --image=checkrpc:0.0.1 -n discov -- sh

___________________________________________________________________________________________________________________

## 生成 k8s addrpc.yaml 配置文件

    goctl kube deploy -name addrpc -namespace discov -image addrpc:0.0.1 -o addrpc.yaml -port 8080

## k8s 部署服务 如果 addrpc namespace 不存在的话，请先通过 kubectl create namespace addrpc 创建

    kubectl create namespace discov

## 部署

    kubectl apply -f addrpc.yaml 

## 测试
    kubectl run -i --tty --rm cli --image=addrpc:0.0.1 -n discov -- sh
___________________________________________________________________________________________________________________

## 生成 k8s apirpc.yaml 配置文件

    goctl kube deploy -name apirpc -namespace discov -image apirpc:0.0.1 -o apirpc.yaml -port 8888

## k8s 部署服务 如果 apirpc namespace 不存在的话，请先通过 kubectl create namespace apirpc 创建

    kubectl create namespace discov

## 部署

    kubectl apply -f apirpc.yaml 

## 测试
    kubectl run -i --tty --rm cli --image=apirpc:0.0.1 -n discov -- sh
__________________________________________________________________________

