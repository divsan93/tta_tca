#!/bin/bash

host="${host:-localhost:8080}"

curl -X POST ${host}/applications -d \
'{
    "name":"Pathfinder",
    "description": "Tackle Pathfinder application.",
    "repository": {
      "name": "tackle-pathfinder",
      "url": "https://github.com/djzager/tackle-pathfinder.git",
      "branch": "main"
    }
}' | jq -M .
