git submodule add https://github.com/facebook/rocksdb rocksdb

git submodule update --init

cd rocksdb && make static_lib
(sudo apt install make g++ libz-dev libsnappy-dev)

[libsnappy-dev on mac](https://stackoverflow.com/questions/22794807/installing-avro-in-mac-os-x/24287418#24287418)