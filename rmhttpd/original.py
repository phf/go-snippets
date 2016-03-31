# except for this line, this is Raluca's original web server
s=socket.socket();s.bind(('',80));s.listen(9)
while(1): c,a=s.accept();c.send(open(c.recv(99)
[5:].split()[0]).read());c.close() #twitcode
