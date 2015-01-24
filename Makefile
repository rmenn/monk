all: monk

monk: messages/monk.capnp.go
	go build

messages/monk.capnp.go: 
	cd messages; capnp compile -ogo monk.capnp

clean:
	rm -rf monk messages/*.capnp.go
