#!/bin/bash

set -euo pipefail

mkdir -p ca
cd ca

pw=agilePw

openssl req -new -text -passout pass:$pw -subj /CN=localhost -out server.req -keyout privkey.pem
openssl rsa -in privkey.pem -passin pass:$pw -out server.key
openssl req -x509 -in server.req -text -key server.key -out server.crt
chmod 600 server.key