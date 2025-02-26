# IOC-golang：一款 GO 语言依赖注入框架

```
  ___    ___     ____                           _                         
 |_ _|  / _ \   / ___|           __ _    ___   | |   __ _   _ __     __ _ 
  | |  | | | | | |      _____   / _` |  / _ \  | |  / _` | | '_ \   / _` |
  | |  | |_| | | |___  |_____| | (_| | | (_) | | | | (_| | | | | | | (_| |
 |___|  \___/   \____|          \__, |  \___/  |_|  \__,_| |_| |_|  \__, |
                                |___/                               |___/ 
```

[![IOC-golang CI](https://github.com/alibaba/IOC-golang/actions/workflows/github-actions.yml/badge.svg)](https://github.com/alibaba/IOC-golang/actions/workflows/github-actions.yml)
[![License](https://img.shields.io/badge/license-Apache%202-4EB1BA.svg)](https://www.apache.org/licenses/LICENSE-2.0.html)

[文档](https://ioc-golang.github.io/cn)

[English Docs](https://ioc-golang.github.io)

[English README](./README.md)

![demo gif](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/video/ioc-golang-demo.gif)

IOC-golang 是一款强大的 Go 语言依赖注入框架，提供了一套完善的 IoC 容器。其能力如下：

- [依赖注入](https://ioc-golang.github.io/cn/docs/getting-started/tutorial/)

  支持任何结构、接口的依赖注入，具备完善的对象生命周期管理机制。

  可以接管对象的创建、参数注入、工厂方法。可定制化对象参数来源。

- [结构代理层](https://ioc-golang.github.io/cn/docs/examples/debug/)

  基于 AOP 的思路，为由框架接管的对象提供默认的结构代理层，在面向接口编程的情景下，可以使用基于结构代理 AOP 层扩展的丰富运维能力。例如接口查询，参数动态监听，方法粒度链路追踪，性能瓶颈分析，分布式场景下全链路方法粒度追踪等。

- [代码生成能力](https://ioc-golang.github.io/cn/docs/reference/iocli/#%E7%BB%93%E6%9E%84%E6%B3%A8%E8%A7%A3)

  我们提供了代码生成工具，开发者可以通过注解的方式标注结构，从而便捷地生成结构注册代码、结构代理、结构专属接口等。

- [可扩展能力](https://ioc-golang.github.io/cn/docs/contribution-guidelines/)

  支持被注入结构的扩展、自动装载模型的扩展、调试 AOP 层的扩展。

- [丰富的预置组件](https://ioc-golang.github.io/cn/docs/examples/)

  提供覆盖主流中间件的预制对象，方便直接注入使用。


## 项目结构

- **autowire：** 提供依赖注入内核，以及单例模型、多例模型两种基本自动装载模型
- **config：** 配置加载模块，负责解析 yaml 格式的配置文件
- **debug：** 调试模块：提供调试 API、提供调试注入层实现
- **extension：** 组件扩展目录：提供基于多种注入模型的预置实现结构：

    - autowire：自动装载模型扩展

        - grpc：grpc 客户端模型

        - config：配置模型

        - rpc：远程过程调用模型
        
        - triple: Dubbo3 支持【待完善】
    - config：配置注入模型扩展结构
    
        - string,int,map,slice
    - normal：多例模型扩展结构

        - redis
    
        - http_server
        - mysql

        - rocketmq
    
        - nacos
- **example：** 示例仓库
- **iocli：** 代码生成/程序调试 工具


## 快速开始

### 安装代码生成工具

```shell
% go install github.com/alibaba/ioc-golang/iocli@latest
% iocli
hello
```

### 依赖注入教程

我们将开发一个具有以下拓扑的工程，在本例子中，可以展示

1. 注册代码生成
2. 接口注入
3. 对象指针注入
4. API 获取对象
5. 调试能力，查看运行中的接口、方法；以及实时监听参数值、返回值。

![ioc-golang-quickstart-structure](https://raw.githubusercontent.com/ioc-golang/ioc-golang-website/main/resources/img/ioc-golang-quickstart-structure.png)


用户所需编写的全部代码：main.go

```go
package main

import (
	"fmt"
	"time"

	"github.com/alibaba/ioc-golang"
)

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type App struct {
	// 将封装了代理层的 main.ServiceImpl1 指针注入到 Service 接口，单例模型，需要在标签中指定被注入结构
	ServiceImpl1 Service `singleton:"main.ServiceImpl1"`

	// 将封装了代理层的 main.ServiceImpl2 指针注入到 Service 接口，单例模型，需要在标签中指定被注入结构
	ServiceImpl2 Service `singleton:"main.ServiceImpl2"`

	// 将封装了代理层的 main.ServiceImpl1 指针注入到他的专属接口 'ServiceImpl1IOCInterface'
  // 注入专属接口的命名规则是 '${结构名}IOCInterface'，注入专属接口无需指定被注入结构，标签值为空即可。
	Service1OwnInterface ServiceImpl1IOCInterface `singleton:""`

	// 将结构体指针注入当前字段
	ServiceStruct *ServiceStruct `singleton:""`
}

func (a *App) Run() {
	for {
		time.Sleep(time.Second * 3)
		fmt.Println(a.ServiceImpl1.GetHelloString("laurence"))
		fmt.Println(a.ServiceImpl2.GetHelloString("laurence"))

		fmt.Println(a.Service1OwnInterface.GetHelloString("laurence"))
		
		fmt.Println(a.ServiceStruct.GetString("laurence"))
	}
}

type Service interface {
	GetHelloString(string) string
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ServiceImpl1 struct {
}

func (s *ServiceImpl1) GetHelloString(name string) string {
	return fmt.Sprintf("This is ServiceImpl1, hello %s", name)
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ServiceImpl2 struct {
}

func (s *ServiceImpl2) GetHelloString(name string) string {
	return fmt.Sprintf("This is ServiceImpl2, hello %s", name)
}

// +ioc:autowire=true
// +ioc:autowire:type=singleton

type ServiceStruct struct {
}

func (s *ServiceStruct) GetString(name string) string {
	return fmt.Sprintf("This is ServiceStruct, hello %s", name)
}

func main() {
	// start
	if err := ioc.Load(); err != nil {
		panic(err)
	}

	// app, err := GetAppIOCInterface 也可以，获取到的是封装了代理层的接口，如下获取到的是未封装的结构体指针。
	app, err := GetApp()
	if err != nil {
		panic(err)
	}
	app.Run()
}

```

上述所说的“代理层”，是框架为“以接口形式注入/获取”的结构体，默认封装的代理，可以扩展一系列运维操作。我们推荐开发者在编写代码的过程中基于接口编程，则所有对象都可拥有运维能力。

编写完毕后，当前目录执行以下命令，初始化 go mod ，拉取最新代码，生成结构注册代码。（mac 环境可能因权限原因需要sudo）：

```bash
% go mod init ioc-golang-demo
% export GOPROXY="https://goproxy.cn"
% go mod tidy
% go get github.com/alibaba/ioc-golang@master
% sudo iocli gen
```

会在当前目录生成：zz_generated.ioc.go，开发者无需关心这一文件，这一文件中就包含了上面使用的 GetApp 方法

```go
//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package main

import (
        autowire "github.com/alibaba/ioc-golang/autowire"
        normal "github.com/alibaba/ioc-golang/autowire/normal"
        "github.com/alibaba/ioc-golang/autowire/singleton"
        util "github.com/alibaba/ioc-golang/autowire/util"
)

func init() {
        normal.RegisterStructDescriptor(&autowire.StructDescriptor{
                Factory: func() interface{} {
                        return &app_{}
                },
        })
        singleton.RegisterStructDescriptor(&autowire.StructDescriptor{
                Factory: func() interface{} {
                        return &App{}
                },
        })
  ...
func GetServiceStructIOCInterface() (ServiceStructIOCInterface, error) {
        i, err := singleton.GetImplWithProxy(util.GetSDIDByStructPtr(new(ServiceStruct)), nil)
        if err != nil {
                return nil, err
        }
        impl := i.(ServiceStructIOCInterface)
        return impl, nil
}

```

查看当前目录文件

```bash
% tree
.
├── go.mod
├── go.sum
├── main.go
└── zz_generated.ioc.go

0 directories, 4 files
```

#### 执行程序

`go run .`

控制台打印输出：

```sh
  ___    ___     ____                           _                         
 |_ _|  / _ \   / ___|           __ _    ___   | |   __ _   _ __     __ _ 
  | |  | | | | | |      _____   / _` |  / _ \  | |  / _` | | '_ \   / _` |
  | |  | |_| | | |___  |_____| | (_| | | (_) | | | | (_| | | | | | | (_| |
 |___|  \___/   \____|          \__, |  \___/  |_|  \__,_| |_| |_|  \__, |
                                |___/                               |___/ 
Welcome to use ioc-golang!
[Boot] Start to load ioc-golang config
[Config] Load default config file from ../conf/ioc_golang.yaml
[Config] Load ioc-golang config file failed. open /Users/laurence/Desktop/workplace/alibaba/conf/ioc_golang.yaml: no such file or directory
 The load procedure is continue
[Boot] Start to load debug
[Debug] Debug port is set to default :1999
[Boot] Start to load autowire
[Autowire Type] Found registered autowire type normal
[Autowire Struct Descriptor] Found type normal registered SD main.serviceStruct_
[Autowire Struct Descriptor] Found type normal registered SD main.app_
[Autowire Struct Descriptor] Found type normal registered SD main.serviceImpl1_
[Autowire Struct Descriptor] Found type normal registered SD main.serviceImpl2_
[Autowire Type] Found registered autowire type singleton
[Autowire Struct Descriptor] Found type singleton registered SD main.App
[Autowire Struct Descriptor] Found type singleton registered SD main.ServiceImpl1
[Autowire Struct Descriptor] Found type singleton registered SD main.ServiceImpl2
[Autowire Struct Descriptor] Found type singleton registered SD main.ServiceStruct
[Debug] Debug server listening at :1999
This is ServiceImpl1, hello laurence
This is ServiceImpl2, hello laurence
This is ServiceImpl1, hello laurence
This is ServiceStruct, hello laurence
...
```

可看到，依赖注入成功，程序正常运行。

### 调试程序

可看到打印出的日志中包含，说明 Debug 服务已经启动。

```bash
[Debug] Debug server listening at :1999
```

新开一个终端，使用 iocli 的调试功能，查看所有拥有代理层的结构和方法。默认端口为 1999。

```bash
% iocli list
main.ServiceImpl1
[GetHelloString]

main.ServiceImpl2
[GetHelloString]
```

监听方法的参数和返回值。以监听 main.ServiceImpl 结构的 GetHelloString 方法为例，每隔三秒钟，函数被调用两次，打印参数和返回值。

```bash
% iocli watch main.ServiceImpl1 GetHelloString
========== On Call ==========
main.ServiceImpl1.GetHelloString()
Param 1: (string) (len=8) "laurence"

========== On Response ==========
main.ServiceImpl1.GetHelloString()
Response 1: (string) (len=36) "This is ServiceImpl1, hello laurence"

========== On Call ==========
main.ServiceImpl1.GetHelloString()
Param 1: (string) (len=8) "laurence"

========== On Response ==========
main.ServiceImpl1.GetHelloString()
Response 1: (string) (len=36) "This is ServiceImpl1, hello laurence"
,,,
```



### 注解分析

```go
// +ioc:autowire=true
代码生成工具会识别到标有 +ioc:autowire=true 注解的对象

// +ioc:autowire:type=singleton
标记注入模型为 singleton 单例模型，还有 normal 多例模型等扩展

```

###  更多

更多代码生成注解可以移步[ioc-golang-cli](https://github.com/alibaba/IOC-golang/tree/master/iocli).查看。

可以移步 [ioc-golang-example](https://github.com/alibaba/IOC-golang/tree/master/example)  查看更多例子和高级使用方法。


### 证书

IOC-golang developed by Alibaba and licensed under the Apache License (Version 2.0).
See the NOTICE file for more information.

### 联系我们

感兴趣的开发者可以加入钉钉群：44638289