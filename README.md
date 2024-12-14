# 描述

goleaf 分布式ID生成系统，基于 go-zero 实现，可根据Snowflake算法和号段模式生成唯一ID

# 架构

![这是图片](https://telegraph-image-92x.pages.dev/file/76226f1bd62a0cf270ebc-a2eea87d05b0259094.png "Magic Gardens")

- goleaf：核心服务，生成分布式id
- sdk：goleaf 服务提供sdk 获取id
- mysql：数据库，业务信息存储
- etcd：获取node id，Snowflake算法使用
- redis：缓存服务器

# 部署

- 前置条件
  - etcd 部署
  - mysql 部署
  - redis 部署

## 数据库表创建
```sql
create table leaf.leaf_alloc
(
    id int auto_increment
        primary key,
    biz_tag varchar(128) default '' not null,
    max_id bigint default 1 not null,
    step int not null,
    description varchar(256) default '' not null ,
    created_at  timestamp default current_timestamp not null,
    updated_at timestamp default current_timestamp not null,
    deleted_at timestamp default null
);

create index leaf_alloc_biz_tag_index
    on leaf.leaf_alloc (biz_tag);
```

## 编译

- make build: 编译，二进制文件生成在 ./bin 目录下
- 服务启动：./bin/goleaf -f etc/goleaf-api.yaml

# 工具

- make swagger: 生成 swagger 文档
- make all: api 文档格式化和生成go代码、swagger文档
- make build：编译，二进制文件生成在 ./bin 目录下