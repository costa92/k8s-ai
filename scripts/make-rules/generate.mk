# ==============================================================================
# Makefile helper functions for generate necessary files
#

SERVICES ?= $(filter-out tools,$(foreach service,$(wildcard ${PROJ_ROOT_DIR}/cmd/*),$(notdir ${service})))


.PHONY: gen.ca.%
gen.ca.%: ## Generate CA files.
	$(eval CA := $(word 1,$(subst ., ,$*)))
	@echo "===========> Generating CA files for $(CA)"
	@${SCRIPTS_DIR}/gen-certs.sh generate-node-cert $(OUTPUT_DIR)/cert $(CA)

.PHONY: gen.ca
gen.ca: $(addprefix gen.ca., $(CERTIFICATES)) ## Generate all CA files.


.PHONY: gen.kubeconfig
gen.kubeconfig: gen.ca ## Generate kubeconfig files.
	@echo "===========> Generating admin kubeconfig file"
	@$(SCRIPTS_DIR)/gen-kubeconfig.sh $(OUTPUT_DIR)/cert/ca.pem \
		$(OUTPUT_DIR)/cert/admin.pem $(OUTPUT_DIR)/cert/admin-key.pem > \
		$(OUTPUT_DIR)/config