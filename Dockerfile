FROM golang:1.21-alpine as builder

WORKDIR /go/src/github.com/thoughtgears/action-iac-generator
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o build/app .

# A distroless container image with some basics like SSL certificates
# https://github.com/GoogleContainerTools/distroless
FROM gcr.io/distroless/static
COPY /templates /templates
COPY --from=builder /go/src/github.com/thoughtgears/action-iac-generator/build/app /app
ENTRYPOINT ["/app"]