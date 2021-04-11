CONFIG_FILENAME_PREFIX ?= "aws"

deploy:
	@make installNgnix
	@echo "----------------"
	@make installMonitoring
	@echo "----------------"
	@make installMyApp
	@echo "----------------"
	@make installExternalDns

installMonitoring:
	@echo "Creating monitoring namespace"
	@kubectl create namespace monitoring 
	@echo "----------------"
	@make installElasticSearch
	@echo "----------------"
	@make installKibana
	@echo "----------------"
	@make installFluentd
	@echo "----------------"
	@make installKubePrometheusStack

installMyApp:
	@echo "Deploying my app"
	@kubectl create -f ./myapp
	@echo "My app deployed successfully"

installFluentd:
	@echo "Deploying fluentd"
	@helm install fluentd --values ./monitoring/fluentd/base-config.yml,./monitoring/fluentd/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 11.11.0 kokuwa/fluentd-elasticsearch
	@echo "Fluentd deployed successfully"

installElasticSearch:	
	@echo "Deploying elastic search"
	@helm install elasticsearch --values ./monitoring/elastic/base-config.yml,./monitoring/elastic/${CONFIG_FILENAME_PREFIX}-config.yml  --namespace monitoring --version 7.12.0 elastic/elasticsearch 
	@echo "Elastic search deployed successfully"

installKibana:
	@echo "Deploying kibana"
	@helm install kibana --values ./monitoring/kibana/base-config.yml,./monitoring/kibana/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 7.12.0 elastic/kibana
	@echo "Kibana deployed successfully"

installKubePrometheusStack:
	@echo "Deploying kube prometheus stack"
	@helm install prometheus --values ./monitoring/prometheus/base-config.yml,./monitoring/prometheus/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 15.2.0 prometheus-community/kube-prometheus-stack 	
	@echo "Kube prometheus deployed successfully"

installNgnix:
	@echo "Deploying ngnix"
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	@helm repo update
	@helm install ngnix --values ./routing/ngnix/base-config.yml,./routing/ngnix/${CONFIG_FILENAME_PREFIX}-config.yml --version 3.30.0 ingress-nginx/ingress-nginx
	@echo "Ngnix deployed successfully"

installExternalDns:
	@echo "Installing need policy in order to the external dns to work"
	./external-dns/put-node-policy.sh
	@echo "Deploying external dns"
	kubectl create -f ./routing/external-dns/deployment.yml
	@echo "External dns deployed successfully"

