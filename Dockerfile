# Start by building the application.
FROM golang:1.8 as build

WORKDIR /go/src/ocrClient
COPY . .
RUN go get -d -v ./...
RUN go install -v ./...

# Now copy it into our base image.
FROM gcr.io/distroless/base
COPY --from=build /go/src/ocrClient/config/config.json /config/config.json
COPY --from=build /go/src/ocrClient/template/ /template/
COPY --from=build /go/src/ocrClient/static/ /static/
COPY --from=build /go/bin/ocrClient /

ENTRYPOINT ["/ocrClient"]
EXPOSE 80
