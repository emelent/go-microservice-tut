# prepare build
FROM golang:1.12.6-stretch as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get -d -v ./...
RUN make build


FROM iron/base
COPY --from=build /app/accountservice accountservice
ENTRYPOINT ["./accountservice"]
EXPOSE 6767