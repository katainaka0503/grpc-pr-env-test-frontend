FROM golang:1.16-buster AS build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build ./greeter_frontend/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /

COPY --from=build /main /greeter_frontend

EXPOSE 50052

USER nonroot:nonroot

ENTRYPOINT ["/greeter_frontend"]