# Software Building Blocks: Containers and the Tools to Run Them

This chapter introduces Docker containers, which are used to run sensors, actuators, and services as independent processes.

---

## Part 1 — Containers and Docker, a Short Guide

No prior container experience is required. The material here is sufficient to build, run, and connect the services used in the project. For more advanced topics, see the linked resources.

### Why containers?

A container packages an application together with everything it needs to run—libraries, runtime, and configuration—into a single isolated unit. The same container behaves consistently across different machines, reducing the classic “it works on my laptop” problem.

Containers are important in this course for three reasons:

- **Process isolation.** The project requires sensors, actuators, control agents, and supporting services to run as separate, independently restartable processes. Containers are the standard way to package and deploy such services (see Course Note 2 on service-oriented architectures).
- **Reproducibility.** Teammates, instructors, and examiners can run the same system and obtain the same behaviour with minimal setup.
- **Dependency management.** Databases, message brokers, dashboards, and application services can coexist without conflicting libraries or runtime environments.

Containers are often confused with virtual machines (VMs). A VM emulates an entire operating system, making it relatively large and resource-intensive. A container shares the host operating system kernel and therefore starts quickly, uses fewer resources, and allows many services to run efficiently on a single laptop. This makes containers particularly well suited for distributed systems composed of multiple independent services.

- What a container is: <https://www.docker.com/resources/what-container/>
- Docker overview: <https://docs.docker.com/get-started/docker-overview/>

### Three words to know

- **Image** , a read-only template (a built program plus its dependencies).
- **Container** , a running instance of an image. One image, many containers.
- **Registry** , a place images are stored and shared, such as Docker Hub.

The lifecycle ties them together: a `Dockerfile` is *built* into an image, an image is
*run* as containers, and images are *pulled* from (or *pushed* to) a registry.

<figure class="diagram">
<img src="figures/course-notes6-fig01.svg" alt="Container lifecycle: a Dockerfile is built into an image, the image is pulled from or pushed to a registry, and the image is run as containers">
<figcaption><em>The container lifecycle: build a Dockerfile into an image, pull or push images via a registry, and run an image as one or more containers.</em></figcaption>
</figure>

### Installing

- **Docker Desktop** (Mac, Windows, Linux): <https://docs.docker.com/get-started/get-docker/>
- **Docker Engine** (Linux servers, no GUI): <https://docs.docker.com/engine/install/>
- **Podman**, a drop-in, daemonless alternative whose commands mirror Docker's:
  <https://podman.io/>
- No install at all, a browser sandbox to experiment in:
  <https://labs.play-with-docker.com/>

Confirm the install with `docker version` and `docker run hello-world`.

### Example 1, run a container

```bash
docker run hello-world                 # pulls a tiny image and runs it
docker run -p 9000:80 nginx            # run a web server; open http://localhost:9000
docker ps                              # list running containers
docker logs <name>                     # see a container's output
docker stop <name>                     # stop it
```

`-p 9000:80` maps port 80 *inside* the container to port 9000 on the laptop. Port mapping
is how a service running in a container becomes reachable from the host.

### Example 2, a tiny Go sensor with a REST API

The first service worth building is a single one: a sensor packaged in its own container.
Before the code, a word on why the image must be **small**.

The lab is only complete when the **whole system runs together** on one laptop: BuildSim
plus a dozen or more sensor, actuator, agent, and supporting-service containers, all at
once. That only fits if the containers built in-house are small. Small images build,
restart, and ship faster, take less disk, and leave more room for the heavyweight services
(a database, a broker) that cannot be shrunk. Making the project's own services tiny is the
single biggest lever a team controls.

This is where **Go** earns its place as the recommended language. A Go program compiles to a
**single static binary** with no runtime, interpreter, or shared libraries to carry along.
The final image can therefore contain *just the binary*, a few megabytes, instead of a full
operating system or language runtime. A Python sensor ships an interpreter and its packages
(hundreds of megabytes); a JVM service ships a virtual machine; a Go sensor ships one file.

Four rules keep an image tiny:

- **Multi-stage build.** Compile in the full `golang` image, then copy only the binary into
  a minimal final image. The toolchain never ships.
- **Static binary.** Build with `CGO_ENABLED=0` so the binary has no C dependencies and can
  run on an empty base image.
- **Minimal base.** Use `scratch` (literally nothing) or `gcr.io/distroless/static` (adds CA
  certificates and time-zone data) for the final stage.
- **A `.dockerignore`.** Keep the build context, and the image, free of source noise, tests,
  and documents.

The example below is a full, self-contained sensor service. It exposes its current reading
over a REST API and updates that reading on a timer, the smallest realistic shape of a sensor
process. It uses only the Go standard library, so the module has **no dependencies** and the
image stays minimal.

`main.go`:

```go
package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"
	"sync"
	"time"
)

type Reading struct {
	Sensor string    `json:"sensor"`
	Room   string    `json:"room"`
	Value  float64   `json:"value"`
	Unit   string    `json:"unit"`
	Time   time.Time `json:"time"`
}

func main() {
	id := env("SENSOR_ID", "temp-A109")
	room := env("ROOM", "A109")
	port := env("PORT", "9000")

	var mu sync.RWMutex
	cur := Reading{Sensor: id, Room: room, Value: 21.0, Unit: "C", Time: time.Now()}

	// Update the reading every 2 s with a small random walk.
	go func() {
		for range time.Tick(2 * time.Second) {
			mu.Lock()
			cur.Value += (rand.Float64() - 0.5) * 0.3
			cur.Time = time.Now()
			mu.Unlock()
		}
	}()

	http.HandleFunc("/value", func(w http.ResponseWriter, r *http.Request) {
		mu.RLock()
		snapshot := cur
		mu.RUnlock()
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(snapshot)
	})
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("ok"))
	})

	log.Printf("sensor %s (room %s) listening on :%s", id, room, port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
```

`go.mod` defines the module. Generate it with `go mod init` (this example has no
third-party dependencies, so there is nothing else to add):

```bash
go mod init example.com/sensor
```

which writes:

```
module example.com/sensor

go 1.25
```

The `go` line records the installed Go version (`go mod init` sets it; check with
`go version`), so the value above will match whatever Go is installed.

`go mod init` prints a hint to run `go mod tidy`. That step only adds requirements for
**third-party** imports; this sensor uses only the standard library, so there is nothing to
add and `go mod tidy` can be skipped. (Once a real dependency is imported, `go mod tidy`
records it and produces a `go.sum`.)

The Docker build needs this file: `go run sensor.go` works on a single file without it, but
the image build runs `go mod download`, which requires a `go.mod`.

With `main.go` and `go.mod` in a folder, the sensor runs directly, **no Docker needed**.
Start it, and call it from a second terminal:

```bash
go run .                         # starts the sensor on :9000
curl -s localhost:9000/value     # {"sensor":"temp-A109","room":"A109","value":21,...}
```

To package the same program as a container, add a `.dockerignore`:

```
Dockerfile
.dockerignore
*.md
```

`Dockerfile`, a multi-stage build producing a binary on an empty base:

```dockerfile
# --- build stage ---
# Use a tag >= the `go` line in go.mod (run `go version` to check).
FROM golang:1.25 AS build
WORKDIR /src
COPY go.mod ./
RUN go mod download
COPY . .
# static binary, debug info stripped (-s -w) for a smaller file
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o /sensor .

# --- run stage ---
FROM scratch
COPY --from=build /sensor /sensor
EXPOSE 9000
ENTRYPOINT ["/sensor"]
```

The `golang:` tag in the build stage must be **the same or newer** than the `go` line in
`go.mod`. If it is older, the build stops at `go mod download` with
`go.mod requires go >= ...`; the fix is to bump the tag to match the installed Go version.

Build, run, and call it:

```bash
docker build -t sensor-temp .
docker run --rm -p 9000:9000 -e SENSOR_ID=temp-A109 -e ROOM=A109 sensor-temp &
curl -s localhost:9000/value
# {"sensor":"temp-A109","room":"A109","value":21.07,"unit":"C","time":"..."}
docker images sensor-temp        # FROM scratch: roughly 6 MB
```

The listen port is set by the `PORT` variable (default `9000`). If another local service
already holds that port, a clash is fixed by changing one value, for example
`PORT=9001 go run .` when running directly, or a different `-p` mapping under Docker.

- Dockerfile reference: <https://docs.docker.com/reference/dockerfile/>
- Multi-stage builds: <https://docs.docker.com/build/building/multi-stage/>
- Docker's Go language guide: <https://docs.docker.com/guides/golang/>
- Dockerfile best practices: <https://docs.docker.com/build/building/best-practices/>

### How small, and why it matters

The same program changes size dramatically with the final base image:

| Final base image | Approx. size | When to use it |
|---|---|---|
| `scratch` | ~6 MB | The default. Just the binary; no shell, no CA certificates. |
| `gcr.io/distroless/static` | ~8 MB | When the service makes **HTTPS** calls (adds CA certs) or needs time zones. |
| `alpine` | ~13 MB | When a shell is wanted for debugging inside the container. |
| `debian:slim` | ~80 MB | Only when **CGo** is required (for example the DuckDB driver). |
| `golang` (the build image) | ~800 MB | Never shipped, build-only. |

A 20-container system on `scratch` images costs a few hundred megabytes of disk; the same
system on full base images would cost gigabytes and start far more slowly. The pattern above,
standard-library Go plus a multi-stage build onto `scratch`, is the recommended template for
every sensor and actuator process in the project.

### Connecting it to the rest of the system

This example exposes its reading over REST so another service can poll it. In the real
project a sensor usually also **publishes** its reading, either to BuildSim over its REST API
(`PUT /api/sensors/{id}/value`, see `buildingsim/docs/lab-quickstart.md`) or to a message
broker (Part 3). Adding that is a few lines in the update loop using `net/http` or an MQTT
client; the container stays exactly as small, because those clients are pure Go.

### Persisting data

A container's filesystem is **ephemeral**: when it is removed, its data is gone. Anything
that must survive a restart, a database file, the Parquet data lake, goes in a **volume**:

```yaml
  timescale:
    image: timescale/timescaledb:latest-pg16
    volumes: ["tsdata:/var/lib/postgresql/data"]
volumes:
  tsdata:
```

- Volumes: <https://docs.docker.com/engine/storage/volumes/>
- Networking: <https://docs.docker.com/engine/network/>

### Example 3, run several services with Docker Compose

A single container is one process; the project is many. Once each service has its own small
image, **Docker Compose** describes them in one `compose.yaml` file and starts them together.

For Compose to build the sensor, its source from Example 2 lives in its own folder:

```
project/
├── compose.yaml
└── sensor/              # the files from Example 2
    ├── main.go
    ├── go.mod
    ├── .dockerignore
    └── Dockerfile
```

The file below builds that image once and runs **two instances** of it, a temperature and a
CO2 sensor, each configured through environment variables and published on a different host
port:

```yaml
services:
  temp-a109:
    build: ./sensor
    environment: { SENSOR_ID: "temp-A109", ROOM: "A109", PORT: "9000" }
    ports: ["9000:9000"]

  co2-a109:
    build: ./sensor
    environment: { SENSOR_ID: "co2-A109", ROOM: "A109", PORT: "9000" }
    ports: ["9001:9000"]
```

```bash
docker compose up --build        # build the image and start both sensors
curl -s localhost:9000/value     # the temperature sensor
curl -s localhost:9001/value     # the CO2 sensor
docker compose down              # stop and remove the containers
```

Recent Docker ships Compose as the built-in `docker compose` subcommand. Older installations
use a separate tool invoked as `docker-compose` (note the hyphen), with the same flags. If
`docker compose version` reports an error, use `docker-compose up --build` instead, or install
the plugin (<https://docs.docker.com/compose/install/>).

Compose places both containers on one network where each is reachable by its **service
name**, so a consumer added later can read `http://temp-a109:9000/value` without knowing any
IP address. A real broker (Mosquitto) and a data store join this file in the parts that
follow; the full recommended stack later in this chapter is exactly such a `compose.yaml`.

- Compose overview: <https://docs.docker.com/compose/>
- Compose quickstart: <https://docs.docker.com/compose/gettingstarted/>
- Compose file reference: <https://docs.docker.com/reference/compose-file/>

### A small command cheat sheet

| Command | What it does |
|---|---|
| `docker build -t name .` | Build an image from the Dockerfile in this folder |
| `docker run -p H:C img` | Run an image, mapping host port H to container port C |
| `docker ps` / `docker ps -a` | List running / all containers |
| `docker logs -f name` | Follow a container's output |
| `docker exec -it name sh` | Open a shell inside a running container |
| `docker compose up -d` | Start the whole stack in the background |
| `docker compose down -v` | Stop the stack and delete its volumes |
| `docker stats` | Live CPU and memory use per container |

- Official get-started workshop: <https://docs.docker.com/get-started/>
- One-page cheat sheet (PDF): <https://docs.docker.com/get-started/docker_cheatsheet.pdf>
- Docker Hub, find official images: <https://hub.docker.com/>

### A note on the laptop budget

Containers are light, but they are not free. Each service uses memory, and some images
(anything built on the Java Virtual Machine) can use half a gigabyte or more at rest.
`docker stats` shows the live usage. The chapters that follow recommend tools partly on how
little they cost to run, because the entire system has to fit on one machine. Keep the
container count and the heavyweight services modest.
