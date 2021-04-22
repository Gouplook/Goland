

#====== 开发备忘录=====



#======mysql 安装备忘录========
A "/etc/my.cnf" from another install may interfere with a Homebrew-built
server starting up correctly.

MySQL is configured to only allow connections from localhost by default

To connect:
    mysql -uroot

mysql@5.6 is keg-only, which means it was not symlinked into /usr/local,
because this is an alternate version of another formula.

If you need to have mysql@5.6 first in your PATH, run:
  echo 'export PATH="/usr/local/opt/mysql@5.6/bin:$PATH"' >> ~/.zshrc

For compilers to find mysql@5.6 you may need to set:
  export LDFLAGS="-L/usr/local/opt/mysql@5.6/lib"
  export CPPFLAGS="-I/usr/local/opt/mysql@5.6/include"


To have launchd start mysql@5.6 now and restart at login:
  brew services start mysql@5.6
Or, if you don't want/need a background service you can just run:
  /usr/local/opt/mysql@5.6/bin/mysql.server start
  
  
  621778 8362401349705



CREATE TABLE `kc_tag` (
`tag_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
`bus_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属企业商户id',
`name` varchar(20) NOT NULL DEFAULT '' COMMENT '标签名称',
`id_del` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否被删除 0=否 1=是',
`ctime` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
PRIMARY KEY (`tag_id`),
KEY `bus_id` (`bus_id`,`id_del`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=84 DEFAULT CHARSET=utf8mb4 COMMENT='卡协服务-单项目可用标签';
