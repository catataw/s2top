# Debug Mode

`s2top` comes with a built-in logging facility and local socket server to simplify debugging at run time.

## Quick Start

If running `s2top` via Docker, debug logging can be most easily enabled as below:
```bash
docker run -ti --rm \
           --name=s2top \
           -e S2TOP_DEBUG=1 \
           -e S2TOP_DEBUG_TCP=1 \
           -p 9000:9000 \
           -v /var/run/docker.sock:/var/run/docker.sock \
           quay.io/vektorlab/s2top:latest
```

Log messages can be followed by connecting to the default listen address:
```bash
curl -s localhost:9000
```

example output:
```
15:06:43.881 ▶ NOTI 002 logger initialized
15:06:43.881 ▶ INFO 003 loaded config param: "filterStr": ""
15:06:43.881 ▶ INFO 004 loaded config param: "sortField": "state"
15:06:43.881 ▶ INFO 005 loaded config switch: "sortReversed": false
15:06:43.881 ▶ INFO 006 loaded config switch: "allContainers": true
15:06:43.881 ▶ INFO 007 loaded config switch: "enableHeader": true
15:06:43.883 ▶ INFO 008 collector started for container: 7120f83ca...
...
```

## Unix Socket

Debug mode is enabled via the `S2TOP_DEBUG` environment variable:

```bash
S2TOP_DEBUG=1 ./s2top
```

While `s2top` is running, you can connect to the logging socket via socat or similar tools:
```bash
socat unix-connect:./s2top.sock stdio
```

## TCP Logging Socket

In lieu of using a local unix socket, TCP logging can be enabled via the `S2TOP_DEBUG_TCP` environment variable:

```bash
S2TOP_DEBUG=1 S2TOP_DEBUG_TCP=1 ./s2top
```

A TCP listener for streaming log messages will be started on the default listen address(`0.0.0.0:9000`)
