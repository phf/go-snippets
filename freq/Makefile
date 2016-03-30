TESTFILE=/usr/share/dict/web2
ALL=freq simple_freq complex_freq
CFLAGS=-std=c11 -Wall -Wextra -Wpedantic -O2

all: $(ALL)

freq: freq.c

simple_freq: simple_freq.go
	go build $^
complex_freq: complex_freq.go
	go build $^

clean:
	rm -rf $(ALL) *.out

bench: $(ALL)
	./xtime ./freq <${TESTFILE} >freq.out
	./xtime ./simple_freq <${TESTFILE} >simple_freq.out
	./xtime ./complex_freq <${TESTFILE} >complex_freq.out
	./xtime python freq.py <${TESTFILE} >python_freq.out
	diff freq.out simple_freq.out
	diff simple_freq.out complex_freq.out
	diff complex_freq.out python_freq.out
