ca.key.pem

```
openssl genrsa 2048 > ca.key.pem
```

ca.key.pem

```
openssl req -new -sha256 -x509 -nodes -days 3600 -key ca.key.pem -out ca.cert.pem -subj "/C=TW/ST=Taiwan/L=Taipei/O=cclin/OU=cclin/CN=ca.cclin/emailAddress=cclin81922@gmail.com"
```

server.key.pem, server.csr

```
openssl req -new -sha256 -keyout server.key.pem -out server.csr -days 365 -newkey rsa:2048 -nodes -subj "/C=TW/ST=Taiwan/L=Taipei/O=cclin/OU=cclin/CN=localhost.localdomain/emailAddress=cclin81922@gmail.com"
```

server.cert.pem, ca.srl

```
openssl x509 -req -days 365 -sha1 -CA ca.cert.pem -CAkey ca.key.pem -CAcreateserial -in server.csr -out server.cert.pem
```

client.key.pem, client.csr

```
openssl req -new -sha256 -keyout client.key.pem -out client.csr -days 365 -newkey rsa:2048 -nodes -subj "/C=TW/ST=Taiwan/L=Taipei/O=cclin/OU=cclin/CN=client/emailAddress=cclin81922@gmail.com"
```

client.cert.pem, ca.srl

```
openssl x509 -req -days 365 -sha1 -CA ca.cert.pem -CAkey ca.key.pem -CAcreateserial -in client.csr -out client.cert.pem
```
