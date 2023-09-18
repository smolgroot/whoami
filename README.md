# Whoami - Geo IP Blocking website

**Whoami** is a simple Go HTTP server providing a "What's my IP?" service.
You can use it for testing purpose to see if your IP address is effectively hidden behind a VPN or a Proxy (HTTP, SOCKS, etc.).

The app is serving a "fake" video streaming website called *Netblix*, which is known to geo-IP-block some of its content.

## Build

To build the app use 

```
go build .
```

and to run it use 

```
./whoami
```

## Geo IP Blocking

The app allows you to watch a movie if your IP is different from:

- `127.0.0.1`
- `::1`

By accessing the website through `0.0.0.0:8088` or `localhost:8088` you should be blocked, but with another network interface, you should be allowed.

> Coming: config file to customize IP blacklist and listening address and port

