exe: clean
	windres -o controls-rc.o controls.rc
	8g controls.go
	go tool pack grc _go_.8 controls.8 controls-rc.o
	8l -o controls.exe _go_.8
	rm *.8 *.o
clean:
	rm -f *.8 *.o *.exe
