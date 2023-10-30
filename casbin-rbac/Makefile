run: alice bob john

alice:
	curl http://localhost:1323/admin/user/list\?user_name\=alice
	curl http://localhost:1323/device/list\?user_name\=alice
	curl -X DELETE http://localhost:1323/device/1\?user_name\=alice

bob:
	curl http://localhost:1323/admin/user/list\?user_name\=bob
	curl http://localhost:1323/device/list\?user_name\=bob
	curl -X DELETE http://localhost:1323/device/1\?user_name\=bob

john:
	curl http://localhost:1323/admin/user/list\?user_name\=john
	curl http://localhost:1323/device/list\?user_name\=john
	curl -X DELETE http://localhost:1323/device/1\?user_name\=john