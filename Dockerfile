FROM golang:1.24-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux  go build -o ./scan ./cmd

FROM golang:1.24-alpine AS release

COPY --from=build /app/scan /
COPY --from=build /app/bin/scan.sh /

ENTRYPOINT ["/scan.sh"]
