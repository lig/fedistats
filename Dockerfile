FROM quay.io/centos/centos:stream9 as build
WORKDIR /usr/src/app

RUN dnf install -y golang

COPY . .
RUN go build -o ./build/ -v ./...


FROM quay.io/centos/centos:stream9 as app
ENTRYPOINT [ "/fedistats-collector" ]

COPY --from=build /usr/src/app/build/fedistats-collector /
