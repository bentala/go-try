FROM golang:alpine

ADD . /app

WORKDIR /app
# This builds a binary first
RUN ["go", "build"]
# Now you can run the executable and pass arguments at the run time.
ENTRYPOINT ["./app"]