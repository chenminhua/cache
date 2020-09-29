 ./test_basic -t 100000
total record number is 100000
value size is 1000
100000 records put in 689938 usec, 6.89938 usec average
100000 records get in 168648 usec, 1.68648 usec average

./ingest_data -t 100000 -s 1000
total record number is 100000
value size is 1000
100000 total records set in 616606 usec,6.16606 usec average, throughput 162.178 MB/s, rps is 162178

total record number is 1000000
value size is 1000
1000000 total records set in 6235130 usec,6.23513 usec average, throughput 160.382 MB/s, rps is 160382

./ingest_data -t 10000000 -s 1000
total record number is 10000000
value size is 1000
10000000 total records set in 64760728 usec,6.47607 usec average, throughput 154.415 MB/s, rps is 154415