FROM scratch

WORKDIR /data

ARG GOARCH=amd64

ENTRYPOINT ["/consul"]
CMD ["agent"]

COPY pkg/bin/linux_${GOARCH}/consul /
