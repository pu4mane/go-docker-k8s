FROM alpine:3.20

RUN apk update

COPY demo /
COPY configs/config.yaml /
CMD ["./demo", "-config=config.yaml"]