FROM golang:alpine as base
Run apk add restic

FROM base as builder
ENV CGO_ENABLED 0
COPY . /src
WORKDIR /src
RUN go mod download
Run apk add restic
RUN go build -o assets/out github.com/chimano/restic-resource/cmd/out

From builder as tester
RUN /src/test.sh

FROM alpine as runner
Run apk add restic
COPY --from=builder /assets/out /opt/resource/out
RUN chmod +x /opt/resource/*
