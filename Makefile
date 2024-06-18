build:
	GOARCH=amd64 GOOS=linux go build -C src -o ../bin/linux-amd64/dbee
	GOARCH=amd64 GOOS=freebsd go build -C src -o ../bin/freebsd-amd64/dbee
	GOARCH=amd64 GOOS=darwin go build -C src -o ../bin/darwin-amd64/dbee
	GOARCH=amd64 GOOS=windows go build -C src -o ../bin/windows-amd64/dbee.exe
	GOARCH=386 GOOS=linux go build -C src -o ../bin/linux-i386/dbee
	GOARCH=386 GOOS=freebsd go build -C src -o ../bin/freebsd-i386/dbee
	GOARCH=386 GOOS=windows go build -C src -o ../bin/windows-i386/dbee.exe
	GOARCH=arm64 GOOS=linux go build -C src -o ../bin/linux-arm64/dbee
	GOARCH=arm64 GOOS=freebsd go build -C src -o ../bin/freebsd-arm64/dbee
	GOARCH=arm64 GOOS=darwin go build -C src -o ../bin/darwin-arm64/dbee
	GOARCH=arm64 GOOS=windows go build -C src -o ../bin/windows-arm64/dbee.exe
	GOARCH=arm GOOS=linux go build -C src -o ../bin/linux-arm/dbee
	GOARCH=arm GOOS=freebsd go build -C src -o ../bin/freebsd-arm/dbee
	GOARCH=arm GOOS=windows go build -C src -o ../bin/windows-arm/dbee.exe

clean:
	go clean
	rm -rf ./bin/*

re: clean build
