FROM golang:1.22-alpine AS builder
WORKDIR /src/app/
COPY . .
RUN pwd && ls -la && cd metrics/go && go build app.go

FROM alpine:latest
RUN addgroup -g 1025 nonroot
RUN adduser -D nonroot -u 1025 -G nonroot
COPY --from=builder /src/app/metrics/go/app /usr/local/bin/gaia-metrics
EXPOSE 3000
USER nonroot

ENTRYPOINT ["gaia-metrics"]
