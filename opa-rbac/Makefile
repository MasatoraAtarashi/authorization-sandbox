test: test-alice test-bob test-john

test-alice:
	curl http://localhost:1323/articles/ -H 'user_name: alice'
	curl http://localhost:1323/articles/1 -H 'user_name: alice'
	curl -X POST http://localhost:1323/articles/ -H 'user_name: alice'
	curl -X PUT http://localhost:1323/articles/1 -H 'user_name: alice'
	curl -X DELETE http://localhost:1323/articles/1 -H 'user_name: alice'

test-bob:
	curl http://localhost:1323/articles/ -H 'user_name: bob'
	curl http://localhost:1323/articles/1 -H 'user_name: bob'
	curl -X POST http://localhost:1323/articles/ -H 'user_name: bob'
	curl -X PUT http://localhost:1323/articles/1 -H 'user_name: bob'
	curl -X DELETE http://localhost:1323/articles/1 -H 'user_name: bob'

test-john:
	curl http://localhost:1323/articles/ -H 'user_name: john'
	curl http://localhost:1323/articles/1 -H 'user_name: john'
	curl -X POST http://localhost:1323/articles/ -H 'user_name: john'
	curl -X PUT http://localhost:1323/articles/1 -H 'user_name: john'
	curl -X DELETE http://localhost:1323/articles/1 -H 'user_name: john'