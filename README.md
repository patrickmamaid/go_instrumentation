# go_instrumentation


Very quick boilerplate code to get you going on golang instrumentation


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