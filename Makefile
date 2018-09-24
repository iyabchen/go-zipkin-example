
image: clean .deps 
	docker build \
		-t svca:1.0 \
		-f Dockerfile-svca \
		.

	docker build \
		-t svcb:1.0 \
		-f Dockerfile-svcb \
		.


Gopkg.lock: Gopkg.toml
	dep ensure -v
	@echo ==================================================
	@echo Gopkg.lock content
	@echo ==================================================
	@cat Gopkg.lock
	@echo ==================================================

.deps: Gopkg.lock
	dep ensure -v
	@touch .deps

clean:
	@rm -rf Gopkg.lock
	@rm -rf .deps
	@rm -rf vendor/


.PHONY: image clean
