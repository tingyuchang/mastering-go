# TCP/IP UDP, and Websocket


- the `net` package of Go Standard Library is all about TCP/IP, UDP, domain name resolution, and UNIX socket.
- `net.Dial()` as a client
- `net.Listen()` as a server
- `net.Dial()` and `net.Listen()`both return `net.Conn` which implement `io.Reader` and `io.Writer`.