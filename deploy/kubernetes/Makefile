CONFIG_FILENAME ?= "minikube"

deploy:
	@make installMonitoring
	@echo "----------------"
	@make installMyApp

installMonitoring:
	@echo "Installing monitoring namespace"
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

installFluentd:
	@echo "Deploying fluentd"
	@helm install fluentd --values .fluentd/base-config.yml,/fluentd/${CONFIG_FILENAME}.yml --namespace monitoring --version 11.11.0 kokuwa/fluentd-elasticsearch

installElasticSearch:	
	@echo "Deploying elastic search"
	@helm install elasticsearch --values .elastic/base-config.yml,./elastic/${CONFIG_FILENAME}.yml  --namespace monitoring --version 7.12.0 elastic/elasticsearch 
	@echo "Elastic search deployed successfully"

upgradeElasticSearch:	
	@echo "Upgrading elastic search"
	@helm upgrade --values ./elastic/base-config.yml,./elastic/${CONFIG_FILENAME}.yml --namespace monitoring --version 7.12.0 elasticsearch elastic/elasticsearch 

installKibana:
	@echo "Deploying kibana"
	@helm install kibana --values ./kibana/base-config.yml,./kibana/${CONFIG_FILENAME}.yml --namespace monitoring --version 7.12.0 elastic/kibana
	@echo "Kibana deployed successfully"

installKubePrometheusStack:
	@echo "Deploying kube prometheus stack"
	@helm install prometheus --values ./prometheus/base-config.yml,./prometheus/${CONFIG_FILENAME}.yml --namespace monitoring --version 15.2.0 prometheus-community/kube-prometheus-stack 	
	@echo "kube prometheus stack deployed successfully"
	@echo "Deploying service/pod monitors"
	@kubectl create -f ./prometheus/monitors/

installNgnix:
	@echo "Deploying ngnix"
	@helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
	@helm helm repo update
	@helm install ngnix --values ./ngnix/base-config.yml,./ngnix/${CONFIG_FILENAME}.yml --version 0.46.0 ingress-nginx/ingress-nginx
