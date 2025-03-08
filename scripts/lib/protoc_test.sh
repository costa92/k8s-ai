#!/usr/bin/env bash

#! /usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail

# The root of the build/dist directory
PROJ_ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"

source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/logging.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/golang.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/version.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/protoc.sh"
# 生成文件存放目录
export LOCAL_OUTPUT_ROOT=${PROJ_ROOT_DIR}/_output

proj::protoc::generate_proto
