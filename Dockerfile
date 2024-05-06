FROM golang:1.22.2-bullseye as builder
RUN mkdir /build 
ADD . /build/
WORKDIR /build 
RUN make build-prod
RUN ls config/

FROM scratch
COPY --from=builder /build/main /build/config/dev.yaml /app/

WORKDIR /app

EXPOSE 8080
CMD ["./main", "-e", "dev"]