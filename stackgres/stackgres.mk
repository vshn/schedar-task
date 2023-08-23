

.PHONY: stackgres
stackgres: export KUBECONFIG = $(KIND_KUBECONFIG)
stackgres: ## Install StackGres
	helm repo add stackgres-charts https://stackgres.io/downloads/stackgres-k8s/stackgres/helm/
	helm upgrade --install --create-namespace --namespace stackgres stackgres-operator  stackgres-charts/stackgres-operator
	# Set simple credentials for development
	NEW_USER=admin &&\
	NEW_PASSWORD=admin &&\
	patch=$$(kubectl create secret generic -n stackgres stackgres-restapi  --dry-run=client -o json \
		--from-literal=k8sUsername="$$NEW_USER" \
		--from-literal=password="$$(echo -n "$${NEW_USER}$${NEW_PASSWORD}"| sha256sum | awk '{ print $$1 }' )") &&\
	kubectl patch secret -n stackgres stackgres-restapi -p "$$patch" &&\
	kubectl patch secrets --namespace stackgres stackgres-restapi --type json -p '[{"op":"remove","path":"/data/clearPassword"}]' | true &&\
	encoded=$$(echo -n "$$NEW_PASSWORD" | base64) && \
	kubectl patch secrets --namespace stackgres stackgres-restapi --type json -p "[{\"op\":\"add\",\"path\":\"/data/clearPassword\", \"value\":\"$${encoded}\"}]" | true

