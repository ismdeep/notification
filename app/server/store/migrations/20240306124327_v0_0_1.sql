-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

create table msgs
(
    id                   varchar(255)    not null
        primary key,
    customer_msg_id      varchar(255)    not null,
    user_id              varchar(255)    not null,
    content              longtext        null,
    status               tinyint         not null,
    err                  longtext        null,
    created_at_unix_nano bigint unsigned not null,
    constraint uq_userid_customermsgid
        unique (customer_msg_id, user_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

create index idx_msgs_created_at_unix_nano
    on msgs (created_at_unix_nano);

create table projects
(
    id                   varchar(255)    not null
        primary key,
    name                 varchar(255)    not null,
    created_at_unix_nano bigint unsigned not null,
    constraint uni_projects_name
        unique (name)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

create table tokens
(
    id                   varchar(255)    not null
        primary key,
    user_id              longtext        null,
    token_name           varchar(255)    not null,
    token                varchar(255)    not null,
    aes_key              varchar(255)    not null,
    created_at_unix_nano bigint unsigned not null
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

create table users
(
    id                   varchar(255)      not null
        primary key,
    username             varchar(50)       not null,
    nickname             varchar(255)      not null,
    avatar               varchar(255)      not null,
    enabled              tinyint default 0 not null,
    digest               varchar(255)      not null,
    created_at_unix_nano bigint unsigned   not null,
    constraint uni_users_username
        unique (username)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;



-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
