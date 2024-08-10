test:
	go test ./... -v

build-docker:
	docker build -t audio-container:v1 .

run-docker:
	docker run --rm -it --device=/dev/snd:/dev/snd audio-container:v1