# Casbin RBAC for REST API
## ルール
- ロールはadmin, writer, readerの3つ
- adminはすべてのAPIにアクセスできる
- writerは一般向けAPIにアクセスできる
- readerは一般向けの参照系APIにアクセスできる

## 結果
```shell
curl http://localhost:1323/admin/user/list\?user_name\=alice
admin user list
curl http://localhost:1323/device/list\?user_name\=alice
device list
curl -X DELETE http://localhost:1323/device/1\?user_name\=alice
device delete
curl http://localhost:1323/admin/user/list\?user_name\=bob
{"message":"Forbidden"}
curl http://localhost:1323/device/list\?user_name\=bob
device list
curl -X DELETE http://localhost:1323/device/1\?user_name\=bob
device delete
curl http://localhost:1323/admin/user/list\?user_name\=john
{"message":"Forbidden"}
curl http://localhost:1323/device/list\?user_name\=john
device list
curl -X DELETE http://localhost:1323/device/1\?user_name\=john
{"message":"Forbidden"}
```