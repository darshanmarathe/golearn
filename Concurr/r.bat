cls
go build %1.go
%1.exe
pause
del *.exe
r %1