# prepare build
FROM golang:1.12.6-stretch as build
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get -d -v ./...
RUN make docker-build


FROM iron/base
# ADD accountservice-linux-amd64 /
COPY --from=build /app/accountservice-linux-amd64 accountservice-linux-amd64 
EXPOSE 6767
ENTRYPOINT ["./accountservice-linux-amd64"]