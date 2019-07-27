# Developer documentation

## Pre-requirements

- `curl`
- `git`
- `go`
- `make`

For a quick start, use the official `golang` Docker image (which has all these tools pre-installed), e.g.

```bash
docker run --rm -ti \
  --user 65534:65534 \
  --volume `pwd`:/go/src/github.com/Kong/konvoy/components/konvoy-control-plane \
  --workdir /go/src/github.com/Kong/konvoy/components/konvoy-control-plane \
  --env HOME=/tmp/home \
  --env GO111MODULE=on \
  golang:1.12.5 bash
export PATH=$HOME/bin:$PATH
```

## Helper commands

```bash
make help
```

## Installing dev tools

Run:

```bash
make dev/tools
```

which will install the following tools at `$HOME/bin`:

1. [Ginkgo](https://github.com/onsi/ginkgo#set-me-up) (BDD testing framework)
2. [Kubebuilder](https://book.kubebuilder.io/quick-start.html#installation) (Kubernetes API extension framework, comes with `etcd` and `kube-apiserver`)
3. [kustomize](https://book.kubebuilder.io/quick-start.html#installation) (Customization of kubernetes YAML configurations)
4. [kubectl](https://kubernetes.io/docs/tasks/tools/install-kubectl/#install-kubectl-binary-with-curl-on-linux) (Kubernetes API client)
5. [KIND](https://kind.sigs.k8s.io/docs/user/quick-start/#installation) (Kubernetes IN Docker)
6. [Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/#linux) (Kubernetes in VM)

ATTENTION: By default, development tools will be installed at `$HOME/bin`. Remember to include this directory into your `PATH`, 
e.g. by adding `export PATH=$HOME/bin:$PATH` line to the `$HOME/.bashrc` file.

## Building

Run:

```bash
make build
```

## Integration tests

 Integration tests will run all dependencies (ex. Postgres). Run:

 ```bash
make integration
```

## Running Control Plane on local machine

1. Run [KIND](https://kind.sigs.k8s.io/docs/user/quick-start) (Kubernetes IN Docker):

```bash
make start/k8s

# set KUBECONFIG for use by `konvoyctl` and `kubectl`
export KUBECONFIG="$(kind get kubeconfig-path --name=konvoy)"
```

2. Run `Control Plane` on local machine:

```bash
make run
```

3. Make a test `Discovery` request to `LDS`:

```bash
make curl/listeners
```

4. Make a test `Discovery` request to `CDS`:

```bash
make curl/clusters
```

## Pointing Envoy at Control Plane

1. Run [KIND](https://kind.sigs.k8s.io/docs/user/quick-start) (Kubernetes IN Docker):

```bash
make start/k8s

# set KUBECONFIG for use by `konvoyctl` and `kubectl`
export KUBECONFIG="$(kind get kubeconfig-path --name=konvoy)"
```

2. Start `Control Plane` on local machine:

```bash
make run
```

3. Start `Envoy` on local machine (requires `envoy` binary to be on your `PATH`):

```bash
make run/example/envoy
```

4. Dump effective `Envoy` config:

```bash
make config_dump/example/envoy
```

## Running Control Plane on Kubernetes

1. Run [KIND](https://kind.sigs.k8s.io/docs/user/quick-start) (Kubernetes IN Docker):

```bash
make start/k8s

# set KUBECONFIG for use by `konvoyctl` and `kubectl`
export KUBECONFIG="$(kind get kubeconfig-path --name=konvoy)"
```

2. Deploy `Control Plane` to [KIND](https://kind.sigs.k8s.io/docs/user/quick-start) (Kubernetes IN Docker):

```bash
make start/control-plane/k8s
```

3. Build `konvoyctl`

```bash
make build/konvoyctl
```

4. Add `Control Plane` to your `konvoyctl` config:

```bash
build/artifacts/konvoyctl/konvoyctl config control-planes add k8s
```

5. Verify that `Control Plane` has been added:

```bash
build/artifacts/konvoyctl/konvoyctl config control-planes list

NAME                      ENVIRONMENT
kubernetes-admin@konvoy   k8s
```

6. List `Dataplanes` connected to the `Control Plane`:

```bash
build/artifacts/konvoyctl/konvoyctl get dataplanes

MESH      NAMESPACE   NAME                        SUBSCRIPTIONS   LAST CONNECTED AGO   TOTAL UPDATES   TOTAL ERRORS
default               demo-app-685444477b-dnx9t   1               21m9s                2               0
```