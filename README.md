# 基于Kratos的微服务最佳实践框架

本框架基于kratos的大仓模式，结合了Kratos-layout和ThreeDotLabs的最佳实践，目录有所调整，支持DDD、CQRS，能够有效应对复杂的业务逻辑

还在开发中，可能存在错误

# 部署运行本演示项目

## 基于本模板创建自己的项目仓库

## 单机部署

## 集群部署

### 准备工作

- 找1台或多台服务器，将其中1台作为ansible管理节点，其他作为运行节点。如果只有一台， 那么管理节点和运行节点是同一台
    在master节点上，生成公钥

    ```bash
    ssh-keygen -C master@192.168.1.110 # 本地生成秘钥， 如果已有则跳过
    ```

- 允许使用非root账号sudo免密操作
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

### 运行节点主机列表配置

默认的配置在 `/etc/ansible/hosts` 中，追加自己的配置，和项目匹配到配置如下：

```ini
[all]
master hostname=master ansible_python_interpreter=/usr/bin/python3 ansible_host=192.168.72.36 ansible_port=22 ansible_user=sqf # ansible_pass='Tsss'
node1 hostname=node1 ansible_python_interpreter=/usr/bin/python3 ansible_host=192.168.72.84 ansible_port=22 ansible_user=user # ansible_pass='Tsss'

[registry]
master

[nacos]
master

[prometheus]
master

[backend]
node1
```

- 方括号[]中是组名，用于对系统进行分类，便于对不同系统进行个别的管理。
- ansible_user 指定登录该服务器的用户名，如果和当前系统已登录的用户名一样，则不需要指定。
- 如果修改组名，需同步更新playbook中的hosts字段

查看主机清单

```bash
ansible all --list-hosts
ansible webservers --list-hosts
ansible dbservers --list-hosts 
```

### 检查主机的连通性

```bash
ansible node1 -m ping 
```

### 允许使用sudo执行命令

修改 `/etc/ansible/ansible.cfg` 中的privilege_escalation配置为：

```ini
[privilege_escalation]
become=True
become_method=sudo
```

### 设置节点免密登录

```bash
ansible-playbook ./deploy/docker/ansible_playbook/ssh_login_no_password.yml
```

如果安装失败，需要手动复制key：`ssh-copy-id user@192.168.72.84`
### 批量设置hostname

```bash
ansible-playbook ./deploy/docker/ansible_playbook/modify_hostname.yml
```

### 部署基础设施

1. 安装docker

    ```bash
    ansible-playbook ./deploy/docker/ansible-playbook/install_docker_online.yml
    ```

2. 部署镜像私有仓库
    项目源码来自[这里](https://github.com/Joxit/docker-registry-ui), 部署操作：

    ```bash
    ansible-playbook ./deploy/docker/ansible_playbook/install_docker_registry.yml
    ```

3. 拉取和打包镜像

   ```bash
   ansible-playbook ./deploy/docker/prepare-infra-images.yml
   ```

4. 部署nacos

   ```bash
   ansible-playbook ./deploy/docker/install_nacos.yml
   ```

5. 部署prometheus和grafana

   ```bash
   ansible-playbook ./deploy/docker/install_prometheus_grafana.yml
   ```

  注意：
  
  - 默认只设置了nacos的数据源，如果需要增加其他数据源，修改 `./deploy/docker/config_prometheus.env` 配置
  - nacos的dashboard配置是： `deploy/docker/config_nacos_grafana_dashboard.json`

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
./bin/training -conf ./application/training/configs
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