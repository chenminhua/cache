git submodule add https://github.com/facebook/rocksdb rocksdb

git submodule update --init

cd rocksdb && make static_lib
(sudo apt install make g++ libz-dev libsnappy-dev)

[libsnappy-dev on mac](https://stackoverflow.com/questions/22794807/installing-avro-in-mac-os-x/24287418#24287418)



https://github.com/facebook/rocksdb/blob/5e221a98b5cbe65a40d945685821877b43f259bc/INSTALL.md
brew install rocksdb
/usr/local/Cellar/rocksdb




./cachebench -type tcp -n 100000 -r 100000 -t set
100000 records set
6.406644 seconds total
99% requests < 1 ms
99% requests < 2 ms
99% requests < 3 ms
100% requests < 23 ms
62 usec average for each request
throughput is 15.608797 MB


./cachebench -type tcp -n 100000 -r 100000 -t get
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 1
62930 records get
37070 records miss
0 records set
5.600641 seconds total
100% requests < 1 ms
54 usec average for each request
throughput is 11.236213 MB/s
rps is 17855.097910

./cachebench -type tcp -n 100000 -r 100000 -t get -P 3
type is tcp
server is localhost
total 100000 requests
data size is 1000
we have 1 connections
operation is get
keyspacelen is 100000
pipeline length is 3
63097 records get
36903 records miss
0 records set
3.063320 seconds total
99% requests < 1 ms
100% requests < 2 ms
87 usec average for each request
throughput is 20.597587 MB/s
rps is 32644.320115