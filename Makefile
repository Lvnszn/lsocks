GOCMD=go

server:
	$(GOCMD) run cmd/server/main.go
client:
	$(GOCMD) run cmd/client/main.go