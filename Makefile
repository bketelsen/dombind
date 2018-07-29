build:
	GOROOT=~/gowasm GOARCH=wasm GOOS=js ~/gowasm/bin/go build -o test.wasm dombind.go
clean:
	rm test.wasm
