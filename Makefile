all: proto build_cmd build_ui 

proto:
	@echo generating grpc code
	@protoc \
		--go_out=. \
		--go_opt=paths=source_relative \
		--go-grpc_out=. \
		--go-grpc_opt=paths=source_relative \
		./api/api.proto

build_cmd:
	@echo building cmd
	@cd cmd && \
		go build -o cmd . && \
		cd .. && \
		mv cmd/cmd build

build_ui:
	@echo building ui
	@cd ui && \
		go build -o ui . && \
		cd .. && \
		mv ui/ui build

.PHONY: proto build_cmd build_ui