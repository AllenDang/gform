exe: clean
	windres -o resource.o resource.rc
	8g testcanvas.go
	gopack grc _go_.8 testcanvas.8 resource.o
	8l -o testcanvas.exe -s _go_.8
	rm *.8 *.o
clean:
	rm -f *.8 *.o *.exe