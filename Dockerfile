# Use a minimal base image
# FROM debian:bullseye-slim

# # Install required packages
# RUN apt-get update && apt-get install -y alsa-utils mpg123

# # Copy your favorite song into the container
# COPY beep.mp3 /home/user/

# # Set the default command to play the audio file with mpg123
# CMD ["mpg123", "/home/user/beep.mp3"]

FROM golang:1.22.5

RUN apt-get update && apt-get install -y alsa-utils mpg123

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /goplayer

ENTRYPOINT ["/goplayer"]