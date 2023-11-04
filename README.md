# 基于Kratos的微服务最佳实践框架

## 安装 Kratos

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## 下载框架

```bash
git clone https://github.com/shiqinfeng1/goMono.git
```

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
