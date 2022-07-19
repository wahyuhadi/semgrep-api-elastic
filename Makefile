compile:
	echo "Compiling for every OS and Platform"
	GOOS=linux GOOS=linux GARCH=amd64 CGO_ENABLED=0 go build   -v -a -installsuffix cgo -o bin/semgrep-to-elastic
