dev: 
	sam build && sam local start-api

build: 
	sam build

deploy: 
	sam build && sam deploy --guided

deps:
   	go mod download

wire:
   	cd cmd && wire