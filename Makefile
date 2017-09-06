APP:=kafkaesque

image:
	docker build -t $(APP) .

run:
	docker run -it --rm \
	-v $(PWD)/config.toml:/go/src/$(APP)/config.toml \
	$(APP)
