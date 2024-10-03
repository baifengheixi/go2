.PHONY: compose skaffold test

compose:
	@echo "Building docker images"
	cd ../deployments && docker-compose up --build

skaffold:
	skaffold dev -f ../deployments/skaffold.yaml

test:
	go test -v ../test/...