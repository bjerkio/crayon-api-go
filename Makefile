all: install_deps download_swagger generate

.PHONY: install_deps
install_deps:
	# go get -u github.com/google/addlicense
	docker pull quay.io/goswagger/swagger
	go generate ./tools.go

.PHONY: download_swagger
download_swagger:
	curl -o ./swagger.json https://api.crayon.com/swagger/v1/swagger.json

.PHONY: generate
generate:
	alias swagger="docker run --rm -it -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger"
	swagger generate client -f \
	./swagger.json -t ./ --skip-validation
	./bin/addlicense -c "Bjerk AS" **/*.go	