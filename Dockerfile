# create an image
FROM golang:1.20.4-alpine

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -v -o /app/vehiclerent

EXPOSE 3021

ENTRYPOINT [ "/app/vehiclerent" ]
CMD [ "serve" ]