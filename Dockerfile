FROM --platform=$BUILDPLATFORM golang:1.23-alpine AS builder

ARG TARGETOS
ARG TARGETARCH

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=$TARGETOS \
    GOARCH=$TARGETARCH

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

ADD cmd/ ./cmd/
ADD pkg/ ./pkg/

RUN go build -o sifar ./cmd/main.go

FROM scratch
COPY --from=builder /build/sifar /bin/sifar
EXPOSE 8000

ENTRYPOINT ["/bin/sifar"]
