# claws beacon


### simple beacon, first part is just a bind shell that can be connected to, in proc of adding more

service dir contains a .service file for systemd to configure the daemon. move `claws` into `/usr/sbin/claws` which should work to add it as a service, so just run `systemctl enable netconnect` and `systemctl start netconnect` to start the service.

build: `go build .`

usage: `./claws`

connect: `nc [host-addr] 9999`
