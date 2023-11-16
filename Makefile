cert:
	openssl genrsa -out cert/id_rsa 525
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub