# Build all by default, even if it's not first
.DEFAULT_GOAL := help

.PHONY: all


# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/all.mk


# ==============================================================================
# Usage

define USAGE_OPTIONS

endef

.PHONY: gen-k8s
gen-k8s: ## Generate all necessary kubernetes related files, such as deepcopy files
	@$(PROJ_ROOT_DIR)/scripts/update-codegen.sh
	# The following command is old generate way with makefile script.
	# Comment here as a code history.
	# $(MAKE) -s generated.files



.PHONY: protoc
protoc: ## Generate api proto files.
	$(MAKE) gen.protoc


## --------------------------------------
## Binaries
## --------------------------------------

##@ Build

.PHONY: build
build: tidy ## Build source code for host platform.
	$(MAKE) go.build


.PHONY: add-copyright
add-copyright: ## Ensures source code files have copyright license headers.
	$(MAKE) copyright.add


.PHONY: copy-githooks
copy-githooks:
	@$(COPY_GITHOOK)
	@chmod +x .git/hooks/*
	@echo "Copy githooks done"


.PHONY: ca
ca: ## Generate CA files for all onex components.
	$(MAKE) gen.ca


## --------------------------------------
## Tidy
## --------------------------------------

.PHONY: tidy
tidy:
	@$(GO) mod tidy

.PHONY: install-tools
install-tools: ## Install CI-related tools. Install all tools by specifying `A=1`.
	$(MAKE) install.ci
	if [[ "$(A)" == 1 ]]; then                                             \
		$(MAKE) _install.other ;                                            \
	fi

.PHONY: release
release: ## Publish a release on the release branch.
	$(MAKE) release.run

	
.PHONY: targets
targets: Makefile ## Show all Sub-makefile targets.
	@for mk in `echo $(MAKEFILE_LIST) | sed 's/Makefile //g'`; do echo -e \\n\\033[35m$$mk\\033[0m; awk -F':.*##' '/^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 }' $$mk;done;



# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk command is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php
.PHONY: help
help: Makefile ## Display this help info.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<TARGETS> <OPTIONS>\033[0m\n\n\033[35mTargets:\033[0m\n"} /^[0-9A-Za-z._-]+:.*?##/ { printf "  \033[36m%-45s\033[0m %s\n", $$1, $$2 } /^\$$\([0-9A-Za-z_-]+\):.*?##/ { gsub("_","-", $$1); printf "  \033[36m%-45s\033[0m %s\n", tolower(substr($$1, 3, length($$1)-7)), $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' Makefile #$(MAKEFILE_LIST)
	@echo -e "$$USAGE_OPTIONS"