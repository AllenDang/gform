exe: clean
exe: clean
	windres -o dialog-rc.o dialog.rc
	8g dialog.go
	gopack grc _go_.8 dialog.8 dialog-rc.o
	8l -o dialog.exe -Hwindowsgui _go_.8
	mt.exe -manifest dialog.manifest -outputresource:dialog.exe
	rm *.8 *.o
clean:
	rm -f *.8 *.o *.exe