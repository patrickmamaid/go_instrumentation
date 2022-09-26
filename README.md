# [Boilerplate] Prometheus instrumentation in go
## golang v1.18

`EST: 5 mins to understand and use`

Very quick boilerplate code to get you going on golang instrumentation with prometheus
1. Easily expose :2112/metrics for any service you want (you can change this port in main.go)
2. Paste into your implementation and make any modifications you deem necessary, this should get you going really quickly


## init:
```shell
# term1 
./init.sh
./run.sh

# term2
./curltest.sh
```

## Requirements:
```shell
brew install watch

```

## Quick notes:
main.go creates a basic counter `pmapp_infinite_ops_2s_total`:
- every 2s counter++
- grep it out by doing `curl -s http://localhost:2112/metrics | grep pmapp_infinite_ops_2s_total` / or run `curltest.sh`