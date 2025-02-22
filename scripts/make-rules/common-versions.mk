# ==============================================================================
# Versions used by all Makefiles
#

GO_TESTS_VERSION ?= v1.6.0
CODE_GENERATOR_VERSION ?= $(call get_go_version,k8s.io/code-generator)
GOLANGCI_LINT_VERSION := v1.64.5