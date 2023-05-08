#!/bin/bash

app=$APPLICATION
env=$TRACK

for i in "$@"
do
  case $i in
    --target=*)
      target="${i#*=}"
      shift # past argument=value
    ;;
    *)
      # ignore
      shift # past argument=value
    ;;
  esac
done

if [[ $env == "" ]]
then
  echo "No env mentioned."
  exit 2
fi

GITHUBIDFILE=id_rsa
if [ -f "GITHUBIDFILE" ]; then
    echo "GITHUBIDFILE exists"
else
    cp ~/.ssh/id_rsa id_rsa
fi