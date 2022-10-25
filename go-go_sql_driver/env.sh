gopath=$(cd `dirname $0`;pwd)
mkdir src
echo "export GOPATH="$gopath"" >> ~/.bash_profile
source ~/.bash_profile
go env -w GO111MODULE=auto
export GO111MODULE=auto
export GOPROXY=https://goproxy.io
go get github.com/go-sql-driver/mysql
#go install github.com/go-sql-driver/mysql@latest
