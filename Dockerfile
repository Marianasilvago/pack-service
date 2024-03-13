FROM golang:alpine as builder
WORKDIR /pack-svc
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o pack-svc cmd/*.go

FROM scratch
COPY --from=builder /pack-svc/pack-svc .
COPY ./frontend/index.html ./frontend/index.html
COPY .env .env

ENTRYPOINT ["./pack-svc", "-configFile", ".env", "http-serve"]
