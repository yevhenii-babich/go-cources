############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder
# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/testapp/
COPY . .
# Build the binary.
RUN CGO_ENABLED=0 go build -o /go/bin/service
############################
# STEP 2 build a small image
############################
FROM jrbeverly/wkhtmltopdf:baseimage
COPY --from=builder /go/bin/service  /app/
ENV APP_TEST_APTH "/app/"
EXPOSE 8081
ENTRYPOINT ["/app/service"]
