
The magic test invocation

	$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp/showrunner -e "GOPATH=/usr/src/myapp" golang:1.7 go test
