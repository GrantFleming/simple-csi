GOOSES := linux darwin
GOARCHES := arm64 amd64

clean:
	rm -rf build/bin

test:
	go test ./...

.PHONY: build
build:
	for GOOS in $(GOOSES); do \
		for GOARCH in $(GOARCHES); do \
			env GOOS=$$GOOS GOARCH=$$GOARCH go build -o build/bin/$$GOOS/$$GOARCH/simple-csi cmd/simple-csi/main.go; \
		done; \
	done

.PHONY: package
package:
	docker build -f build/package/simple-csi/Dockerfile --platform linux/amd64 -t grantfl/simple-csi .

.PHONY: publish
publish:
	docker push grantfl/simple-csi

.PHONY: deploy
deploy:
	kubectl apply -f deployments/simple-csi

.PHONY: undeploy
undeploy:
	kubectl delete -f deployments/simple-csi
