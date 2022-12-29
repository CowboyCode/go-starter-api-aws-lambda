create table actionlog
(
    id         int(11) unsigned auto_increment
        primary key,
    actiontype varchar(60)  default '' not null,
    message    text                    null,
    filename   varchar(255) default '' null,
    location   varchar(255) default '' null,
    line       varchar(50)             null,
    logtime    datetime                not null
)
    charset = utf8mb4;

create index type
    on actionlog (actiontype);