build:
	rm -rf out > /dev/null
	mkdir out
	pwd
	. build.sh darwin linux windows
run:
	go run cmd/historder/*