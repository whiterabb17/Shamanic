set GOOS=windows
bash -c "cp -r package backup"
XCOPY resources\gomambojambo.exe gomambojambo.exe /Y
REM Obfuscating Source Code
.\gomambojambo.exe -calls -deadcode -loops -srcpath .\package\ -strings -verbose -writechanges

del gomambojambo.exe
REM Building: DEBUG Client
set GOARCH=amd64
go build -o ShamanCli.exe -ldflags="-w -s" ./package/main.go

go build -o Shaman.exe -ldflags="-w -s -H=windowsgui" ./package/main.go

REM set GOOS=linux
REM go build -o Shaman.nix -ldflags="-w -s" ./package/main.go

REM set GOOS=darwin
REM go build -o Shaman.mac -ldflags="-w -s" ./package/main.go
