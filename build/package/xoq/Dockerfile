FROM golang:1.19-alpine as build
WORKDIR /build
COPY cmd cmd
COPY pkg pkg
COPY go.mod go.mod
COPY go.sum go.sum
RUN go build ./cmd/xoq
FROM golang:1.19-alpine as run
WORKDIR /run
COPY --from=build /build/xoq .
RUN addgroup -S xoq-runner && adduser -S xoq-runner -G xoq-runner
USER xoq-runner
CMD ["./xoq"]