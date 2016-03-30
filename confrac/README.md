# Concurrent Fractals

Rendering fractals concurrently. Maybe even in parallel if we're lucky?

## Summary

There's a neat little program in *The Go Programming Language* that
renders fractals. I wanted to see what the state of goroutines is
these days, so I took the original sequential program and remixed
it into two concurrent versions.

The first concurrent version is *sane* because it starts one goroutine
for each *core* the machine has. The second concurrent version is
*insane* because it starts one goroutine for each *pixel* we render.
Imagine my surprise when I find that even the insane version finishes
in a reasonable amount of time; it just needs a *lot* of memory on most
platforms. The sane version even gives me a decent speedup (see below)!

## Running

There's a simple Makefile you can use:

	make bench

It'll build all three variants and run them against each other. On my
ancient AMD Athlon 64 X2 with two cores, I get this:

	./xtime ./single >single.png
	1.39u 0.01s 1.39r 46144kB ./single

	./xtime ./cores >cores.png
	1.17u 0.00s 0.68r 46256kB ./cores

	./xtime ./pixels >pixels.png
	3.30u 0.09s 1.81r 294480kB ./pixels

That's actually pretty impressive. I mean that's 1,048,576 goroutines
being launched! Try doing that with threads. Here's my slightly newer
laptop with four cores:

	./xtime ./single >single.png
	0.80u 0.00s 0.81r 7684kB ./single

	./xtime ./cores >cores.png
	0.99u 0.00s 0.53r 7856kB ./cores

	./xtime ./pixels >pixels.png
	2.42u 0.11s 0.86r 63344kB ./pixels

And here's some big-ish machine with 24 cores that our department has
lying around for undergraduates:

	./xtime ./single >single.png
	0.93u 0.07s 0.90r 12864kB ./single

	./xtime ./cores >cores.png
	1.33u 0.09s 0.32r 17072kB ./cores

	./xtime ./pixels >pixels.png
	7.74u 5.45s 1.21r 16388kB ./pixels

It's nice to see that the sane version actually scales with the number
of cores. But what fascinates me is the difference in memory usage. If
you have a good explanation, be sure to contact me!

## License

Copyright © 2016 Peter H. Froehlich.
https://creativecommons.org/licenses/by-nc-sa/4.0/

Based on example code from *The Go Programming Language*, page 61.
Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
