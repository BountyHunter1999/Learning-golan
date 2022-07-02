dev:
	docker build --target dev . -t go
build:
	docker build --target build . -t videos
prod:
# produce a container image without golang sdk 
# withou binary that we can go ahead and run somewhere
	docker build --target runtime . -t videos
run:
	docker run -it --rm -p 8085:8000 -v ${PWD}:/work go sh
run_app:
	docker run -it -p 8085:8000 videos
change_owner:
	sudo chmod -R 777 videos 
.PHONY: dev build run prod run_app change_owner