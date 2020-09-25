#!/bin/bash

for d in ./packages/* ; do
  # Avoid running command on domain models
  if [ $d != "./packages/domain-models" ]
  then
    # If the package is python based
    if [ $d = "./packages/player-web-scraper" || $d = "./packages/team-web-scraper" ]
    then
      if []

    # If the package is golang based
    if [ $d = "./packages/player-stats" || $d = "./packages/team-stats" ]
    then
      if []

    # Change directory and run the second argument
    (cd "$d" && $1);
    if [[ $? -ne 0 ]]; then
      exit 1
    else
      continue
    fi
  fi
done