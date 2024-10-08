CREATE TABLE IF NOT EXISTS nid_node_id
(
    id                bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    created_time      datetime(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
    updated_time      datetime(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
    instance_name     varchar(255)    NOT NULL DEFAULT '' COMMENT '实例名称',
    instance_id       varchar(255)    NOT NULL DEFAULT '' COMMENT '实例ID',
    node_id           bigint          NOT NULL DEFAULT 0 COMMENT '节点id',
    node_id_status    mediumint       NOT NULL DEFAULT 0 COMMENT '节点状态',
    instance_metadata json            NOT NULL COMMENT '实例元数据',
    expired_at        datetime(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '失效时间',
    access_token      varchar(255)    NOT NULL DEFAULT '' COMMENT '令牌；用于续订和释放ID',
    PRIMARY KEY (id),
    UNIQUE INDEX (instance_id, node_id),
    INDEX (expired_at)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
    COMMENT = '节点';

CREATE TABLE IF NOT EXISTS nid_node_serial
(
    id              bigint unsigned NOT NULL AUTO_INCREMENT COMMENT 'ID',
    created_time    datetime(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP() COMMENT '创建时间',
    updated_time    datetime(3)     NOT NULL DEFAULT CURRENT_TIMESTAMP() ON UPDATE CURRENT_TIMESTAMP() COMMENT '更新时间',
    instance_id     varchar(255)    NOT NULL DEFAULT '' COMMENT '实例ID',
    current_node_id bigint          NOT NULL DEFAULT 0 COMMENT '当前节点id',
    PRIMARY KEY (id),
    UNIQUE INDEX (instance_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
    COMMENT = '节点系列号';
