set GOOS=windows
set GOARCH=386
go build -ldflags="-s -w" -o ccc32.exe
if %errorlevel% neq 0 pause

set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o ccc64.exe
if %errorlevel% neq 0 pause
