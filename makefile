REMOTE_HOST := 
APP_NAME := longplan-backend

dbcreate:
	migrate create -ext sql -dir cmd/db/migrations/ -seq longplan
dbmigrate:
	method=migrate go run ./cmd/db/
dbupgrade:
	method=up go run ./cmd/db/
dbdowngrade:
	method=down go run ./cmd/db/
dbdrop:
	method=drop go run ./cmd/db/
dbreset:
	method=reset go run ./cmd/db/

stopair:
	-kill $$(lsof -ti:8000) 2>/dev/null
