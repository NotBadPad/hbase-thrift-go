hbase-thrift-go
===============

在go中使用thrift连接hbase。
由于hbase官方提供的thrift文件使用的thrift.0.9.1生成的文件存在错误（应该是thrift工具的一个bug，thrift会把binary生成[]byte，然而该字段可能用作map的key，这在go中是不允许的，解决方法就是将binary改成string，似乎这个问题将会在0.9.2中解决），这里使用是个修改版的，详情见https://github.com/sdming/goh。
