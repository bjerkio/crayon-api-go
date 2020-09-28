all: install_deps generate

.PHONY: install_deps
install_deps:
	docker pull quay.io/goswagger/swagger
	go generate ./tools.go

.PHONY: generate
generate:
	alias swagger="docker run --rm -it -e GOPATH=$$HOME/go:/go -v $$HOME:$$HOME -w $$(pwd) quay.io/goswagger/swagger"
	swagger generate client -f \
	./swagger.json -t ./ --skip-validation
	./bin/addlicense -c "Bjerk AS" **/*.go	