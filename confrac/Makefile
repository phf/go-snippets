ALL=single cores pixels

all: $(ALL)

single: single.go fractals.go
	go build $^

cores: cores.go fractals.go
	go build $^

pixels: pixels.go fractals.go
	go build $^

bench: $(ALL)
	./xtime ./single >single.png
	./xtime ./cores >cores.png
	./xtime ./pixels >pixels.png
	diff --from-file *.png

clean:
	rm -rf single cores pixels *.png
