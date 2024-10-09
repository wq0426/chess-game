chess: cmd/server/main.go
	GOOS=linux GOARCH=amd64 go build -o chess cmd/server/main.go

.PHONY: upload
upload:
	scp chess root@120.53.45.195:/root/class

.PHONY: clean
clean:
	rm -f chess