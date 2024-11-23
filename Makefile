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
			env GOOS=$$GOOS GOARCH=$$GOARCH go build -o build/bin/$$GOOS/$$GOARCH/csi-driver cmd/csi-driver/main.go; \
		done; \
	done

.PHONY: package
package:
	docker build -f build/package/csi-driver/Dockerfile --platform linux/amd64 -t grantfl/csi-driver .

.PHONY: publish
publish:
	docker push grantfl/csi-driver

.PHONY: deploy
deploy:
	kubectl apply -f deployments/csi-driver

.PHONY: undeploy
undeploy:
	kubectl delete -f deployments/csi-driver
