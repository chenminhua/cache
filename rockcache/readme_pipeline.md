redis-benchmark -c 1 -n 100000 -d 1000 -t set,get -r 100000 -P 3

redis-benchmark -c 1 -n 100000 -d 1000 -t set,get -r 100000 -P 3
====== SET ======
  100002 requests completed in 1.22 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

99.65% <= 0.1 milliseconds
99.95% <= 0.2 milliseconds
99.97% <= 0.3 milliseconds
99.97% <= 0.4 milliseconds
99.98% <= 0.7 milliseconds
99.98% <= 0.9 milliseconds
99.99% <= 2 milliseconds
100.00% <= 2 milliseconds
81968.85 requests per second

====== GET ======
  100002 requests completed in 1.12 seconds
  1 parallel clients
  1000 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

100.00% <= 0 milliseconds
89207.84 requests per second