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
  # 安装 protoc-gen-gogo
  GO111MODULE=on GOPROXY=off go install k8s.io/code-generator/cmd/go-to-protobuf/protoc-gen-gogo

  proj::protoc::check_protoc

  local package=${1}
  proj::protoc::protoc "${package}"
  proj::protoc::format "${package}"
}

function proj::protoc::check_protoc() {
  if [[ -z "$(/usr/bin/which protoc)" || "$(protoc --version)" != "libprotoc ${PROTOC_VERSION}"* ]]; then
    echo "Generating protobuf requires protoc ${PROTOC_VERSION}."
    echo "Run scripts/install-protoc.sh or download and install the"
    echo "platform-appropriate Protobuf package for your OS from"
    echo "https://github.com/protocolbuffers/protobuf/releases"
    return 1
  fi
}


# Generates $1/api.pb.go from the protobuf file $1/api.proto
# $1: Full path to the directory where the api.proto file is
function proj::protoc::protoc() {
  local package=${1}
  gogopath=$(dirname "$(proj::util::find-binary "protoc-gen-gogo")")

  (
    cd "${package}"

    # This invocation of --gogo_out produces its output in the current
    # directory (despite gogo docs saying it would be source-relative, it
    # isn't).  The inputs to this function do not all have a common root, so
    # this works best for all inputs.
    PATH="${gogopath}:${PATH}" protoc \
      --proto_path="$(pwd -P)" \
      --proto_path="${PROJ_ROOT_DIR}/vendor" \
      --proto_path="${PROJ_ROOT_DIR}/third_party/protobuf" \
      --gogo_out=paths=source_relative,plugins=grpc:. \
      api.proto
  )
}

# Compares the contents of $1 and $2
# Echo's $3 in case of error and exits 1
function proj::protoc::diff() {
  local ret=0
  diff -I "gzipped FileDescriptorProto" -I "0x" -Naupr "${1}" "${2}" || ret=$?
  if [[ ${ret} -ne 0 ]]; then
    echo "${3}"
    exit 1
  fi
}

proj::protoc::loaded() {
  true
}


function proj::protoc::install() {
  # run in a subshell to isolate caller from directory changes
  (
    # 检查是否已经安装了正确版本的protoc
    if command -v protoc &>/dev/null && [[ "$(protoc --version)" == "libprotoc ${PROTOC_VERSION}"* ]]; then
      proj::log::info "protoc v${PROTOC_VERSION} is already installed."
      return 0
    fi

    local os
    local arch
    local download_folder
    local download_file

    os=$(proj::util::host_os)
    arch=$(proj::util::host_arch)
    download_folder="protoc-v${PROTOC_VERSION}-${os}-${arch}"
    download_file="${download_folder}.zip"

    mkdir -p "${HOME}/bin" || return 1
    cd "${HOME}/bin" || return 1
    
    local url
    if [[ ${os} == "darwin" ]]; then
      # TODO: switch to universal binary when updating to 3.20+
      url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-osx-x86_64.zip"
    elif [[ ${os} == "linux" && ${arch} == "amd64" ]]; then
      url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip"
    elif [[ ${os} == "linux" && ${arch} == "arm64" ]]; then
      url="https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-aarch_64.zip"
    else
      proj::log::info "This install script does not support ${os}/${arch}"
      return 1
    fi
    
    proj::util::download_file "${url}" "${download_file}"
    unzip -o "${download_file}" -d "${download_folder}"
    mv "${download_folder}/bin/protoc" $HOME/bin/protoc
    chmod +rX $HOME/bin/protoc
    rm -rvf "${download_file}" "${download_folder}"
    
    proj::log::info "$HOME/bin/protoc v${PROTOC_VERSION} installed."
  )
}

# Marker function to indicate protoc.sh has been fully sourced
proj::protoc::loaded() {
  return 0
}
