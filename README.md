[![](https://img.shields.io/badge/GreatSQL-官网-orange.svg)](https://greatsql.cn/)
[![](https://img.shields.io/badge/GreatSQL-论坛-brightgreen.svg)](https://greatsql.cn/forum.php)
[![](https://img.shields.io/badge/GreatSQL-博客-brightgreen.svg)](https://greatsql.cn/home.php?mod=space&uid=10&do=blog&view=me&from=space)
[![](https://img.shields.io/badge/License-Apache_v2.0-blue.svg)](https://gitee.com/GreatSQL/GreatSQL/blob/master/LICENSE)
[![](https://img.shields.io/badge/release-1.2.1-blue.svg)](https://gitee.com/GreatSQL/gt-checksum/releases/tag/1.2.1)
![输入图片说明](Menu/%E5%BE%AE%E4%BF%A1%E5%9B%BE%E7%89%87_20230329094139.png)
# 数据库巡检工具 DB OSInspection
## 简介：
DB OSInspection 数据库巡检工具，实现对数据库一键巡检，对实例有大概了解。
## 愿景
一秒了解数据库
## 快速上手
```go
./DB_OSInspection -u 用户名 -p 密码  
```
即可实现全表巡检
### 参数解析
所有输入后面都需要跟空格，否则会错误
- -u 
输入用户名  默认为root
- -p 
输入密码
- -nw 
输入链接模式  默认为tcp
- -P 
输入IP号和端口号  默认为localhost:3306
- -m  
模式选择all全部巡检、table只巡检表、index只巡检索引、variables只巡检重要参数、status只巡检重要状态、user只巡检用户、privileges只巡检权限
- -v
输出版本号
- -h
输出帮助
## 已实现功能
### 实现表巡检
- 大小超过10G的表
- 索引超过6个的表
- 碎片率超过50%的表
- 行数超过1000万行的表
- 非默认字符集的表
- 含有大字段的表
- varchar定义超长的表
- 无主键/索引的表
### 索引巡检
- 重复索引
- 索引列超过5个的索引
- 无用索引
### 重要参数
- version
- innodb_buffer_pool_size
- innodb_flush_log_at_trx_commit
- innodb_log_file_size
- innodb_log_files_in_group
- innodb_file_per_table
- innodb_max_dirty_pages_pct
- sync_binlog
- max_connections
- table_open_cache
- table_definition_cache
### 重要状态指标
- Uptime
- Opened_files
- Opened_table_definitions
- Opened_tables
- Max_used_connections
- Threads_created
- Threads_connected
- Aborted_connects
- Aborted_clients
- Table_locks_waited
- Innodb_buffer_pool_wait_free
- Innodb_log_waits
- Table_locks_waited
- Innodb_row_lock_waits
- Innodb_row_lock_time_avg
- Binlog_cache_disk_use
- Created_tmp_disk_tables
### 用户检查
- 无密码用户
- %用户
### 权限检查