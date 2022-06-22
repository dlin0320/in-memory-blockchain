FROM golang
RUN mkdir -p /server_image
WORKDIR /server_image
COPY . .
RUN go mod download
RUN go build -o server_image
ENTRYPOINT ["./server_image"]

