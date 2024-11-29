
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