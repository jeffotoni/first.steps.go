#################################################
# Dockerfile distroless
#################################################
FROM golang:1.15 as builder
WORKDIR /go/src/main
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go
RUN cp main /go/bin/main

############################
# STEP 2 build a small image
############################
FROM gcr.io/distroless/base
COPY --from=builder /go/bin/main /
CMD ["/main"]
