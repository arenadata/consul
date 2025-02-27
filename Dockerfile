FROM scratch

WORKDIR /data

ARG GOARCH=amd64

ENTRYPOINT ["/yp"]
CMD ["agent"]

COPY pkg/bin/linux_${GOARCH}/yp /
