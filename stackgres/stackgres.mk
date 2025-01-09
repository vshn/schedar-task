

.PHONY: stackgres
stackgres: export KUBECONFIG = $(KIND_KUBECONFIG)
stackgres: ## Install StackGres
	helm repo add stackgres-charts https://stackgres.io/downloads/stackgres-k8s/stackgres/helm/ --force-update
	helm upgrade --install --create-namespace --namespace stackgres stackgres-operator  stackgres-charts/stackgres-operator --wait
	kubectl -n stackgres wait --for condition=Available deployment/stackgres-operator --timeout 120s

	# wait max 60 seconds for secret to be created - it takes little bit longer now for secret to appear, therefore we need a mechanism to block execution until it appears
	echo "waiting for stackgres-restapi-admin secret creation..."
	@for i in $$(seq 1 60); do \
        if kubectl get secret stackgres-restapi-admin -n stackgres > /dev/null 2>&1; then \
            break; \
        else \
            sleep 1; \
        fi; \
	done;

	# Set simple credentials for development
	NEW_USER=admin &&\
	NEW_PASSWORD=password &&\
	patch=$$(kubectl create secret generic -n stackgres stackgres-restapi-admin  --dry-run=client -o json \
		--from-literal=k8sUsername="$$NEW_USER" \
		--from-literal=password="$$(echo -n "$${NEW_USER}$${NEW_PASSWORD}"| sha256sum | awk '{ print $$1 }' )") &&\
	kubectl patch secret -n stackgres stackgres-restapi-admin -p "$$patch" &&\
	kubectl patch secrets --namespace stackgres stackgres-restapi-admin --type json -p '[{"op":"remove","path":"/data/clearPassword"}]' | true &&\
	encoded=$$(echo -n "$$NEW_PASSWORD" | base64) && \
	kubectl patch secrets --namespace stackgres stackgres-restapi-admin --type json -p "[{\"op\":\"add\",\"path\":\"/data/clearPassword\", \"value\":\"$${encoded}\"}]" | true


