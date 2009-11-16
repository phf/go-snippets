all: rmhttpd phfhttpd

rmhttpd: rmhttpd.6
	6l -o rmhttpd rmhttpd.6
rmhttpd.6: rmhttpd.go
	6g rmhttpd.go

phfhttpd: phfhttpd.6
	6l -o phfhttpd phfhttpd.6
phfhttpd.6: phfhttpd.go
	6g phfhttpd.go

clean:
	rm -rf *.6 rmhttpd phfhttpd
