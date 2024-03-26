当前文件夹下的内容为： 通过absilbe使用docker在服务器集群上部署本系统

- `ansible_playbook`目录存放ansible运行脚本
- `tmpl`目录存放配置文件模板和docker-compose文件模板，这些模板文件在ansible_playbook中被使用，并且其中的变量在执行ansible_playbook 命令时被替换
- `ansible_playbook/var.yaml`中存放所有可配置的变量，可按照实际环境修改

