
using Go = import "go.capnp";

# Messages for pupil and master to talk to each other
$Go.package("messages");
$Go.import("github.com/sudharsh/monk");

@0x970b9de393122a12;

struct Pupil {
	uuid @0 :Text;
	url @1 :Text;
}

struct Response {
	success @0 :Bool;
}
