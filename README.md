# 基于Kratos的微服务最佳实践框架

本框架基于kratos的大仓模式，结合了Kratos-layout和ThreeDotLabs的最佳实践，目录有所调整，支持DDD、CQRS，能够有效应对复杂的业务逻辑

# 部署运行本演示项目

## 基于本模板创建自己的项目仓库

## 单机部署

## 集群部署

### 准备工作

- 找1台或多台服务器，将1台作为ansible管理节点，其他作为运行节点。如果只有一台， 那么管理节点和运行节点是同一台，主机名和dns只需要配置一次就行。
- 免密登录远端服务器
    在master节点上，生成公钥

    ```bash
    ssh-keygen -C master@192.168.1.110 # 本地生成秘钥， 如果已有则跳过
    ```

    如果有独立运行节点，拷贝master节点的公钥到运行节点上

    ```bash
    ssh-copy-id -f -i ~/.ssh/id_rsa.pub sqf@192.168.1.184
    ```

    检查是否生效: `ssh 'sqf@192.168.1.184'`, 如果微提示输入密码，表示已生效
    如果未生效， 参考这里解决：https://www.slw.ac.cn/article/linux-cmd-remotelogin.html
    如果本机也是作为被ansible管理的主机，也需要设置本机免密登录（ssh-copy-id到本机）

- 下载安装dns域名管理工具
  
    ```bash
    mkdir hostctl && cd hostctl
    wget https://github.com/guumaster/hostctl/releases/download/v1.1.4/hostctl_1.1.4_linux_64-bit.tar.gz 
    tar -zxvf hostctl_1.1.4_linux_64-bit.tar.gz
    sudo mv hostctl /usr/local/bin/
    cd .. && rm -rf hostctl
    hostctl -v
    ```

    hostctl的使用

    ```bash
    # 查看dns
    hostctl list
    # 添加dns记录
    sudo hostctl add domains ansible node1 --ip 192.168.72.84  # 其中 ansible 表示管理域， node是域名
    # 启用禁用本地域名解析
    sudo hostctl disable  ansible
    sudo hostctl enable  ansible
    # 删除域名解析
    sudo hostctl remove domains ansible node1
    ```

- 设置主机名
  一般设置管理节点的主机名为master， 运行节点的主机名为nodeX， X为1,2,3...一次编号

  ```bash
   # 登录管理节点后，示例：ssh sqf@192.168.1.184
   hostnamectl set-hostname master
   # 登录运行节点后
   hostnamectl set-hostname node1
   hostnamectl set-hostname node2
   ...
  ```

  如果主机较多， 也可以使用下面介绍的ansible来批量设置

- 非root账号
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

### 在管理节点上安装ansible

centos7

```bash
sudo yum install -y epel-release
sudo dnf install ansible
# sudo dnf install ansible-collection-community-general
ansible --version  # 输出版本信息，例如： ansible 2.9.27
```

ubuntu

```bash
$ sudo apt update
$ sudo apt install software-properties-common
$ sudo add-apt-repository --yes --update ppa:ansible/ansible
$ sudo apt install ansible
```

安装完成后，默认配置文件在 `/etc/ansible/` 下

```bash
ansible.cfg  hosts  roles
```

### 运行节点主机列表配置

默认的配置在 `/etc/ansible/hosts` 中，追加自己的配置，例如：

```ini
[all]
master hostname=master ansible_python_interpreter=/usr/bin/python3 ansible_ssh_host=192.168.72.36 ansible_ssh_port=22 ansible_ssh_user=sqf # ansible_ssh_pass='Tsss'
node1 hostname=node1 ansible_python_interpreter=/usr/bin/python3 ansible_ssh_host=192.168.72.84 ansible_ssh_port=22 ansible_ssh_user=user # ansible_ssh_pass='Tsss'
[registry]
master
[webservers]
node1
[dbservers]
node1
```

- 方括号[]中是组名，用于对系统进行分类，便于对不同系统进行个别的管理。
- ansible_user 是指定登录该服务器的用户名，如果不指定，将使用当前系统已登录的用户名，一个节点只需要指定一处

查看服务器列表

```bash
ansible all --list-hosts
ansible webservers --list-hosts
ansible dbservers --list-hosts 
```

### 其他基础操作举例

检查主机的连通性

```bash
ansible node1 -m ping 
```

使用sudo执行命令
命令行下

```bash
# -m command 表示使用command模块，可以不写，默认就是使用command模块
# -a表示模块的参数args
ansible all -m command -a 'tail /etc/sudoers' -become=yes --become-method=sudo
```

全局修改: 修改 `/etc/ansible/ansible.cfg` 中的如下配置

```ini
[privilege_escalation]
become=True
become_method=sudo
```

修改后在playbook或命令行中可以不加 `become/become-method` 这些配置了

将文件直接传输到all组中的所有服务器

```bash
# mode=600 文件属性 
# owner=mdehaan group=mdehaan 文件所有者
ansible all -m copy -a "src=/etc/hosts dest=/tmp/hosts mode=600 owner=mdehaan group=mdehaan"
```

批量设置hostname

```bash
ansible-playbook ./deploy/docker/cluster/ansible_playbook/modify_hostname.yml
```

### 检查剧本有效性及彩排

举例：

```bash
ansible-playbook --syntax-check deploy/docker/cluster/ansible-playbook/install_docker-online.yml
ansible-playbook --check deploy/docker/cluster/ansible-playbook/install_docker-online.yml
```

### 部署基础设施

2. 安装docker

    ```bash
    ansible-playbook ./deploy/docker/cluster/ansible-playbook/install_docker_online.yml
    ```

2. 自建镜像仓库
    项目来自[这里](https://github.com/Joxit/docker-registry-ui), 部署操作：

    ```bash
    ansible-playbook ./deploy/docker/cluster/ansible_playbook/install_docker_registry.yml
    ```

3. 制作镜像
   
   ```bash
   ./deploy/docker/cluster/prepare-infra-images.sh
   ```

4. 部署nacos
    方式： 拷贝docker-compose.yml到运行节点，并启动

    ```bash

    ```

部署mysql集群
方式：拷贝docker-compose.yml到运行节点，并启动

# 开发微服务

## 添加自己的服务



## 初始化和编译

```bash
make all
```

## 添加一个proto文件,根据业务编写proto接口

## 生成protobuffer的客户单代码和服务端代码

```bash
make api
```

## 本地运行

```bash
./bin/training -conf ./app/training/configs
```

## Docker运行

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

## 框架和功能详解


#### 领域层

领域层的职责范围和设计规范

1. 定义领域实体的数据结构，例如 `internal/training/domain/training/user.go` `internal/training/domain/training/training.go`
2. 定义`Repository`存储库接口，满足依赖倒置原则。repo在adpters中实现，在app层的useCase中使用，而不绑定到领域实体，便于测试以及实现CQRS，因为CQRS的查询是不需要经过领域实体的，可以直接使用repo
3. 实现领域实体的行为。如果行为较多，可以拆分为多个文件。注意！！！这是围绕「行为」定义和实现实体的对外接口函数，而不是围绕字段进行set和get实现的类似贫血模型接口
4. 领域层的所有错误都必须定义名称，不能直接返回`errors.New(...)` 或者`fmt.Errorf(...)`，而是应该先全局定义再返回 `var ErrXxxx = errors.New("xxx") ... return ErrXxxx`
5. 