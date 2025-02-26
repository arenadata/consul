FROM scratch

WORKDIR /data

ARG GOARCH=amd64

COPY pkg/bin/linux_${GOARCH}/consul /

ENTRYPOINT ["/consul"]
