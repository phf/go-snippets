import sys

def printable(c):
    if ord(c) < 32:
        return '-'
    else:
        return c

histo = {}
input = sys.stdin.read()

for c in input:
    histo[c] = histo.get(c, 0) + 1

for c in sorted(histo):
    print "%.2x  %c  %d" % (ord(c), printable(c), histo[c])
