.PHONY: build_cmd
build_cmd:
	cd cmd && \
		go build -o cmd . && \
		cd .. && \
		mv cmd/cmd build

.PHONY: generate
generate:
	@echo generating grpc code
	@protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./api/api.proto
