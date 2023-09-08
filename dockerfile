FROM golang:latest as builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN CGO_ENABLED=0 go build -o main .
#EXPOSE 80
#CMD [ "./main" ]

FROM gcr.io/distroless/static-debian11
COPY --from=builder /app/main /
EXPOSE 80
CMD ["/main"]