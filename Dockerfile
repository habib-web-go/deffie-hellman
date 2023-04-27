FROM golang:latest
WORKDIR /app
ARG AUTH_PORT
ENV AUTH_PORT $AUTH_PORT
EXPOSE $AUTH_PORT
COPY . .
RUN go build -o server .
CMD ["./server"]