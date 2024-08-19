FROM golang:1.22.5

RUN apt-get update && apt-get install -y pulseaudio mpg123

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

CMD ["tail", "-f", "/dev/null"]