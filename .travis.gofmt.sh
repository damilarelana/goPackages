#!/bin/bash

if [ -n "$(gofmt -1 .)"]; then
  echo "Go code is not formatted:"
  gofmt -d .
  exit 1
fi