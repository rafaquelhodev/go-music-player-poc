FROM golang:1.22.5

RUN apt-get update && apt-get install -y pulseaudio mpg123

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /goplayer

ENTRYPOINT ["/goplayer"]