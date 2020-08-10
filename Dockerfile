FROM golang:1.13-alpine as builder

RUN apk --no-cache add make git bash fish ca-certificates wget

# This configures all the go tooling to resolve dependencies from the vendor/
# folder instead of trying to download dependencies.  We currently want the
# dockerized build to be completely self-contained, since we currently don't
# have a way for builds running on our CI servers to download dependencies
# from internal git servers which require authentication.  At some point,
# we should set up an Artifactory as a go mod proxy, so the build only needs
# authenticate and download from that.
ENV GOFLAGS="-mod=vendor"
ENV CGO_ENABLED=0

WORKDIR /project

COPY ./Makefile ./go.mod ./go.sum /project/
COPY ./vendor /project/vendor

COPY ./ /project

FROM builder as build

RUN make clean build


FROM alpine:3.10

WORKDIR /app

COPY --from=build /project/build/$TARGET /project/website   /app/

COPY --from=build /project/website/*   /app/website/


CMD ./encryptgig
