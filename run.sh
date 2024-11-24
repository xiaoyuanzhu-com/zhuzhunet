#!/bin/bash

function run_api() {
	go run .
}

function run_ui() {
	cd ui
	npm run dev
}

function build_docker() {
	docker build -t zhuzhunet .
    # docker run --rm --name zhuzhunet -p 8080:8080 -v $(pwd)/configs:/config zhuzhunet
}

if [ "$1" = "ui" ]; then
	run_ui
elif [ "$1" = "api" ]; then
	run_api
elif [ "$1" = "docker" ]; then
	build_docker
else
	echo "Usage: $0 [ui|api|docker]"
	exit 1
fi
