CURRENT_DIR=$(shell pwd)

pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge

build:
	CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go

proto-gen:
	bash ${CURRENT_DIR}/script/gen-proto.sh
	ls genproto/*.pb.go | xargs -n1 -IX bash -c "sed -e '/bool/ s/,omitempty//' X > X.tmp && mv X{.tmp,}"
server:
	go run cmd/main.go
	
migrate_up:
	migrate -path migrations/ -database postgres://postgres:sdy12197@localhost:5432/reatingdb up

migrate_down:
	migrate -path migrations/ -database postgres://postgres:sdy12197@localhost:5432/reatingdb down

migrate_force:
	migrate -path migrations/ -database postgres://postgres:sdy12197@localhost:5432/reatingdb force 2
