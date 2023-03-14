COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .
RUN go build -o ./server ./server/cmd/server

EXPOSE 8080

CMD ["./server"]