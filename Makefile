
# NOTE: 誤操作防止のためtarget指定なしの場合はエラー扱いにする
all:
	@echo Please specify the target. >&2
	@exit 1

inspect:
	docker run \
		--rm \
		-v ${PWD}:/local/src \
		openapitools/openapi-generator-cli \
		validate \
			-i /local/src/openapi/reference/root.yaml

gen:
	docker run \
		--rm \
		-v ${PWD}:/local/src \
		openapitools/openapi-generator-cli \
		generate \
			-i /local/src/openapi/reference/root.yaml \
			-g go-server \
			-o /local/src
