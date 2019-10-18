#!/bin/bash

rm -f ./keys.js ./keys.js.map
gopherjs build keys.go -m

