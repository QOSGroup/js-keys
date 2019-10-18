#!/bin/bash

rm -f ./keys.js ./keys.js.map
gopherjs build hdpath.go -m

