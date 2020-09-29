cachebench -type http -n 100000 -r 100000 -t set

100000 records set
9.992472 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 5 ms
100% requests < 8 ms
97 usec average for each request
throughput is 10.007534 MB/s
rps is 10007.534001

cachebench -type http -n 100000 -r 100000 -t get

9.445663 seconds total
99% requests < 1 ms
100% requests < 3 ms
92 usec average for each request
throughput is 6.682855 MB/s
rps is 10586.869015


cachebench -type tcp -n 100000 -r 100000 -t set
100000 records set
5.142323 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
99% requests < 4 ms
100% requests < 5 ms
49 usec average for each request
throughput is 19.446463 MB/s
rps is 19446.463488

cachebench -type tcp -n 100000 -r 100000 -t get
4.653813 seconds total
100% requests < 1 ms
44 usec average for each request
throughput is 13.531914 MB/s
rps is 21487.755539