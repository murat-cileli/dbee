BUILD_PARAMS = go build -C src -ldflags '-s -w' -trimpath -o ../bin

build:	clean
	GOARCH=amd64 GOOS=linux ${BUILD_PARAMS}/linux-amd64/dbee
	GOARCH=386 GOOS=linux ${BUILD_PARAMS}/linux-i386/dbee

	GOARCH=amd64 GOOS=freebsd ${BUILD_PARAMS}/freebsd-amd64/dbee
	GOARCH=386 GOOS=freebsd ${BUILD_PARAMS}/freebsd-i386/dbee

	GOARCH=amd64 GOOS=darwin ${BUILD_PARAMS}/darwin-amd64/dbee
	GOARCH=arm64 GOOS=darwin ${BUILD_PARAMS}/darwin-arm64/dbee

	GOARCH=amd64 GOOS=windows ${BUILD_PARAMS}/windows-amd64/dbee.exe
	GOARCH=386 GOOS=windows ${BUILD_PARAMS}/windows-i386/dbee.exe
	
clean:
	go clean
	rm -rf ./bin/*
