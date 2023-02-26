::windows
SET CGO_ENABLED=0
SET GOOS=windows
SET GOARCH=amd64
go build -o go-build.exe main.go

::linux
@REM SET CGO_ENABLED=0
@REM SET GOOS=linux
@REM SET GOARCH=amd64
@REM go build -o go-build main.go
