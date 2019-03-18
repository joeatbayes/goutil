export GOPATH=$PWD
#:  Run get-dependencies before running this
#:  script.

# Build the library modules to be sure we still have a clean build
# 
go build jutil/commandLineParser.go
go build jutil/consulUtil.go
go build jutil/timer.go
go build jutil/util.go
go build jutil/fileUtil.go



