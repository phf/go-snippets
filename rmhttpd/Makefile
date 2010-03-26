all: rmhttpd phfhttpd

rmhttpd: rmhttpd.8
	8l -o rmhttpd rmhttpd.8
rmhttpd.8: rmhttpd.go
	8g rmhttpd.go

phfhttpd: phfhttpd.8
	8l -o phfhttpd phfhttpd.8
phfhttpd.8: phfhttpd.go
	8g phfhttpd.go

clean:
	rm -rf *.8 rmhttpd phfhttpd
