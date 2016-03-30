all: single cores

single: single.go fractals.go
	go build $^

cores: cores.go fractals.go
	go build $^

bench: single cores
	time ./single >single.png
	time ./cores >cores.png
	md5sum *.png

clean:
	rm -rf single cores *.png
