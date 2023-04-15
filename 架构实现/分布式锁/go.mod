module algorithm/distributeLock

go 1.20

//这里会有包的一些冲突, 需要使用go get -u获取最新的版本
//当使用了go modules的情况下，不建议使用"go get"命令。而应该使用"go mod"命令来管理依赖。
//在go modules中，可以通过修改go.mod文件来升级依赖库的版本。因为"go get"命令可能会导致依赖库的版本混乱
require (
	github.com/go-redsync/redsync v1.4.2
	github.com/gomodule/redigo v2.0.0+incompatible
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
)
