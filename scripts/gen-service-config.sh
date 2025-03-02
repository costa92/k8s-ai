#!/usr/bin/env bash

# Copyright 2022 Lingfei Kong <colin404@foxmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/onexstack/onex.
#


PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${PROJ_ROOT_DIR}/scripts/common.sh"

if [ $# -ne 4 ];then
    proj::log::error "Usage: gen-service-config.sh SERVICE_NAME ENV_FILE TEMPLATE_FILE OUTPUT_DIR"
    exit 1
fi

export SERVICE_NAME=$1
ENV_FILE=$2
TEMPLATE_FILE=$3
OUTPUT_DIR=$4


if [ ! -d ${OUTPUT_DIR} ];then
  mkdir -p ${OUTPUT_DIR}
fi

echo "PROJ_ALL_COMPONENTS: ${PROJ_ALL_COMPONENTS}"
echo "PROJ_CLIENT_SIDE_COMPONENTS: ${PROJ_CLIENT_SIDE_COMPONENTS}"
echo "PROJ_SERVER_SIDE_COMPONENTS: ${PROJ_SERVER_SIDE_COMPONENTS}"


if [[ " ${PROJ_ALL_COMPONENTS[*]} " != *" ${SERVICE_NAME} "* ]]; then
  echo "SERVICE_NAME: ${SERVICE_NAME} not in PROJ_ALL_COMPONENTS"
  exit 0
fi

function get_apiserver_systemd_file()
{
  cat << EOF
# apiserver.service generated by scripts/gen-systemd.sh. DO NOT EDIT.
[Unit]
Description=Systemd unit file for apiserver
Documentation=https://github.com/onexstack/onex/blob/master/manifests/installation/onex/systemd/README.md

[Service]
WorkingDirectory=${PROJ_INSTALL_DIR}
ExecStartPre=/usr/bin/mkdir -p ${PROJ_DATA_DIR}/apiserver
ExecStartPre=/usr/bin/mkdir -p ${PROJ_LOG_DIR}
ExecStart=/opt/onex/bin/onex-apiserver --bind-address=${PROJ_APISERVER_BIND_ADDRESS} --secure-port ${PROJ_APISERVER_SECURE_PORT} --etcd-servers ${PROJ_APISERVER_ETCD_SERVERS} --client-ca-file=${PROJ_APISERVER_CLIENT_CA_FILE} --tls-cert-file=${PROJ_APISERVER_TLS_CERT_FILE} --tls-private-key-file=${PROJ_APISERVER_TLS_PRIVATE_KEY_FILE} --v=${PROJ_APISERVER_V_LEVEL}
Restart=always
RestartSec=5
StartLimitInterval=0

[Install]
WantedBy=multi-user.target
EOF
}

source ${ENV_FILE}

echo "TEMPLATE_FILE: ${TEMPLATE_FILE}"

# Some customized processing
case ${TEMPLATE_FILE} in
  *onex.systemd.tmpl.service)
    echo "TEMPLATE_FILE 111: ${TEMPLATE_FILE}"
    if [ "${SERVICE_NAME}" == "apiserver" ];then
      echo "SERVICE_NAME: ${SERVICE_NAME}"
      get_apiserver_systemd_file > ${OUTPUT_DIR}/${SERVICE_NAME}.service
      exit 0
    fi
    if [ "${SERVICE_NAME}" == "onexctl" ];then
      exit 0
    fi
    ;;
  *apiserver.config.tmpl.yaml)
    exit 0
    ;;
  *)
    ;;
esac

suffix=$(echo $TEMPLATE_FILE | awk -F'.' '{print $NF}')
${PROJ_ROOT_DIR}/scripts/gen-config.sh ${ENV_FILE} ${TEMPLATE_FILE} > ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix}

# Fix
if [[ "${TEMPLATE_FILE}" =~ .*onex.systemd.tmpl.service ]] && [[ "${SERVICE_NAME}" =~ onex.*controller ]];then
  escaped_onex_config_dir="$(sed -e 's/[\/&]/\\&/g' <<< "${ONEX_CONFIG_DIR}")"
  sed -i "/ExecStart=/s/$/ --kubeconfig=${escaped_onex_config_dir}\/config/" ${OUTPUT_DIR}/${SERVICE_NAME}.${suffix}
fi
