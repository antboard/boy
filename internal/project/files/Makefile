NAME=example
# 基本git信息采集
GIT_COMMIT=$(shell git rev-parse --short HEAD)
GIT_BRANCH=$(shell git name-rev --name-only HEAD)
GIT_TAG=$(shell git describe --abbrev=0 --tags --always --match "v*")
# 时间
BUILD_DATE=$(shell date +%s)
# 编译
GO_VERSION=$(shell go version)
# 路径
CurPath=$(shell pwd)
all:
	cd ${CurPath}/cmd/service && \
	go build -ldflags                           \
	"                                           \
	-X \"main.Commit=${GIT_COMMIT}\" \
	-X 'main.BuildDate=${BUILD_DATE}' \
	-X 'main.GitTag=${GIT_TAG}' \
	-X 'main.GitBranch=${GIT_BRANCH}' \
	-X 'main.GoVersion=${GO_VERSION}' \
	-w -s                               \
	"                                           \
	-o ${CurPath}/output/${NAME} .

run:
	cd output && ./${NAME}

clean:
	rm -rf output && mkdir output
