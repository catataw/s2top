FROM quay.io/vektorcloud/go:1.12

RUN apk add --no-cache make

WORKDIR /app
COPY go.mod .
RUN go mod download

COPY . .
RUN make build && \
    mkdir -p /go/bin && \
    mv -v s2top /go/bin/

FROM scratch
ENV TERM=linux
COPY --from=0 /go/bin/s2top /s2top
ENTRYPOINT ["/s2top"]
