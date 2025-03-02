#!/usr/bin/env bash


# shellcheck disable=SC2034 # Variables sourced in other scripts.

# Common utilities, variables and checks for all build scripts.
set -eEuo pipefail

# Unset CDPATH, having it set messes up with script import paths
unset CDPATH

USER_ID=$(id -u)
GROUP_ID=$(id -g)

# This will canonicalize the path
PROJ_ROOT_DIR=$(cd "$(dirname "${BASH_SOURCE[0]}")"/.. && pwd -P)

source "${PROJ_ROOT_DIR}/scripts/lib/init.sh"


# The variable PROJ_SERVER_SIDE_COMPONENTS is used to define onex server-side components.
# These components need to installed as a service.
declare -Ax PROJ_SERVER_SIDE_COMPONENTS=(
  ["api"]="apiserver"
)


# The variable PROJ_CLIENT_SIDE_COMPONENTS used to define onex client-side components.
# These components no need to installed as a service, but used as a command line.
declare -Ax PROJ_CLIENT_SIDE_COMPONENTS=(
  ["ctl"]="projctl"
)


# The variable ONEX_ALL_COMPONENTS is used to define all onex components.
# 12 useable components (@2024.01.01)
declare -Ax PROJ_ALL_COMPONENTS
for key in "${!PROJ_CLIENT_SIDE_COMPONENTS[@]}"; do
  PROJ_ALL_COMPONENTS["$key"]="${PROJ_CLIENT_SIDE_COMPONENTS[$key]}"
done
for key in "${!PROJ_SERVER_SIDE_COMPONENTS[@]}"; do
  PROJ_ALL_COMPONENTS["$key"]="${PROJ_SERVER_SIDE_COMPONENTS[$key]}"
done