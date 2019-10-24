#!/bin/bash

rm -f ./qos-keys.js ./qos-keys.js.map
gopherjs build qos-keys.go -m

