FROM golang:latest as build

ADD . /src
WORKDIR /src
RUN mkdir -p dist && go build -o dist/example cmd/example/main.go  && chmod +x dist/example

FROM ubuntu:latest 
COPY --from=build /src/dist /app
RUN /app/example export bash