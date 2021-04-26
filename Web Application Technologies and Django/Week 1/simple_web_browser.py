import socket

sckt0=socket.socket(socket.AF_INET, socket.SOCK_STREAM)
sckt0.connect(('data.pr4e.org', 80))
cmd='GET http://data.pr4e.org/page1.htm HTTP/1.0\r\n\r\n'.encode()
sckt0.send(cmd)

while True:
	data=sckt0.recv(512)
	if len(data)<1:
		break
	print(data.decode(),end='')

sckt0.close()