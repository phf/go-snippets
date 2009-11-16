TESTFILE=/usr/share/dict/cracklib-small
GOCC=6g
GOLD=6l

all: freq simple_freq complex_freq

freq: freq.c

simple_freq: simple_freq.6
	6l -o simple_freq simple_freq.6
simple_freq.6: simple_freq.go
	6g simple_freq.go

complex_freq: complex_freq.6
	6l -o complex_freq complex_freq.6
complex_freq.6: complex_freq.go
	6g complex_freq.go

clean:
	rm -rf *.6 *.8 freq simple_freq complex_freq *.out

test:
	time ./freq <${TESTFILE} >freq.out
	time ./simple_freq <${TESTFILE} >simple_freq.out
	time ./complex_freq <${TESTFILE} >complex_freq.out
	time python freq.py <${TESTFILE} >python_freq.out
	diff freq.out simple_freq.out
	diff simple_freq.out complex_freq.out
	diff complex_freq.out python_freq.out
