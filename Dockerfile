FROM golang:alpine as builder
COPY . /restic-resource
WORKDIR /restic-resource
RUN go build -tags check -o /assets/check
RUN go build -tags out -o /assets/out
RUN go build -tags in -o /assets/in

FROM alpine as runner
COPY --from=builder /assets/check /opt/resource/check
COPY --from=builder /assets/out /opt/resource/out
COPY --from=builder /assets/in /opt/resource/in