FROM golang:alpine as base
Run apk add restic

FROM base as builder
ENV CGO_ENABLED 0
COPY . /src
WORKDIR /src
RUN go mod download
Run apk add restic
RUN go build -o assets/out github.com/chimano/restic-resource/cmd/out
RUN go build -o assets/in github.com/chimano/restic-resource/cmd/in
RUN go build -o assets/check github.com/chimano/restic-resource/cmd/check

From builder as tester
RUN /src/test.sh

FROM alpine as runner
Run apk add restic
COPY --from=builder /assets/out /opt/resource/out
COPY --from=builder /assets/in /opt/resource/in
COPY --from=builder /assets/check /opt/resource/check
RUN chmod +x /opt/resource/*
