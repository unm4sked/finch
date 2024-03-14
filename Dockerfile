FROM golang:1.22 as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
RUN go mod tidy
COPY . ./

RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o /bin/api cmd/api/main.go 

FROM scratch

COPY --from=build /bin/api /bin/api
EXPOSE 3000

ENTRYPOINT ["./bin/api"]