FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /grpc-pr-env-test-frontend /grpc-pr-env-test-frontend

EXPOSE 50052

USER nonroot:nonroot

ENTRYPOINT ["/greeter_frontend"]