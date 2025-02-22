# ==============================================================================
#  Makefile helper functions for tools
#
# Rules name starting with `_` mean that it is not recommended to call directly through make command, 
# like `make _install.gotests`, you should run `make tools.install.gotests` instead.
#

# Specify tools category.
CODE_GENERATOR_TOOLS= client-gen conversion-gen deepcopy-gen defaulter-gen informer-gen lister-gen prerelease-lifecycle-gen \
					  register-gen applyconfiguration-gen go-to-protobuf

# code-generator is a makefile target not a real tool.
CI_WORKFLOW_TOOLS := code-generator golangci-lint goimports wire
# unused tools in this project: gentool
OTHER_TOOLS := mockgen uplift git-chglog addlicense kratos kind go-apidiff gotests \
			   cfssl go-gitlint kustomize kafkactl kube-linter kubeconform kubectl \
			   helm-docs db2struct gentool air swagger license gothanks kubebuilder \
			   go-junit-report controller-gen protoc-go-inject-tag

MANUAL_INSTALL_TOOLS := kafka

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

.PHONY: _install.ci
_install.ci: $(addprefix tools.install., $(CI_WORKFLOW_TOOLS)) ## Install necessary tools used by CI/CD workflow.

.PHONY: _install.other
_install.other: $(addprefix tools.install., $(OTHER_TOOLS)) ## Install other tools.

.PHONY: tools.verify.code-generator
tools.verify.code-generator: $(addprefix _verify.code-generator., $(CODE_GENERATOR_TOOLS)) ## Verify a specified tool.
.PHONY: _verify.code-generator.%
_verify.code-generator.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.code-generator.$*; fi


.PHONY: _install.code-generator
_install.code-generator: $(addprefix _install.code-generator., $(CODE_GENERATOR_TOOLS)) ## Install all necessary code-generator tools.

.PHONY: _install.code-generator.%
_install.code-generator.%: ## Install specified code-generator tool.
	@if [ -z "$(CODE_GENERATOR_VERSION)" ]; then \
		echo "Warning: CODE_GENERATOR_VERSION is not set, using default version."; \
	else \
		echo "===========> Installing code-generator tool: $*" \
		-$(GO) install k8s.io/code-generator/cmd/$*@$(CODE_GENERATOR_VERSION) || true; \
	fi;


.PHONY: _install.golangci-lint
_install.golangci-lint: ## Install golangci-lint.
	@$(GO) install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)
	@$(SCRIPTS_DIR)/add-completion.sh golangci-lint bash


.PHONY: _install.gotests
_install.gotests: ## Install gotests.
	@$(GO) install github.com/cweill/gotests/gotests@$(GO_TESTS_VERSION)
