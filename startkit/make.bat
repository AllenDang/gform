@echo off
windres resource.rc -o temp-rc.o 
go tool 8g -I %GOPATH%\pkg\windows_386 app.go
go tool pack grc _go_.8 app.8 temp-rc.o
go tool 8l -L %GOPATH%\pkg\windows_386 -o startkit.exe -s -Hwindowsgui _go_.8
del /F /Q *.8 *.o