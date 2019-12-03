build: packr2
	go build -o scaffold ./cmd/scaffold

packr2:
	if command -v packr2; then cd ./pkg/scaffold && packr2; fi

.PHONY: run
