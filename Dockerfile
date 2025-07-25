# builder
FROM --platform=$BUILDPLATFORM golang:alpine AS build-env

ARG TARGETOS
ARG TARGETARCH

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH

WORKDIR /build
COPY . .
RUN go build -o sifar ./cmd/main.go

FROM scratch
WORKDIR /app
COPY --from=build-env /build/sifar /app/
EXPOSE 8000

ENTRYPOINT ["./sifar"]