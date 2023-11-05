.PHONY: run test

run:
	@echo "Starting the server..."
	@go run main.go

test: test-admin test-doctor test-nurse test-family

test-admin:
	curl 'http://localhost:1323/admin/user/list' -H 'user: admin'
	curl 'http://localhost:1323/patient/123/record' -H 'user: admin'
	curl 'http://localhost:1323/patient/124/record' -H 'user: admin'
	curl 'http://localhost:1323/patient/123/record/summary' -H 'user: admin'
	curl 'http://localhost:1323/patient/124/record/summary' -H 'user: admin'

test-doctor:
	curl 'http://localhost:1323/admin/user/list' -H 'user: doctor'
	curl 'http://localhost:1323/patient/123/record' -H 'user: doctor'
	curl 'http://localhost:1323/patient/124/record' -H 'user: doctor'
	curl 'http://localhost:1323/patient/123/record/summary' -H 'user: doctor'
	curl 'http://localhost:1323/patient/124/record/summary' -H 'user: doctor'

test-nurse:
	curl 'http://localhost:1323/admin/user/list' -H 'user: nurse'
	curl 'http://localhost:1323/patient/123/record' -H 'user: nurse'
	curl 'http://localhost:1323/patient/124/record' -H 'user: nurse'
	curl 'http://localhost:1323/patient/123/record/summary' -H 'user: nurse'
	curl 'http://localhost:1323/patient/124/record/summary' -H 'user: nurse'

test-family:
	curl 'http://localhost:1323/admin/user/list' -H 'user: family'
	curl 'http://localhost:1323/patient/1/record' -H 'user: family'
	curl 'http://localhost:1323/patient/123/record' -H 'user: family'
