# 基于Kratos的微服务最佳实践框架
本框架基于kratos的大仓模式，微服务入口在app下。默认有3个服务。

## 安装 Kratos

```bash
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

## 下载框架

1. 在github上fork该goMono仓库，例如fork后仓库为：`https://github.com/xiaoming/didi.git`
2. clone到本地，例如：`git clone https://github.com/xiaoming/didi.git`
3. 替换`github.com/shiqinfeng1/goMono` 为`github.com/xiaoming/didi`

```bash
git clone https://github.com/shiqinfeng1/goMono.git
```

## 添加自己的服务
直接拷贝app下的服务目录，


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

### 工程目录

工程目录结合了Kratos-layout和ThreeDotLabs的最佳实践，支持DDD、CQRS，能够有效应对复杂的业务逻辑

#### 领域层

领域层的职责范围和设计规范

1. 定义领域实体的数据结构，例如 `internal/trainings/domain/training/user.go` `internal/trainings/domain/training/training.go`
2. 定义`Repository`存储库接口，满足依赖倒置原则。repo在adpters中实现，在app层的useCase中使用，而不绑定到领域实体，便于测试以及实现CQRS，因为CQRS的查询是不需要经过领域实体的，可以直接使用repo
3. 实现领域实体的行为。如果行为较多，可以拆分为多个文件。注意！！！这是围绕「行为」定义和实现实体的对外接口函数，而不是围绕字段进行set和get实现的类似贫血模型接口
4. 领域层的所有错误都必须定义名称，不能直接返回`errors.New(...)` 或者`fmt.Errorf(...)`，而是应该先全局定义再返回 `var ErrXxxx = errors.New("xxx") ... return ErrXxxx`
5. 