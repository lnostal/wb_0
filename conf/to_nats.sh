#!/bin/bash


model=$(cat conf/model.json)
conf/pub wb "$model"

conf/pub wb "so why should I feel crappy about something that makes me happy?"

test=$(cat conf/test.json)
conf/pub wb "$test"

conf/pub wb ""

test_mpi=$(cat conf/test_multiple_items.json)
conf/pub wb "$test_mpi"