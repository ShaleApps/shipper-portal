# Makefile

# Define variables
IMAGE_NAME = shaleapps/{{SERVICE_NAME}}
K8S_DEPLOY_NAME = {{SERVICE_NAME}}
DOCKERFILE = Dockerfile

# Default target
all: build

# Deploy
deploy: build push restart

# Target to build Docker image
build:
	docker build --build-arg ACCESS_TOKEN=$(GITHUB_TOKEN) --platform linux/amd64 --tag $(IMAGE_NAME) -f $(DOCKERFILE) .

# Push image to Docker Hub
push:
	docker push $(IMAGE_NAME):latest

restart:
	kubectl rollout restart deploy/$(K8S_DEPLOY_NAME)

k8s_apply:
	kubectl apply -n agnus -f k8s/$(ENV)/deployment.yaml
	kubectl apply -n agnus -f k8s/$(ENV)/service.yaml

# Target for dev environment
k8s_apply_dev:
	$(MAKE) k8s_apply ENV=dev

# Target for prod environment
k8s_apply_prod:

	$(MAKE) k8s_apply ENV=prod

test:
	go test -v ./...

# Target to clean up Docker image
clean:
	docker rmi $(IMAGE_NAME)

up-dev:	switch-context-port-forward
	docker-compose up -d

up-dev-envar: switch-context-port-forward
	docker-compose up -d && \
    DB_DATABASE="{{SERVICE_NAME_SNAKE_CASE}}" \
    DB_PORT="54325" \
    go run main.go

switch-context-port-forward:
	@kubectl config use-context gke_agnus-project-546094_us-central1_vrt-agnus-d-gke-k8s-usc1

.PHONY: down
down: kill docker-down

docker-down:
	@docker-compose down --remove-orphans

.PHONY: kill
kill:
	@lsof -t -i:5555 | xargs kill || true
	@lsof -t -i:8081 | xargs kill || true
	@lsof -t -i:8500 | xargs kill || true

up: switch-context-port-forward
	docker-compose up -d

rename-service:
	@echo "Please enter your service name:" ; \
	read originalServiceName ; \
	snakeCaseServiceName=`echo $$originalServiceName | tr '-' '_'` ; \
	camelCaseServiceName=`echo $$originalServiceName | awk 'BEGIN{FS="-|_";OFS=""}{for(i=2;i<=NF;i++)$$i=toupper(substr($$i,1,1)) substr($$i,2)} 1'` ; \
	if [ "$$(uname)" = "Darwin" ]; then \
	  grep -rl '{{SERVICE_NAME}}' ./ | xargs sed -i '' "s/{{SERVICE_NAME}}/$$originalServiceName/g" ; \
	  grep -rl '{{SERVICE_NAME_SNAKE_CASE}}' ./ | xargs sed -i '' "s/{{SERVICE_NAME_SNAKE_CASE}}/$$snakeCaseServiceName/g" ; \
	  grep -rl '{{SERVICE_NAME_CAMEL_CASE}}' ./ | xargs sed -i '' "s/{{SERVICE_NAME_CAMEL_CASE}}/$$camelCaseServiceName/g" ; \
	else \
	  grep -rl '{{SERVICE_NAME}}' ./ | xargs sed -i "s/{{SERVICE_NAME}}/$$originalServiceName/g" ; \
	  grep -rl '{{SERVICE_NAME_SNAKE_CASE}}' ./ | xargs sed -i "s/{{SERVICE_NAME_SNAKE_CASE}}/$$snakeCaseServiceName/g" ; \
	  grep -rl '{{SERVICE_NAME_CAMEL_CASE}}' ./ | xargs sed -i "s/{{SERVICE_NAME_CAMEL_CASE}}/$$camelCaseServiceName/g" ; \
	fi
	go fmt ./...
	go mod tidy

