# private ethereum network on kubernetes

## need the following:
1. persistent storage: when nodes fail, path on kubernetes worker
2. ethmonitor: stats and dashboard
3. deployment service file: ports, port forwarding, address forwarding, startup geth
4. deployment start: kubectl create -f ../manifests/geth-service.yaml
5.
