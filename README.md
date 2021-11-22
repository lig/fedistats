# Fedistats 

https://gitlab.com/lig/fedistats


## About

`Fedistats` collects metrics of a Mastodon instance using the Public Timeline API.


## How to Use

### Build

```shell
go build -o ./build/ -v ./...
```

Creates `./build/fedistats-collector` binary.


## Configure

Create `./config/fedistats.yaml`.

```yaml
mastodon:
  server: "https://fosstodon.org"  # Required. The Mastodon instance to connect to.
  client_id: "<your app client id>"  # Required.
  client_secret: "<your app client secret>"  # Required.
  access_token: "<your app access token>"  # Required.
metrics:
  path: "/metrics"  # Optional. Path on which metrics are being served.
  port: 2112  # Optional. Port on which metrics are being served.
```


# Run

Launch using:

```shell
./build/fedistats-collector
```

Point your Prometheus instance to `http://your-host.tld:<configured port><configured path>`


## License

[GNU Affero General Public License v3.0](./LICENSE)
