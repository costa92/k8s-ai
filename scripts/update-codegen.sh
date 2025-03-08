#!/usr/bin/env bash

# shellcheck disable=2046 # printf word-splitting is intentional

set -o errexit
set -o nounset
set -o pipefail

# 获取kube版本
KUBE_VERSION="${KUBE_VERSION:-1}"

# 项目文件
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..

source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/protoc.sh"
