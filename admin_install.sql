
set explicit_defaults_for_timestamp = 1;

CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `salt` varchar(10) NOT NULL DEFAULT '',
  `sex` int(11) NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL DEFAULT '',
  `last_login` datetime DEFAULT NULL,
  `last_ip` varchar(15) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;


INSERT INTO `t_user` VALUES (1,'admin','e10adc3949ba59abbe56e057f20f883e','',1,'admin@admin.com','2019-07-27 11:49:37','127.0.0.1',0,'2019-07-27 03:49:36','2019-07-27 03:49:36');


CREATE TABLE `t_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_name` varchar(45) NOT NULL DEFAULT '' COMMENT '项目名',
  `project_describe` varchar(100) NOT NULL DEFAULT '' COMMENT '项目描述',
  `project_url` varchar(100) NOT NULL COMMENT '项目URL地址',
  `test_ip` varchar(15) NOT NULL COMMENT '测试环境IP',
  `release_ip` varchar(15) NOT NULL COMMENT '预发布IP',
  `pro_ip` varchar(15) NOT NULL COMMENT '生产环境IP',
  `is_monitor` tinyint(2) NOT NULL DEFAULT '1' COMMENT '1开启监控，2不监控，默认开启',
  `monitor_url` varchar(100) NOT NULL COMMENT '监控触发URL',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;




CREATE TABLE `admin`.`project_label`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) UNSIGNED NOT NULL DEFAULT '0',
  `label_name` varchar(20) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE = InnoDB;




CREATE TABLE `admin`.`project_api`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '项目关联ID',
  `project_label_id` int(11) UNSIGNED NOT NULL DEFAULT '0' COMMENT '项目标签关联ID',
	`api_name` varchar(20) NOT NULL DEFAULT '' COMMENT '接口名称',
	`api_url` varchar(50) NOT NULL DEFAULT '' COMMENT '接口URL,域名后台部分',
	`method` tinyint(1) NOT NULL DEFAULT '1' COMMENT '方法名称：1-GET 2-POST 3-PUT 4-PATCH 5-DELETE',
	`intro` varchar(255) NOT NULL DEFAULT '' COMMENT '接口简介',
	`status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0开发中，1上线使用，2停止使用',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
	KEY `project_id` (`project_id`),
	KEY `project_label_id` (`project_label_id`)
) ENGINE = InnoDB;
