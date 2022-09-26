#!/bin/bash

watch -d $(curl -s http://localhost:2112/metrics | grep pmapp_infinite_ops_2s_total)


