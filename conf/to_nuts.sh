#!/bin/bash


#model=$(cat conf/model.json)
#src/nats/examp/pub wb "$model"
#
#src/nats/examp/pub wb "so why should I feel crappy about something that makes me happy?"
#
#test=$(cat conf/test.json)
#src/nats/examp/pub wb "$test"
#
#src/nats/examp/pub wb ""

test_mpi=$(cat conf/test_multiple_items.json)
src/nats/examp/pub wb "$test_mpi"