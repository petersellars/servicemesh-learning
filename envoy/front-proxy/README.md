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

## Running the Front-Proxy Example
To run the front proxy example clone this repository and change to the
`front-proxy` directory.

```
$ git clone https://github.com/petersellars/servicemesh-learning.git
$ cd servicemesh-learning/envoy/front-proxy
```

To build our containers, run:

```
$ docker-compose up --build -d
```

This command starts a single instance of the front proxy and two service
instances, one configured as "service1" and the other as "service2", `--build`
means to build the containers before starting up, and `-d` means run them in
detached mode.

Running `docker-compose ps` should show the following output:

```
$ docker-compose ps
          Name                         Command               State                            Ports                         
----------------------------------------------------------------------------------------------------------------------------
front-proxy_front-envoy_1   /docker-entrypoint.sh /bin ...   Up      10000/tcp, 0.0.0.0:8000->80/tcp, 0.0.0.0:8001->8001/tcp
front-proxy_service1_1      /bin/sh -c /usr/local/bin/ ...   Up      10000/tcp, 80/tcp                                      
front-proxy_service2_1      /bin/sh -c /usr/local/bin/ ...   Up      10000/tcp, 80/tcp  
```

## Sending Traffic
Docker Compose has mapped port 8000 on the front-proxy to your local network.
Open your browser to http://localhsot:8000/service/1, or run `curl
localhost:8000/service/1`. You should see:

```
$ curl localhost:8000/service/1
Hello from behind Envoy (service 1)! hostname: 7280438e7e3d resolvedhostname: 172.18.0.2
```

Going to http://localhost:8000/service/2 should result in:

```
curl localhost:8000/service/2
Hello from behind Envoy (service 2)! hostname: d90edb93e196 resolvedhostname: 172.18.0.3
```

You're connecting to Envoy, operating as a front proxy, which is, in turn,
sending requests to service 1 or service 2

## Envoy Configuration
This example configures Envoy statically for demonstration purposes. Other
examples here will harness its power by dynamically configuring it.

Let's look at how we configured Envoy. To get the right services set up, Docker
Compose looks at the [`docker-compose.yaml`](./docker-compose.yaml) file.
You'll see the following definition for the `front-envoy` service:

```
  front-envoy:
    build:
      context: .
      dockerfile: Dockerfile-frontenvoy
    volumes:
      - ./front-envoy.yaml:/etc/front-envoy.yaml
    networks:
      - envoymesh
    expose:
      - "80"
      - "8001"
    ports:
      - "8000:80"
      - "8001:8001"
```

Let's examine this from top to bottom:

1. Build a container using the `Dockerfile-frontenvoy` file located in the
   current directory
2. Mount the `front-envoy.yaml` file in this directory as
   `/etc/front-envoy.yaml`
3. Create and use a Docker network named "`envoymesh`" for this container
4. Expose ports 80 (for general traffic) and 8001 (for the admin server)
5. Map the host port 8000 to container port 80, and the host port 8001 to
   container port 8001

You'll see the following definition for the `service` service:

```
  service1:
    build:
      context: ../example-service
      dockerfile: Dockerfile-service
    volumes:
      - ./service-envoy.yaml:/etc/service-envoy.yaml
    networks:
      envoymesh:
        aliases:
          - service1
    environment:
      - SERVICE_NAME=1
    expose:
      - "80"
```

Let's examine this from top to bottom:

1. Build a container using the `Dockerfile-service` file located in the `../example-service` directory
2. Mount the `service-envoy.yaml` file in the this directory as `/etc/service-envoy.yaml`
3. Create and use a Docker network named "`envoymesh`" for this container and provide the `service1` alias (alternative hostname) within the network
4. Set the environment variable `SERVICE_NAME` based on the service number
5. Expose port 80 (for general traffic)

### Front-Proxy Envoy Configuration

### Service Envoy Configuration

## Stopping the Front-Proxy Example

To stop, remove and clean the example containers, run:
```
docker-compose down -v
```

## Architecture

![envoy_font_proxy](http://www.plantuml.com/plantuml/proxy?cache=no&src=https://raw.githubusercontent.com/petersellars/servicemesh-learning/envoy-initial/envoy/front-proxy/c4_component.puml)

## References

* Learn Envoy: [On your Laptop](https://www.envoyproxy.io/learn/on-your-laptop)
* Turbine Labs: [Getting Started with Envoy on your Laptop](https://blog.turbinelabs.io/getting-started-with-envoy-on-your-laptop-1b1a7073fd8e)
* Envoy GitHub: [Front-Proxy Example Code](https://github.com/envoyproxy/envoy/tree/release/v1.13/examples/front-proxy)
