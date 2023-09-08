FROM golang:1.21 AS build-stage

WORKDIR /app

COPY . ./

RUN go mod tidy

RUN CGO_ENABLED=0 GOOS=linux go build -o service main.go

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build-stage /app/service /app/service

ENV PORT=58089

EXPOSE ${PORT}

#USER nonroot:nonroot

ENTRYPOINT ["/app/service"]