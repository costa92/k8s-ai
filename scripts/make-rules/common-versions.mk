# ==============================================================================
# Versions used by all Makefiles
#

GO_TESTS_VERSION ?= v1.6.0
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
GOLANGCI_LINT_VERSION := v1.64.5
GIT_CHGLOG_VERSION ?= v0.15.4
GO_IMPORTS_VERSION ?= v0.30.0
LICENSE_VERSION ?= v5.0.4
GSEMVER_VERSION ?= v0.9.1
GO_GIT_LINT_VERSION ?= v1.1.1
WIRE_VERSION ?= $(call get_go_version,github.com/google/wire)
MOCKGEN_VERSION ?= $(call get_go_version,github.com/golang/mock)
