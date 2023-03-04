status:
	@docker ps -a --format "table {{.ID}}\t{{.Names}}\t{{.Status}}" | grep notification

swag-doc:
	swag init -g main.go

test:
	go test ./api/handler/... -count=1

test-create:
	docker run --name notification-db \
    		-p 3306:3306 \
    		-e MYSQL_ROOT_PASSWORD=notification1235 \
    		-e MYSQL_DATABASE=notification \
    		-d mysql:8 --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
	docker run --name notification-redis \
			-p 6379:6379 \
			-v $(CURDIR)/configs/test-redis.conf:/redis.conf \
			-d redis redis-server /redis.conf
	go test github.com/ismdeep/notification/testing/waitdbup -count=1
	go test github.com/ismdeep/notification/testing/prepare  -count=1

test-up:
	docker start notification-db
	docker start notification-redis

test-down:
	docker stop notification-db
	docker stop notification-redis

test-clean: test-down
	docker rm notification-db
	docker rm notification-redis

test-renew:
	-make test-clean
	-make test-create

pre-push:
	go mod tidy
	make test
