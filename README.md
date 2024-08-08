# sqlbuilder-zero

go-zero sqlbuilder model template.

把 goctl 默认的 model 模板替换成 github.com/huandu/go-sqlbuilder 方式, 目的: 

* 为了支持不同的数据库, 如达梦, 人大金仓等.
* 基于 sqlbuilder 可以编写更多通用的方法, 提升开发效率

## Usage

```shell
mkdir -p .template/go-zero
```

将本仓库的 .template/go-zero/model 目录文件复制到项目 .template/go-zero 下

目录结构

```shell
.template
└── go-zero
    └── model
        ├── customized.tpl
        ├── delete.tpl
        ├── err.tpl
        ├── field.tpl
```

## gen

```shell
jzero gen
```
