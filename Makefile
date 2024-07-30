cert-priv:
	openssl genrsa -out cert/id_rsa 525
cert-pub:
	openssl rsa -in cert/id_rsa -pubout -out cert/id_rsa.pub

cert: cert-priv cert-pub