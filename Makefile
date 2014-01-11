all:
	go build -gcflags "-N -l" -o bin/gotour code.google.com/p/go-tour/gotour && bin/gotour

test-go:
	go test code.google.com/p/go-tour
