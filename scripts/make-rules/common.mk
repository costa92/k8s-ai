
# ==============================================================================
# Includes
# include the common make file

ifeq ($(origin PROJ_ROOT_DIR),undefined)
PROJ_ROOT_DIR :=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))
endif


include $(PROJ_ROOT_DIR)/scripts/make-rules/common-versions.mk



# 设置默认 shell 为 bash,因为某些环境中 sh 不是链接到 bash
# -o errexit: 命令返回非0退出码时立即退出
# -o pipefail: 管道中任一命令失败则整个管道失败
# +o nounset: 允许使用未定义的变量
SHELL := /usr/bin/env bash -o errexit -o pipefail +o nounset

# 设置 shell 标志
# -e: 命令失败时退出
# -c: 从字符串中读取命令
.SHELLFLAGS = -ec

# 为 bash shell 设置错误退出标志
export SHELLOPTS := errexit



# Linux 命令设置
# FIND: 查找文件,排除 third_party 和 vendor 目录
FIND := find . ! -path './third_party/*' ! -path './vendor/*'
# XARGS: 从标准输入构建和执行命令行,当输入为空时不运行命令
XARGS := xargs --no-run-if-empty

# Helper function to get dependency version from go.mod
get_go_version = $(shell if ! go list -m $1 >/dev/null 2>&1; then echo "Error: Failed to get version for $1" >&2; exit 1; fi; go list -m $1 | awk '{print $$2}')

# Helper function to install a dependency	
define go_install
$(info ===========> Installing $(1)@$(2))
$(GO) install $(1)@$(2)
endef


# Copy githook scripts when execute makefile
COPY_GITHOOK:=$(shell cp -f githooks/* .git/hooks/)

# Specify components which need certificate
ifeq ($(origin CERTIFICATES),undefined)
CERTIFICATES=apiserver admin
endif

# 设置构建输出目录
BUILD_DIR := $(PROJ_ROOT_DIR)/build


# ==============================================================================
# Build options
#
PRJ_SRC_PATH :=github.com/costa92/k8s-ai

COMMA := ,
SPACE :=
SPACE +=


ifeq ($(origin OUTPUT_DIR),undefined)
OUTPUT_DIR := $(PROJ_ROOT_DIR)/_output
$(shell mkdir -p $(OUTPUT_DIR))
endif


MANIFESTS_DIR=$(PROJ_ROOT_DIR)/manifests
SCRIPTS_DIR=$(PROJ_ROOT_DIR)/scripts



APIROOT ?= $(PROJ_ROOT_DIR)/pkg/api
APISROOT ?= $(PROJ_ROOT_DIR)/pkg/apis