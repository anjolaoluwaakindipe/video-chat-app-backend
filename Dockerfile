FROM golang:alpine
RUN apk add git
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
RUN go build -o ./output/dist


CMD ["./output/dist"]
