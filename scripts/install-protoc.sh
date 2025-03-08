#!/usr/bin/env bash



# This script is convenience to download and install protoc in third_party.
# Usage: `hack/install-protoc.sh`.

set -o errexit
set -o nounset
set -o pipefail

PROJ_ROOT_DIR=$(dirname "${BASH_SOURCE[0]}")/..
source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"
source "${PROJ_ROOT_DIR}/scripts/lib/protoc.sh"

proj::protoc::install
