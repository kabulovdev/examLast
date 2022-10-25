pull_submodule:
	git submodule update --init --recursive

update_submodule:
	git submodule update --remote --merge
run:
	go run cmd/main.go

run_script:
	./scripts/gen-proto.sh

swag:
	swag init -g ./api/router.go -o api/docs

	