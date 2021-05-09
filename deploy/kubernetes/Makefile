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
	@helm install fluentd --values ./fluentd/base-config.yml,./fluentd/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 11.11.0 kokuwa/fluentd-elasticsearch
	@echo "Fluentd deployed successfully"

installElasticSearch:	
	@echo "Deploying elastic search"
	@helm install elasticsearch --values ./elastic/base-config.yml,./elastic/${CONFIG_FILENAME_PREFIX}-config.yml  --namespace monitoring --version 7.12.0 elastic/elasticsearch 
	@echo "Elastic search deployed successfully"

installKibana:
	@echo "Deploying kibana"
	@helm install kibana --values ./kibana/base-config.yml,./kibana/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 7.12.0 elastic/kibana
	@echo "Kibana deployed successfully"

installKubePrometheusStack:
	@echo "Deploying kube prometheus stack"
	@helm install prometheus --values ./prometheus/base-config.yml,./prometheus/${CONFIG_FILENAME_PREFIX}-config.yml --namespace monitoring --version 15.2.0 prometheus-community/kube-prometheus-stack 	
	@echo "Kube prometheus deployed successfully"

installNgnix:
	@echo "Deploying ngnix"
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	@helm repo update
	@helm install ngnix --values ./ngnix/base-config.yml,./ngnix/${CONFIG_FILENAME_PREFIX}-config.yml --version 3.30.0 ingress-nginx/ingress-nginx
	@echo "Ngnix deployed successfully"

installExternalDns:
	@echo "Installing need policy in order to the external dns to work"
	./external-dns/put-node-policy.sh
	@echo "Deploying external dns"
	kubectl create -f ./external-dns/deployment.yml
	@echo "External dns deployed successfully"

