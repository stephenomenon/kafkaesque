APP := kafkaesque

build:
	docker build -t $(APP) .

run:
	docker run -it --rm --name so-$(APP) $(APP)
