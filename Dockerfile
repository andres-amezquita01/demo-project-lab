# Stage 1: Build
FROM public.ecr.aws/docker/library/golang:alpine3.18 AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
#RUN go get -u github.com/swaggo/swag/cmd/swag
#RUN go run /go/pkg/mod/github.com/swaggo/swag@v1.16.1/cmd/swag/main.go init
RUN go build -o main .

# Stage 2: Final Image
FROM public.ecr.aws/docker/library/alpine:3.14

COPY --from=builder /app/main /app/main
COPY ./boot.yaml /app/
RUN mkdir /app/docs
COPY ./docs/* /app/docs/
WORKDIR /app
CMD ["/app/main"]
