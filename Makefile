all:
	go build -gcflags "-N -l" -o bin/gotour code.google.com/p/go-tour/gotour && bin/gotour
# go build -gcflags "-N -l" -o bin/gotour code.google.com/p/go-tour/gotour && pkill -f gotour ; nohup bin/gotour && tail -f 

test-go:
	go test code.google.com/p/go-tour
