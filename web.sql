DROP TABLE IF EXISTS `web_news`;
CREATE TABLE `web_news` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '新闻标题',
  `created_on` varchar(100) DEFAULT '' COMMENT '时间',
  `content` text COMMENT '内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='新闻管理';
select * from web_news;



DROP TABLE IF EXISTS `web_member`;
CREATE TABLE `web_member` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `identity` varchar(100) DEFAULT '' COMMENT '身份',
  `name` varchar(100) DEFAULT '' COMMENT '姓名',
  `phone` varchar(100) DEFAULT '' COMMENT '电话',
  `mail` varchar(100) DEFAULT '' COMMENT '邮箱',
  `research` varchar(100) DEFAULT '' COMMENT '研究方向',
  `achievement` text COMMENT '成果',
  `introduction` text COMMENT '内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='成员管理';
select * from web_member;

DROP TABLE IF EXISTS `web_image`;
CREATE TABLE `web_image` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `address` text COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='图片管理';
select * from web_image;

DROP TABLE IF EXISTS `web_achievement`;
CREATE TABLE `web_achievement` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '名称',
  `category` varchar(100) DEFAULT '' COMMENT '类别',
  `address` varchar(100) DEFAULT '' COMMENT '地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='成果管理';
select * from web_achievement;




DROP TABLE IF EXISTS `web_manager`;
CREATE TABLE `web_manager` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) DEFAULT '' COMMENT '账号',
  `password` varchar(50) DEFAULT '' COMMENT '密码',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

INSERT INTO `web_manager` (`id`, `username`, `password`) VALUES ('1', '1', '1');


DROP TABLE IF EXISTS `web_article`;
CREATE TABLE `web_article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(100) DEFAULT '' COMMENT '论文标题',
  `journal` varchar(100) DEFAULT '' COMMENT '期刊',
  `author` varchar(100) DEFAULT '' COMMENT '第一作者',
  `authors` varchar(100) DEFAULT '' COMMENT '其他作者',
  `date` varchar(100) DEFAULT '' COMMENT '时间',
  `link` varchar(100) DEFAULT '' '详情页链接',
  `papercode` varchar(100) DEFAULT '' '代码',
  `theyear` varchar(10) DEFAULT '' COMMENT '论文年份',
  `abstract` text COMMENT '摘要',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='论文管理';
select * from web_article;


DROP TABLE IF EXISTS `web_project`;
CREATE TABLE `web_project` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT '' COMMENT '项目名称',
  `link` varchar(100) DEFAULT '' COMMENT '项目链接',
  `theyear` varchar(10) DEFAULT '' COMMENT '项目年份',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
select * from web_project;