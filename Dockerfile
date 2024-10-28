FROM golang:1.22.1-alpine3.19 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 make

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /bin/api /api

USER nonroot:nonroot

ENTRYPOINT ["/api"]
