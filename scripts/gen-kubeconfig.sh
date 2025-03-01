#!/usr/bin/env bash

# This script is used to generate a kubeconfig file for the apiserver.
# It is used to connect to the apiserver from the client.
# The kubeconfig file is generated from the CA certificate, the client certificate and the client key.

caFile="$1"
certFile="$2"
keyFile="$3"

APISERVER_HOST=${APISERVER_HOST:-127.0.0.1}
APISERVER_SECURE_PORT=${APISERVER_SECURE_PORT:-52443}



function gen_kubeconfig() {
  host_os=$(uname -s)
  if [[ "${host_os}" == "Darwin" ]]; then
    cn=`openssl x509 -in $2 -noout -subject |  awk -F'CN=' '{print $2}' | awk -F',' '{print $1}'`
  elif [[ "${host_os}" == "Linux" ]]; then
    cn=`openssl x509 -in $2 -noout -text|awk -F'CN = ' '/Subject.*CN/{print $NF}'`
  else
    echo "Unsupported host OS.  Must be Linux or Mac OS X."
    exit 1
  fi

  # 生成kubeconfig文件
	cat << EOF
apiVersion: v1
clusters:
- cluster:
    server: https://${APISERVER_HOST}:${APISERVER_SECURE_PORT}
    certificate-authority-data: `base64 -i $1 |tr -d '\n'`
  name: ${cn}
contexts:
- context:
    cluster: ${cn}
    user: ${cn}
  name: default
current-context: default
kind: Config
preferences: {}
users:
- name: ${cn}
  user:
    client-certificate-data: `base64 -i $2 |tr -d '\n'`
    client-key-data: `base64 -i $3 |tr -d '\n'`
EOF
}

gen_kubeconfig ${caFile} ${certFile} ${keyFile}
