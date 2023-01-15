FROM golang:1.17.7-bullseye as builder
RUN mkdir /build 
WORKDIR /build 
COPY go.* ./
RUN go mod download
COPY . .
RUN make build-prod
RUN ls config/

FROM scratch
COPY --from=builder /build/main /build/config/dev.yaml /app/

WORKDIR /app

EXPOSE 8080
CMD ["./main", "-e", "dev"]
