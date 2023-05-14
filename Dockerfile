FROM public.ecr.aws/docker/library/golang:alpine3.18
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
