



#项目文档
- 前端框架：[element-ui](https://github.com/ElemeFE/element)
- 后台框架：[gin](https://github.com/gin-gonic/gin)

## 1.基本介绍

### 1.1 项目介绍
[在线预览]()
> goSystemManage 是一个基于vue 和 go 语言 开发出一套 前后端分离的后台管理系统，其目的是 为了快速搭建一套后台管理系统，并集成jwt鉴权、动态路由、动态表单、表单生成器、代码生成器等功能，提供快速且便利的后台管理。

###1.2 项目结构

````
  │─docker            (docker服务目录)
  │─docs              (文件目录) 
  │─sql               (数据库脚本)   
  ├─goAdmin           （后端文件夹）
  │  ├─base           （基础内核）
  │  ├─config         （配置包）
  │  ├─ctrl           （API接口）
  │  ├─gateway        （路由网关）
  │  ├─global         （全局对象）
  │  ├─handler        （过滤拦截）
  │  ├─init           （初始化）
  │  ├─model          （结构体层）
  │  ├─resource       （静态资源包）
  │  ├─service         (服务)
  │  └─util           （工具包）
  └─vueWeb            （前端文件）
      ├─static        （发布模板）
      └─src           （源码包）
````

## 2.功能介绍
>   本项目主要包含分为两个部分：一：系统配置，其中主要包含：用户管理，菜单管理，角色管理；二：代码生成器，其中包含：数据库表结构生成器、代码生成器、表单生成器等部分；
>通过这部分内容进行定制开发。
### 2.1 主要功能


## 3.技术选型
- 前端 ： 用基于`vue` 的`Element-UI` 构建基础页面
- 后端 ： 基于 `go` 语言 搭建 restful 风格的web框架
- 数据库 ： 采用MySQL ，并借助`gorm` 实现对数据库的基本操作
- 缓存 ： 基于 `redis` 管理 当前活跃用户的 `jwt`令牌
