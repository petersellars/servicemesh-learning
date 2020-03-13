# Envoy Front-Proxy Example
This example creates a simple front-proxy topology, which sends traffic to two
backend services.

The example runs a straightforward Go HTTP server application, defined in the
[`example-service/service.go`](../example-service/service.go). An Envoy process
runs in the same container as a sidecar, configured with the
[`service-envoy.yaml`](./service-envoy.yaml) file. Finally, the
[`Dockerfile-service`](../example-service/Dockerfile-service) creates a
container that runs Envoy and the service on startup.

The front-proxy runs Envoy, configured with the
[`front-envoy.yaml`](./front-envoy.yaml) file, and uses the
[`Dockerfile-frontenvoy`](./Dockerfile-frontenvoy) as its container definition.

The [`docker-compose.yaml`](./docker-compose.yaml) file describes how to build,
package and run the front proxy and services together.

# Running the Front-Proxy Example
To run the front proxy example clone this repository and change to the
`front-proxy` directory.

```
$ git clone https://github.com/petersellars/servicemesh-learning.git
$ cd servicemesh-learning/envoy/front-proxy
```

## Build & Run the Example

To build our containers run:
```
docker-compose up --build -d
```

To stop, remove and clean up after running the example use:
```
docker-compose down -v
```

## Architecture

![envoy_font_proxy](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/petersellars/servicemesh-learning/envoy-initial/envoy/front-proxy/c4_component.puml)

## References

* Learn Envoy: [On your Laptop](https://www.envoyproxy.io/learn/on-your-laptop)
* Turbine Labs: [Getting Started with Envoy on your Laptop](https://blog.turbinelabs.io/getting-started-with-envoy-on-your-laptop-1b1a7073fd8e)
* Envoy GitHub: [Front-Proxy Example Code](https://github.com/envoyproxy/envoy/tree/release/v1.13/examples/front-proxy)
