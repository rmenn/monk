
using Go = import "go.capnp";

# Messages for pupil and master to talk to each other
$Go.package("messages");
$Go.import("github.com/sudharsh/monk");

@0x970b9de393122a12;

struct Map(Key, Value) {
  entries @0 :List(Entry);
  struct Entry {
    key @0 :Key;
    value @1 :Value;
  }
}

struct Pupil {
	uuid @0 :Text;
	url @1 :Text;
	facts @2 :Map(Text, Text);
}

struct Response {
	success @0 :Bool;
}


