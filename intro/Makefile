dev:
	docker build --target dev . -t go
build:
	docker build --target build . -t customer-app
prod:
# produce a container image without golang sdk 
# withou binary that we can go ahead and run somewhere
	docker build --target runtime . -t customer-app-prod
run:
	docker run -it --user=root -v ${PWD}:/work go sh
run_app:
	docker run customer-app-prod
.PHONY: dev build run prod run_app