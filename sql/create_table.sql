create database oj;
use oj;
-- 用户表
CREATE TABLE `oj_user`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT COMMENT '',
    `user_id`    VARCHAR(255) NOT NULL COMMENT '用户id',
    `username`   VARCHAR(50)  NOT NULL COMMENT '',
    `password`   VARCHAR(255) NOT NULL COMMENT '',
    `user_role`  VARCHAR(255) NOT NULL COMMENT '',
    `email`      VARCHAR(50)  NOT NULL COMMENT '',
    `phone`      VARCHAR(255) NOT NULL COMMENT '',
    `createTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updateTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`  INT      DEFAULT 0 COMMENT '逻辑删除标志（0: 未删除，1: 已删除）'
);


CREATE UNIQUE INDEX idx_user_id ON oj_user (user_id);
-- 分类表
CREATE TABLE `oj_type`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT COMMENT '',
    `parent_id`  VARCHAR(255) NOT NULL COMMENT '',
    `tag_name`   VARCHAR(50)  NOT NULL COMMENT '',
    `createTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updateTime` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`   INT      DEFAULT 0 COMMENT '逻辑删除标志（0: 未删除，1: 已删除）'
);

CREATE UNIQUE INDEX idx_parent_id ON oj_type (parent_id);
CREATE INDEX idx_tag_name ON oj_type (tag_name);
-- 问题表
CREATE TABLE `oj_question`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT COMMENT '',
    `question_id` VARCHAR(255) NOT NULL COMMENT '问题id',
    `input`       VARCHAR(255) NOT NULL COMMENT '输入',
    `output`      VARCHAR(255) NOT NULL COMMENT '输出',
    `title`       VARCHAR(255) NOT NULL COMMENT '标题',
    `content`     VARCHAR(255) NOT NULL COMMENT '文本内容',
    `tags`        VARCHAR(255) NOT NULL COMMENT '标签',
    `submitNum`   INT          NOT NULL COMMENT '提交总数',
    `passNum`     INT          NOT NULL COMMENT '通过总数',
    `max_mem`     INT          NOT NULL COMMENT '最大内存',
    `max_runtime` INT          NOT NULL COMMENT '最大运行时间',
    `createTime`  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updateTime`  DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`    INT      DEFAULT 0 COMMENT '逻辑删除标志（0: 未删除，1: 已删除）'
);
CREATE UNIQUE INDEX idx_question_id ON oj_question (question_id);
-- 提交表
CREATE TABLE `oj_question_submit_info`
(
    `id`                 INT PRIMARY KEY AUTO_INCREMENT COMMENT '',
    `question_submit_id` VARCHAR(255) NOT NULL COMMENT '问题提交id',
    `language`           VARCHAR(255) NOT NULL COMMENT '',
    `code`               VARCHAR(255) NOT NULL COMMENT '',
    `status`             VARCHAR(255) NOT NULL COMMENT '',
    `user_id`            VARCHAR(255) NOT NULL COMMENT '',
    `question_id`        VARCHAR(255) NOT NULL COMMENT '',
    `judge_info`         VARCHAR(255) NOT NULL COMMENT '',
    `createTime`         DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updateTime`         DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
    `isDelete`           INT      DEFAULT 0 COMMENT '逻辑删除标志（0: 未删除，1: 已删除）',
);
CREATE UNIQUE INDEX idx_question_submit_id ON oj_question_submit_info (question_submit_id);

