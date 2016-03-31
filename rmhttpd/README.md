# Go Tiny Web Servers Go

This all started when my friend Raluca posted a #twitcode web server in
Python:

https://twitter.com/ralucam/status/3658147537

I got curious what a simple, really small web server would look like in
Go, so I hacked one. Actually I hacked two, one that's close to Raluca's
and one that's a little more "full-featured" if I dare apply that term.
None of these are in any way "serious" of course.

## Python Versions

There's `original.py` which has (almost) exactly Raluca's original code;
this version is not runnable. The there's `edited.py` which is my
minimum-distance edit to make the web server runnable in a semi-convenient
way.

## Go Versions

There's `rmhttpd.go` which tries to stay close to Raluca's little hack,
at least in spirit. Properly formatted Go code is nowhere near as dense
as the Python hack, thankfully. Then there's `phfhttpd.go` which tries
to do some error checking, supports directory listings, and even speaks
something like a sad version of HTTP.
