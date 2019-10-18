build:
	rm -rf out > /dev/null
	mkdir out
	GOOS=darwin GOARCH=amd64 go build -o out/historder_darwin cmd/historder/*
	GOOS=linux GOARCH=amd64 go build -o out/historder_linux cmd/historder/*
	GOOS=windows GOARCH=amd64 go build -o out/historder_windows cmd/historder/*
run:
	go run cmd/historder/*