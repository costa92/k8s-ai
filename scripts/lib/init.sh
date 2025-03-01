#!/usr/bin/env bash

# Copyright 2025 Qiuhong Long <costa9293@gmail.com>. All rights reserved.
# Use of this source code is governed by a MIT style
# license that can be found in the LICENSE file. The original repo for
# this file is https://github.com/costa92/k8s-ai.


set -o errexit
set +o nounset
set -o pipefail

# Short-circuit if init.sh has already been sourced
[[ $(type -t k8sai::init::loaded) == function ]] && return 0

# Unset CDPATH so that path interpolation can work correctly
# https://github.com/minerrnetes/minerrnetes/issues/52255
unset CDPATH



# Default use go modules
export GO111MODULE=on

# The root of the build/dist directory
PROJ_ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/../.." && pwd -P)"
SCRIPTS_DIR="${PROJ_ROOT_DIR}/scripts"

# 输出目录
K8SAI_OUTPUT_SUBPATH="${K8SAI_OUTPUT_SUBPATH:-_output}"
K8SAI_OUTPUT="${PROJ_ROOT_DIR}/${K8SAI_OUTPUT_SUBPATH}"


# source the common.sh file
# source "${SCRIPTS_DIR}/lib/common.sh"

source "${SCRIPTS_DIR}/lib/util.sh"
source "${SCRIPTS_DIR}/lib/logging.sh"

# 初始化本地环境
proj::init::local() {
 return 0
}
