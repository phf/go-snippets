ALL=single cores pixels

all: $(ALL)

single: single.go fractals.go
	go build $^

cores: cores.go fractals.go
	go build $^

pixels: pixels.go fractals.go
	go build $^

bench: $(ALL)
	./xtime ./single >single.out.png
	./xtime ./cores >cores.out.png
	./xtime ./pixels >pixels.out.png
	diff --from-file *.png

clean:
	rm -rf single cores pixels *.out.png
