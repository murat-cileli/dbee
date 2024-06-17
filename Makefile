build:
	GOARCH=amd64 GOOS=darwin go build -C src -o ../bin/darwin-amd64/dbee
	GOARCH=amd64 GOOS=linux go build -C src -o ../bin/linux-amd64/dbee
	GOARCH=amd64 GOOS=windows go build -C src -o ../bin/windows-amd64/dbee.exe

clean:
	go clean
	rm ./bin/darwin-amd64/dbee
	rm ./bin/linux-amd64/dbee
	rm ./bin/windows-amd64/dbee.exe