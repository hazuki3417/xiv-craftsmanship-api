# 概要: docker imageビルドを行うタスクランナー

VERSION = latest
ENVIRONMENT = develop
DOCKER_BASE_DIR = docker
DOCKER_IMAGE_NAME = xiv-craftsmanship-api
REGISTRY =

DOCKER_IMAGE_TAG = ${VERSION}-${ENVIRONMENT}
DOCKER_CONTEXT_PATH = ${realpath ../../../}
DOCKER_FILE_PATH = ${DOCKER_BASE_DIR}/${ENVIRONMENT}/${DOCKER_IMAGE_NAME}/Dockerfile
DOCKER_FILE_FULL_PATH = ${DOCKER_CONTEXT_PATH}/${DOCKER_FILE_PATH}


# NOTE: 誤操作防止のためtarget指定なしの場合はエラー扱いにする
all:
	@echo Please specify the target
	@exit 1

################################################################################
# タスク
################################################################################

# NOTE: マルチプラットフォーム未対応のビルド
# usage: make build VERSION={versionName}
build:
	docker build \
		-t ${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} \
		-f ${DOCKER_FILE_FULL_PATH} \
		${DOCKER_CONTEXT_PATH}

# NOTE: マルチプラットフォーム対応のビルド・デプロイ
#       マルチプラットフォームビルドする場合は以下の制約が発生します
#       - マルチプラットフォームビルドをする場合、ビルド後のimageはlocal registryに保存できない（仕様）
#       - マルチプラットフォームビルドの成果物としてdocker imageは指定できない
#       - マルチプラットフォームビルドの成果物としてtarは指定できるが、必要なファイルが不足しているためloadに失敗する
build-deploy:
	docker buildx build \
		-t ${REGISTRY}/${DOCKER_IMAGE_NAME}:${DOCKER_IMAGE_TAG} \
		-f ${DOCKER_FILE_FULL_PATH} \
		--platform linux/amd64,linux/arm64 \
		--push \
		${DOCKER_CONTEXT_PATH}
