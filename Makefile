build:
	GOARCH=amd64 GOOS=darwin go build -C src -o ../bin/darwin-amd64/dbee
	GOARCH=amd64 GOOS=linux go build -C src -o ../bin/linux-amd64/dbee
	GOARCH=amd64 GOOS=windows go build -C src -o ../bin/windows-amd64/dbee.exe
	GOARCH=386 GOOS=linux go build -C src -o ../bin/linux-i386/dbee
	GOARCH=386 GOOS=windows go build -C src -o ../bin/windows-i386/dbee.exe

clean:
	go clean
	rm ./bin/darwin-amd64/dbee
	rm ./bin/linux-amd64/dbee
	rm ./bin/windows-amd64/dbee.exe
	rm ./bin/linux-i386/dbee
	rm ./bin/windows-i386/dbee.exe