CURRENT_DIR=$(shell pwd)

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
server:
	go run cmd/main.go

create migrate:
	migrate create -ext sql -dir migrations -seq create_xxx_table

migrate up:
	migrate -source file://migrations/ -database postgres://postgres:sdy12197@localhost:5432/postsdb up

migrate down:
	migrate -source file://migrations/ -database postgres://postgres:sdy12197@localhost:5432/postsdb down