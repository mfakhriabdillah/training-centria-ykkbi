FROM golang:1.22.4-alpine as build
WORKDIR /app
COPY . .
RUN go build -o myapp
# EXPOSE 8000
# CMD [ "./myapp" ]

# Multistage
FROM alpine:latest
COPY --from=build /app/myapp myapp
EXPOSE 8000
CMD [ "./myapp" ]