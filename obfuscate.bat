set GOOS=windows
bash -c "cp -r package backup"
XCOPY resources\gomambojambo.exe gomambojambo.exe /Y
REM Obfuscating Source Code
.\gomambojambo.exe -calls -deadcode -loops -srcpath .\package\ -strings -verbose -writechanges

del gomambojambo.exe
REM Building: DEBUG Client
set GOARCH=amd64

REM set GOOS=linux
REM go build -o Shaman.nix -ldflags="-w -s" ./package/main.go

REM set GOOS=darwin
REM go build -o Shaman.mac -ldflags="-w -s" ./package/main.go
