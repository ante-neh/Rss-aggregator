build:
	@ go build -o bin/binary ./cmd/api/main.go

run:build
	@bin/binary 

	