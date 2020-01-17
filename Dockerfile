FROM golang:1.13.4
WORKDIR /app
COPY . .
RUN make build
EXPOSE 8080
CMD ["/app/main"]