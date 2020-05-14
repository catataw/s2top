# Build

To build `s2top` from source, ensure you have [dep](https://github.com/golang/dep) installed and run:

```bash
go get github.com/catataw/s2top && \
cd $GOPATH/src/github.com/catataw/s2top && \
make build
```

To build a minimal Docker image containing only `s2top`:
```bash
make image
```

Now you can run your local image:

```bash
docker run -ti --name s2top --rm -v /var/run/docker.sock:/var/run/docker.sock s2top
```
