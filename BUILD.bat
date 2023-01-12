REM Removing Stale Builds
del Shaman.exe
del ShamanCLI.exe
cd package
set GOARCH=amd64
REM Building: DEBUG Client
go build -o ../ShamanCLI.exe -ldflags="-w -s"

REM Building: PRODUCTION Client
go build -o ../Shaman.exe -ldflags="-w -s -H=windowsgui"

REM Complete
cd ..