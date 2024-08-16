# sqlbuilder-zero

**自动同步至 https://github.com/jzero-io/jzero/tree/main/.template/go-zero/model**

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
        ├──.............
```

## 特性

* 将原生 sql 全部替换为 sqlbuilder 进行构建
* 新增方法: BulkInsert: 批量插入数据
* 新增方法: FindByCondition: 基于不同条件查询数据
* 新增方法: FindOneByCondition: 基于不同条件查询数据, limit 1
* 新增方法: PageByCondition: 基于不同条件获取分页数据
* 新增方法: UpdateFieldsByCondition: 基于不同条件更新某几个 fields
* 新增方法: DeleteByCondition

## gen

```shell
jzero gen
```
