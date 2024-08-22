args = $(foreach a,$($(subst -,_,$1)_args),$(if $(value $a),$a="$($a)"))

test:
	go test ./... -v

build-docker:
	docker build -t audio-container:v1 .

run-docker:
	docker run --rm -it -e PULSE_SERVER=unix:${XDG_RUNTIME_DIR}/pulse/native \
	-v ${XDG_RUNTIME_DIR}/pulse/native:${XDG_RUNTIME_DIR}/pulse/native \
	-v ~/.config/pulse/cookie:/root/.config/pulse/cookie \
	--device=/dev/snd:/dev/snd \
	audio-container:v1 --bpm=${bpm}

build-dev-docker:
	docker compose build

run-dev-docker:
	docker compose up