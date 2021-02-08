FROM golang:alpine AS builder

# set image variables
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# move to：/build
WORKDIR /build

# copy go.mod 和 go.sum and download package
COPY go.mod .
COPY go.sum .
RUN go mod download

# copy code into coontainer
COPY . .

# build project o2o-center-api 
RUN go build -o o2o-center-api .

###################
# create a small image
###################
FROM scratch

#copy all static file into small image
COPY ./templates /templates
COPY ./static /static
COPY ./conf /conf

# from builder image, cpoy /dist/app into current folder
COPY --from=builder /build/o2o-center-api /

# the command that need to be run
#ENTRYPOINT ["/o2o-center-api", "conf/config.ini"]