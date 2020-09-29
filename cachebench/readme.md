### redis benchmark
redis-benchmark -c 1 -n 100000 -d 1000 -t set,get -r 100000

```
====== SET ======
  100000 requests completed in 3.23 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

99.90% <= 0.1 milliseconds
99.99% <= 0.2 milliseconds
99.99% <= 0.3 milliseconds
100.00% <= 0.4 milliseconds
100.00% <= 0.5 milliseconds
100.00% <= 0.6 milliseconds
100.00% <= 0.9 milliseconds
100.00% <= 2 milliseconds
30969.34 requests per second

====== GET ======
  100000 requests completed in 3.12 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

100.00% <= 0 milliseconds
32010.24 requests per second
```