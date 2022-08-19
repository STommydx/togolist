FROM gcr.io/distroless/base-debian11
COPY togolist /
ENTRYPOINT ["/togolist"]
