set GOPATH=%cd%
::  Run get-dependencies.bat before running this
::  script.

:: Build the library modules to be sure we still have a clean build
 
go build jutil/commandLineParser.go
go build jutil/consulUtil.go
go build jutil/fileUtil.go
go build jutil/timer.go
go build jutil/util.go