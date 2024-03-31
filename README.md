# 基于Kratos的微服务最佳实践框架

本框架基于kratos的大仓模式，结合了Kratos-layout和ThreeDotLabs的最佳实践，目录有所调整，技术架构和目录支持DDD、CQRS，解耦业务逻辑和技术框架，能够有效应对复杂的业务逻辑的开发。

## 目录说明

```bash

├── api
│   └── user
│       └── v1
├── app
│   ├── gomono-admin
│   ├── gomono-bff
│   │   ├── api
│   │   │   └── v1
│   │   ├── cmd
│   │   └── internal
│   │       ├── adapters
│   │       ├── application
│   │       ├── conf
│   │       ├── ports
│   │       └── service
│   ├── gomono-biz-trainer
│   │   ├── api
│   │   │   └── v1
│   │   ├── cmd
│   │   └── internal
│   │       ├── adapters
│   │       ├── application
│   │       │   ├── command
│   │       │   └── query
│   │       ├── conf
│   │       ├── domain
│   │       │   └── hour
│   │       ├── ports
│   │       └── service
│   ├── gomono-biz-training
│   │   ├── api
│   │   │   └── v1
│   │   ├── cmd
│   │   └── internal
│   │       ├── adapters
│   │       ├── application
│   │       │   ├── command
│   │       │   └── query
│   │       ├── conf
│   │       ├── domain
│   │       │   └── training
│   │       ├── ports
│   │       └── service
│   ├── gomono-gateway
│   └── gomono-task-migration
├── deploy
│   ├── docker
│   │   ├── biz
│   │   ├── biz-bff
│   │   ├── biz-gateway
│   │   ├── cluster-init
│   │   ├── images
│   │   ├── infra-docker-registry
│   │   ├── infra-mysql
│   │   ├── infra-nacos
│   │   ├── infra-prometheus
│   │   └── infra-redis
│   ├── k8s
│   └── standalone
│       ├── bin
│       └── config
├── third_party
│   ├── errors
│   ├── google
│   │   ├── api
│   │   └── protobuf
│   │       └── compiler
│   ├── openapi
│   │   └── v3
│   └── validate
└── utils
    ├── auth
    ├── client
    ├── conf
    ├── container
    ├── decorator
    ├── discovery
    ├── log
    ├── registrar
    ├── trace
    └── types
```

- api 接口定义输出文档
- app 业务程序，一个应用对应一个目录，每个目录内部对应一个标准的golang工程，典型的文件夹包括：
  - `api` 接口定义
  - `cmd` 主命令/子命令入口定义
  - `internal` 应用的内部模块，不可以被外部应用引用
  - `main.go` 主程序入口
- deploy 程序部署相关的自动化部署脚本，典型的部署脚本包括：
  - docker集群
  - k8s
  - standalone裸机部署
- third_party 第三方依赖，典型的有googl的基础功能的pb文件
- utils 包含所有应用都会用到的基础工具包，如果该工具包在所有项目中都通用，可以独立出来变成公司级别的kit基础包

## docker集群的环境准备和快速部署

### 环境准备

找1台或多台服务器，将其中1台（master节点）作为ansible管理节点，其他作为运行节点。如果只有一台， 那么管理节点和运行节点是同一台

- 在master节点上，生成公钥

    ```bash
    ssh-keygen -C master@192.168.1.110 # 本地生成秘钥， 如果已有则跳过
    ```

- 在每台服务器上配置允许非root账号进行sudo免密操作
  测试环境可以直接使用root账号， 如果为了安全性则使用非root账号， 但需要确保具有sudo权限。
  
  新建非root账号

    ```bash
    sudo useradd ansible
    sudo passwd ansible
    ```

  配置sudo权限

  ```bash
  sudo chmod u+w -v /etc/sudoers
  sudo sh -c 'echo "ansible ALL=(ALL)       NOPASSWD: ALL" >> /etc/sudoers'
  sudo chmod u-w -v /etc/sudoers
  ```

### 在管理节点（master）上安装ansible

centos7

```bash
sudo yum install -y epel-release
sudo dnf install ansible
ansible --version  # 输出版本信息，例如： ansible 2.9.27
```

ubuntu

```bash
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible
```

安装完成后，默认配置文件在 `/etc/ansible/` 下

```bash
ls /etc/ansible/
ansible.cfg  hosts  roles
```

### 在管理节点（master）上配置节点主机列表及ansible

默认的配置在 `/etc/ansible/hosts` 中，追加自己的配置,配置文件demo在 `./hosts.ansible`.

配置字段说明：

- 方括号[]中是组名，用于对业务/功能系统进行分类，便于对不同系统进行个别的管理。下面为主机名称
- ansible_user 指定登录该服务器的用户名，如果和当前系统已登录的用户名一样，则不需要指定。
- 注意：如果修改组名，需同步更新playbook中的hosts字段

```ini
[all]
master hostname=master ansible_python_interpreter=/usr/bin/python3 ansible_host=192.168.72.36 ansible_port=22 ansible_user=sqf # ansible_pass='Tsss'
node1 hostname=node1 ansible_python_interpreter=/usr/bin/python3 ansible_host=192.168.72.84 ansible_port=22 ansible_user=user # ansible_pass='Tsss'

[images]
master

[registry]
master

[prometheus]
master

[jaeger]
master

[nacos]
node1

[mysql]
node1

[mysql_slave]
node1

[redis]
node1

[gateway]
node1

[bff]
node1

[backend]
node1
```

配置修改好之后，可以通过如下命令查看主机清单：

  ```bash
  ansible all --list-hosts  # all 为组名
  ansible webservers --list-hosts # webservers 为组名
  ansible dbservers --list-hosts  # dbservers 为组名
  ... 
  ```

检查主机的连通性

  ```bash
  ansible node1 -m ping 
  ```

允许使用sudo执行命令

  修改 `/etc/ansible/ansible.cfg` 中的privilege_escalation配置为：

  ```ini
  [privilege_escalation]
  become=True
  become_method=sudo
  ```

### 设置所有工作节点免密登录

原理：将master节点的的公钥复制到节点上

登录master管理节点，执行：

```bash
ansible-playbook ./deploy/docker/cluster-init/ssh_login_no_password.yml
```

如果安装失败，需要手动复制key：`ssh-copy-id user@192.168.72.84`

### 批量设置hostname

```bash
ansible-playbook ./deploy/docker/cluster-init/modify_hostname.yml
```

### 批量设置dns

```bash
ansible-playbook ./deploy/docker/cluster-init/config_dns.yml
```

## 部署基础设施

复制deploy文件夹到管理节点，并在管理节点上执行下述操作。

1. 安装docker和docker-compose

    ```bash
    ansible-playbook ./deploy/docker/infra-docker-registry/install_docker_online.yml
    ```

    > 注意：如果执行palybook时报错：
    ` fatal: [node1]: FAILED! => {"changed": false, "msg": "The Python 2 bindings for rpm are needed for this module. If you require Python 3 support use the `dnf`Ansible module instead.. The Python 2 yum module is needed for this module. If you require Python 3 support use the dnf Ansible module instead."}`, 需要将`/etc/ansible/hosts`中的`ansible_python_interpreter=/usr/bin/python3` 改为 `ansible_python_interpreter=/usr/bin/python2`， 或者，在ansible-playbook命令后面添加参数：`-e ansible_python_interpreter=/usr/bin/python2`

2. 部署镜像私有仓库
  
    私有仓库的项目源码来自[这里](https://github.com/Joxit/docker-registry-ui), 部署操作：

    ```bash
    ansible-playbook ./deploy/docker/infra-docker-registry/install_docker_registry.yml
    ```

    部署完成后，私有仓库在 `http://<域名/IP>:8080/` 上提供服务。（使用域名时，可能需要配置本地DNS）

3. 拉取和打包基础设施的镜像

   ```bash
   ansible-playbook ./deploy/docker/images/prepare-infra-images.yml
   ```

4. 部署nacos

    部署

    ```bash
    ansible-playbook ./deploy/docker/infra-nacos/install.yml
    ```

    停用

    ```bash
    ansible-playbook ./deploy/docker/infra-nacos/stop.yml
    ```

5. 部署prometheus和grafana

    部署

    ```bash
    ansible-playbook ./deploy/docker/infra-prometheus/install.yml
    ```

    停用

    ```bash
    ansible-playbook ./deploy/docker/infra-prometheus/stop.yml
    ```

    注意：

    - 默认只设置了nacos的数据源，如果需要增加其他数据源，修改 `./deploy/docker/infra-prometheus/config_prometheus.env` 配置
    - nacos的dashboard配置是： `deploy/docker/infra-prometheus/config_nacos_grafana_dashboard.json`

6. 部署mysql

    部署

    ```bash
    ansible-playbook ./deploy/docker/infra-mysql/install.yml
    ```

    停用

    ```bash
    ansible-playbook ./deploy/docker/infra-mysql/stop.yml
    ```

7. 部署redis

    部署

    ```bash
    ansible-playbook ./deploy/docker/infra-redis/install.yml
    ```

    停用

    ```bash
    ansible-playbook ./deploy/docker/infra-redis/stop.yml
    ```

## 开发微服务

1. 添加自己的服务
  参考 `app/gomono-biz-xxx` 下的文件夹复制或者新建一份，注意：`main.go` 必须放在`app/gomono-biz-xxx`目录下
2. 添加一个proto文件,根据业务编写proto接口
3. 生成protobuffer的客户单代码和服务端代码

    ```bash
    make api
    ```

4. 编译应用镜像

  如果没有私有仓库，按照上面的步骤部署一个私有仓库。
  修改Makefile文件中的`CONTAINER_REGISTRY`为镜像仓库的地址， 例如：`node1:8080`
  编译所有镜像,并推送镜像到私有仓库：

  ```bash
  make all
  ```

## Docker集群部署

根据实际项目需求部署具体哪些应用

```bash
# 对应的停止命令是： ansible-playbook ./deploy/docker/biz-gateway/stop.yml
ansible-playbook ./deploy/docker/biz-gateway/install.yml

# 对应的停止命令是： ansible-playbook ./deploy/docker/biz-bff/stop.yml
ansible-playbook ./deploy/docker/biz-bff/install.yml
```

## 本地运行

```bash
./bin/training -conf ./application/training/configs
```

## 框架和功能详解

### 运行环境/模式

有3种运行环境

- product 生产环境
- test 测试环境
- debug 调试环境

有如下的区别

- 配置文件的源
  生产环境环境下只能从配置中心获取配置，如果没读到，那么将无法运行
  测试环境key从配置中心，也可以从本地读取配置，如果都没读到，那么将无法运行
  调试环境值只从本地读取配置，如果没读到，那么将无法运行

#### 领域层

领域层的职责范围和设计规范

1. 定义领域实体的数据结构，例如 `internal/training/domain/training/user.go` `internal/training/domain/training/training.go`
2. 定义`Repository`存储库接口，满足依赖倒置原则。repo在adpters中实现，在app层的useCase中使用，而不绑定到领域实体，便于测试以及实现CQRS，因为CQRS的查询是不需要经过领域实体的，可以直接使用repo
3. 实现领域实体的行为。如果行为较多，可以拆分为多个文件。注意！！！这是围绕「行为」定义和实现实体的对外接口函数，而不是围绕字段进行set和get实现的类似贫血模型接口
4. 领域层的所有错误都必须定义名称，不能直接返回`errors.New(...)` 或者`fmt.Errorf(...)`，而是应该先全局定义再返回 `var ErrXxxx = errors.New("xxx") ... return ErrXxxx`
