CREATE TABLE IF NOT EXISTS nid_node_id
(
    id                bigint       NOT NULL AUTO_INCREMENT COMMENT 'ID',
    created_time      datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_time      datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    uuid              varchar(36)  NOT NULL DEFAULT '' COMMENT 'UUID',
    instance_id       varchar(255) NOT NULL DEFAULT '' COMMENT '实例ID',
    instance_name     varchar(255) NOT NULL DEFAULT '' COMMENT '实例名称',
    instance_metadata json         NOT NULL COMMENT '实例元数据',
    node_id           bigint       NOT NULL DEFAULT 0 COMMENT '节点id',
    node_id_status    mediumint    NOT NULL DEFAULT 0 COMMENT '节点状态',
    expired_at        datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '失效时间',
    PRIMARY KEY (id),
    UNIQUE INDEX (uuid),
    INDEX (instance_id, node_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
    COMMENT = '节点';

CREATE TABLE IF NOT EXISTS nid_node_serial
(
    id              bigint       NOT NULL AUTO_INCREMENT COMMENT 'ID',
    created_time    datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    updated_time    datetime(3)  NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    instance_id     varchar(255) NOT NULL DEFAULT '' COMMENT '实例ID',
    current_node_id bigint       NOT NULL DEFAULT 0 COMMENT '当前节点id',
    PRIMARY KEY (id),
    UNIQUE INDEX (instance_id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4
    COMMENT = '节点系列号';