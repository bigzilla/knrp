# Knative Reverse Proxy

Knative Reverse Proxy (knrp) for local development of Knative application with/without magic DNS.

## Install

```bash
go install github.com/injustease/knrp@latest
```

## Usage

### Run Knative cluster

The easiest way is using [Knative quickstart](https://knative.dev/docs/getting-started/quickstart-install/).

### Fetch the Knative networking layer for `<gateway>`

Fetch the External IP.

```bash
# Kourier
kubectl --namespace kourier-system get service kourier

# Istio
kubectl --namespace istio-system get service istio-ingressgateway

# Contour
kubectl --namespace contour-external get service envoy
```

Example output Kourier (External IP is pending and it is intentional, use localhost instead), e.g. `localhost:80` or simply `localhost`.

```bash
NAME      TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)                      AGE
kourier   LoadBalancer   10.96.97.137   <pending>     80:32354/TCP,443:32488/TCP   33m
```

### Fetch the Knative service URL for `<service>`

Fetch the Service URL.

```bash
# kn
kn service list

# kubectl
kubectl get services.serving.knative.dev
```

Example output with service `hello` in the namespace `default`, e.g. `http://hello.default.example.com`.

```bash
# without magic DNS
NAME    URL                                       LATEST        AGE   CONDITIONS   READY   REASON
hello   http://hello.default.example.com   hello-00001   34m   3 OK / 3     True    

# with magic DNS
NAME    URL                                       LATEST        AGE   CONDITIONS   READY   REASON
hello   http://hello.default.127.0.0.1.sslip.io   hello-00001   34m   3 OK / 3     True    
```

### Run knrp

Command `knrp <gateway> <service>`

```bash
knrp localhost http://hello.default.127.0.0.1.sslip.io
```

knrp Output.

```bash
2022/04/06 16:05:19 knrp run at port: 52070
```

For more information run `knrp --help`.

### Expose to the internet with ngrok

You now expose Knative service to the internet with [ngrok](https://ngrok.com/) by using knrp address.

```bash
ngrok http 52070
```
