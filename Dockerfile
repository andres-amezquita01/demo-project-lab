# Stage 1: Build
FROM public.ecr.aws/docker/library/golang:alpine3.18 AS builder

RUN mkdir /app
COPY . /app
WORKDIR /app
RUN go build -o main .

# Stage 2: Final Image
FROM public.ecr.aws/docker/library/alpine:3.14

COPY --from=builder /app/main /app/main

CMD ["/app/main"]
