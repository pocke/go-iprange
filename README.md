go-iprange
==========

Installation
--------------

```sh
go get github.com/pocke/go-iprange
```


Usage
------

One IP.

```go
r, err := iprange.New("192.168.0.1")
if err != nil {
  panic(err)
}
r.IncludeStr("192.168.0.1") // => true
r.IncludeStr("192.168.0.2") // => false
```

IP with CIDR.

```go
r, err = iprange.New("192.168.0.0/24")
if err != nil {
  panic(err)
}
r.IncludeStr("192.168.0.1") // => true
r.IncludeStr("192.168.0.2") // => true
r.IncludeStr("192.168.1.2") // => false
```

Comma sepalated IP.

```go
r, err = iprange.New("192.168.0.0/24,172.0.0.0/16,192.168.1.1")
if err != nil {
  panic(err)
}
r.IncludeStr("192.168.0.1") // => true
r.IncludeStr("192.168.1.1") // => true
r.IncludeStr("172.0.10.11") // => true
```

IPv6

```go
r, err = iprange.New("2001:0db8:bd05:01d2:288a:1fc0:0001:10ee")
if err != nil {
  panic(err)
}
r.IncludeStr("2001:0db8:bd05:01d2:288a:1fc0:0001:10ee") // => false
r.IncludeStr("192.168.0.1") // => false
```

With TCP connection.

```go
r, err := iprange.New("192.168.0.1")
if err != nil {
  panic(err)
}

l, _ := net.ListenTCP("tcp", addr)
conn, err := l.Accept()
r.IncludeTCP(conn)
```
