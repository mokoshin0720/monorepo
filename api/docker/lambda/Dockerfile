FROM golang:1.20-alpine3.16 as build

ENV GO111MODULE=on

WORKDIR /go/src/app
# Copy dependencies list
COPY api/go.mod api/go.sum ./
# Build
COPY api/ .

RUN ls -la

RUN go build -o main cmd/server

# Copy artifacts to a clean image
FROM alpine:3.16
COPY --from=build /go/src/app/main /main
ENTRYPOINT [ "/main" ]