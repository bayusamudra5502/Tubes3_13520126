FROM golang:1.18

RUN mkdir -p /app
COPY . /app

RUN git clone https://github.com/githubnemo/CompileDaemon.git /CompileDaemon

WORKDIR /CompileDaemon
RUN go build -o /usr/bin/CompileDaemon .

WORKDIR /app
RUN go get

EXPOSE 8080

# Hot Reload
CMD CompileDaemon -log-prefix=false -build="go build" -command="./Tubes3_13520128"
