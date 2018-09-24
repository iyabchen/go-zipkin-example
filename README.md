
# zipkin-go example

An example of zipkin-go based on code in go blog series http://callistaenterprise.se/blogg/teknik/2017/02/17/go-blog-series-part12/. In the above series, opentracing zipkin go library was used. This example uses [zipkin-go](https://zipkin.io/pages/existing_instrumentations.html) which supports V2 API instead.

### Build and Deploy 

It requires docker, docker-compose installed. From the same folder with Makefile, run

```sh
make image
# This will use 18080 and port 19411, and can be changed from docker-compose.yml.
docker-compose up -d
```
