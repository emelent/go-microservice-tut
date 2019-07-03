# prepare build
FROM golang:1.12.6-stretch as build
RUN mkdir /app 
COPY ./ /app
WORKDIR /app/accounts
RUN make all


FROM iron/base
COPY --from=build /app/accounts/build/* /app/
RUN ls -lah /app/
ENTRYPOINT ["./app/accountservice"]
EXPOSE 6767
HEALTHCHECK --interval=5s --timeout=5s CMD ["./app/healthcheck", "-port=6767"] || exit 1
