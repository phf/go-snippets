# Comparing Freqs

There are four implementations of `freq` here, two in Go, one in Python,
and the original C code; I've added the C version simply for convenience,
I hope this is "fair use" of the code. The `Makefile` has a `bench` target
that makes sure the programs agree in their output and also times them for
comparison.

## Findings 2010

As it stands, the Go implementations perform worse than the C code but
better than the Python code in terms of speed. This roughly confirms
claims made on the Go site. However, both Go versions are more verbose,
at least in my current programming style. So at least for small toy
programs, Go has not met that goal yet.

## Findings 2016

Seven years ago I didn't have to say `-O2` for the C version to beat the
Go version, now that's the only way. So we've come along quite nicely I
think? Also it used to be that the Go versions were not as close as they
are now to C, they were more in the middle between Python and C. Now
Python is the clear loser. Here's the data:

	./xtime ./freq </usr/share/dict/web2 >freq.out
	0.02u 0.00s 0.02r 5488kB ./freq

	./xtime ./simple_freq </usr/share/dict/web2 >simple_freq.out
	0.04u 0.00s 0.04r 9472kB ./simple_freq

	./xtime ./complex_freq </usr/share/dict/web2 >complex_freq.out
	0.28u 0.00s 0.28r 7888kB ./complex_freq

	./xtime python freq.py </usr/share/dict/web2 >python_freq.out
	1.12u 0.00s 1.12r 34096kB python freq.py

Sadly I don't have comparable results from 2010 anymore, would be nice
to have the exact numbers. (Fascinating aside: In all that time, my home
machine changed exactly zero.)
