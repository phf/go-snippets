# minimal changes to original.py to make it run conveniently
import socket;s=socket.socket();s.bind(('',8080));s.listen(9)
while 1:c,a=s.accept();c.send(open("."+c.recv(1024).split()
[1]).read());c.close() #twitcode
