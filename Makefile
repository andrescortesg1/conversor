test:
	~/go/bin/gocyclo -ignore "_test" .
	go test -v -coverprofile cover.out .
	go tool cover -html cover.out -o cover.html

clean:
	rm cover.out cover.html
