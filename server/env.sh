#!/bin/bash

cat .env.sample | while read line
do
  echo $line
  if [ $(echo $line | grep -e '=') ] && [ ${line:0:1} != '#' ]; then
    export $line
  else
    continue
  fi
done