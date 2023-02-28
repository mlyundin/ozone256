
build-all:
	cd checkout && GOOS=linux make build
	cd loms && GOOS=linux make build
	cd notifications && GOOS=linux make build

run-all: build-all
	sudo docker compose up --force-recreate --build

build-all-w:
	wsl --shell-type login --cd /mnt/d/repo/ozon/Homework make build-all

run-all-w: build-all-w
	docker compose up --force-recreate --build

precommit:
	cd checkout && make precommit
	cd loms && make precommit
	cd notifications && make precommit
