
all:
	go build userbased.go
	./userbased
	go build itembased.go
	./itembased
