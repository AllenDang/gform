exe: clean
exe: clean
	windres -o dialog-rc.o dialog.rc
	8g dialog.go
	gopack grc _go_.8 dialog.8 dialog-rc.o
	8l -o dialog.exe _go_.8
	rm *.8 *.o
clean:
	rm -f *.8 *.o *.exe