
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
  `create_time` timestamp(6) NULL DEFAULT NULL,
  `update_time` timestamp(6) NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;