FROM golang:1.18-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o grpc-pr-env-test-frontend

FROM golang:1.18-buster

WORKDIR /

COPY --from=build /app/grpc-pr-env-test-frontend /grpc-pr-env-test-frontend

EXPOSE 50052

CMD ["/grpc-pr-env-test-frontend"]