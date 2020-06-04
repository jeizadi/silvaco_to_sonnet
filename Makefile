build_all: build_mac build_win build_linux
	@echo "Done building all"

build_mac:
	@go build -o bin/mac/stos_mac *.go
	@echo "Done building mac"

build_win:
	@GOOS=windows go build -o bin/windows/stos_win.exe *.go
	@echo "Done building win"

build_linux:
	@GOOS=linux go build -o bin/linux/stos_linux *.go
	@echo "Done building linux"