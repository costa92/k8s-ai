# ==============================================================================
#  Makefile helper functions for tools
#
# Rules name starting with `_` mean that it is not recommended to call directly through make command, 
# like `make _install.gotests`, you should run `make tools.install.gotests` instead.
#

# Specify tools category.
CODE_GENERATOR_TOOLS := client-gen conversion-gen deepcopy-gen defaulter-gen informer-gen lister-gen prerelease-lifecycle-gen \
					   register-gen applyconfiguration-gen go-to-protobuf

# code-generator is a makefile target not a real tool.
CI_WORKFLOW_TOOLS := code-generator golangci-lint goimports wire

# Tools that need to be installed manually
MANUAL_INSTALL_TOOLS := kafka

# Other tools used in this project
OTHER_TOOLS := mockgen uplift git-chglog addlicense kratos kind go-apidiff gotests \
			   cfssl go-gitlint kustomize kafkactl kube-linter kubeconform kubectl \
			   helm-docs db2struct air swagger license gothanks kubebuilder \
			   go-junit-report controller-gen protoc-go-inject-tag

# Main installation targets
.PHONY: tools.install
tools.install: install.ci _install.other tools.print-manual-tool ## Install all tools.

.PHONY: tools.print-manual-tool
tools.print-manual-tool:
	@echo "===========> The following tools may need to be installed manually:"
	@echo $(MANUAL_INSTALL_TOOLS) | awk 'BEGIN{RS=" "} {printf("%15s%s\n","- ",$$0)}'

.PHONY: tools.install.%
tools.install.%: ## Install a specified tool.
	@echo "===========> Installing $*"
	@$(MAKE) _install.$*

.PHONY: tools.verify.%
tools.verify.%: ## Verify a specified tool.
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

# CI tools installation
.PHONY: _install.ci
_install.ci: $(addprefix tools.install., $(CI_WORKFLOW_TOOLS)) ## Install necessary tools used by CI/CD workflow.

.PHONY: _install.other
_install.other: $(addprefix tools.install., $(OTHER_TOOLS)) ## Install other tools.

# Code generator tools verification and installation
.PHONY: tools.verify.code-generator
tools.verify.code-generator: $(addprefix _verify.code-generator., $(CODE_GENERATOR_TOOLS)) ## Verify code generator tools.

.PHONY: _verify.code-generator.%
_verify.code-generator.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.code-generator.$*; fi

.PHONY: _install.code-generator
_install.code-generator: $(addprefix _install.code-generator., $(CODE_GENERATOR_TOOLS)) ## Install all code generator tools.

.PHONY: _install.code-generator.%
_install.code-generator.%: ## Install specified code generator tool.
	@if [ -z "$(CODE_GENERATOR_VERSION)" ]; then \
		echo "Warning: CODE_GENERATOR_VERSION is not set, using default version."; \
	else \
		echo "===========> Installing code-generator tool: $*" \
		-$(GO) install k8s.io/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION) || true; \
	fi

# Individual tool installation targets
.PHONY: _install.golangci-lint
_install.golangci-lint: ## Install golangci-lint with bash completion.
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh golangci-lint bash

.PHONY: _install.golangci-lint.zsh
_install.golangci-lint.zsh: ## Install golangci-lint with zsh completion.
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh golangci-lint zsh

.PHONY: _install.gotests
_install.gotests: ## Install gotests.
	@$(GO) install github.com/cweill/gotests/gotests@$(GO_TESTS_VERSION)

.PHONY: _install.git-chglog
_install.git-chglog: ## Install git-chglog for CHANGELOG generation.
	@$(GO) install github.com/git-chglog/git-chglog/cmd/git-chglog@$(GIT_CHGLOG_VERSION)

.PHONY: _install.goimports
_install.goimports: ## Install goimports.
	@$(GO) install golang.org/x/tools/cmd/goimports@$(GO_IMPORTS_VERSION)

.PHONY: _install.wire
_install.wire: ## Install wire dependency injection tool.
	@if [ -z "$(WIRE_VERSION)" ]; then \
		echo "Warning: WIRE_VERSION is not set, using default version."; \
	else \
		echo "===========> Installing wire" \
		$(GO) install github.com/google/wire/cmd/wire@$(WIRE_VERSION); \
	fi

.PHONY: _install.mockgen
_install.mockgen: ## Install mockgen mock generator.
	@if [ -z "$(MOCKGEN_VERSION)" ]; then \
		echo "Warning: MOCKGEN_VERSION is not set, using default version."; \
	else \
		echo "===========> Installing mockgen" \
		$(GO) install github.com/golang/mock/mockgen@$(MOCKGEN_VERSION); \
	fi

.PHONY: _install.uplift
_install.uplift: ## Install uplift with bash completion.
	@export UPLIFT_INSTALL_DIR=$(HOME)/bin && \
		curl --retry 10 -sL https://raw.githubusercontent.com/gembaadvantage/uplift/main/scripts/install | bash -s -- '--no-sudo'
	@$(SCRIPTS_DIR)/add-completion.sh uplift bash

.PHONY: _install.license
_install.license: ## Install license generator tool.
	@$(GO) install github.com/nishanths/license/v5@$(LICENSE_VERSION)

.PHONY: _install.gsemver
_install.gsemver: ## Install gsemver with bash completion.
	@$(GO) install github.com/arnaud-deprez/gsemver@$(GSEMVER_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh gsemver bash

.PHONY: _install.gsemver.zsh
_install.gsemver.zsh: ## Install gsemver with zsh completion.
	@$(GO) install github.com/arnaud-deprez/gsemver@$(GSEMVER_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh gsemver zsh

.PHONY: _install.go-gitlint
_install.go-gitlint: ## Install go-gitlint commit message linter.
	@$(GO) install github.com/marmotedu/go-gitlint/cmd/go-gitlint@$(GO_GIT_LINT_VERSION)