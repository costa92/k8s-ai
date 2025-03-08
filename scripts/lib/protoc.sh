#!/usr/bin/env bash

set -o errexit
set -o nounset
set -o pipefail



[[ $(type -t proj::protoc::loaded) == "function" ]] && return 0

# The root of the build/dist directory
PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/../..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"


#PROTOC_VERSION=23.4
PROTOC_VERSION=28.3

function proj::protoc::generate_proto(){
  proj::golang::setup_env
}


proj::protoc::loaded() {
  true
}


