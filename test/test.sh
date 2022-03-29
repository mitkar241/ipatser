#!/bin/env bash

IP="localhost"
PORT="8000"

function header() {
    msg="$1"

    echo ""
    echo "##########"
    echo "# ${msg}"
    echo "##########"
}

header "Listing all Movies"
curl -sS -X GET "http://${IP}:${PORT}/movies" | jq


header "Adding Movies"
curl -sS -X POST "http://${IP}:${PORT}/movies?movieid=4&moviename=movie4"
curl -sS -X POST "http://${IP}:${PORT}/movies?movieid=1&moviename=movie3"
curl -sS -X POST "http://${IP}:${PORT}/movies?movieid=2&moviename=movie1"
curl -sS -X POST "http://${IP}:${PORT}/movies?movieid=3&moviename=movie2"

header "Listing all Movies"
curl -sS -X GET "http://${IP}:${PORT}/movies" | jq


header "Deleting Movie by ID"
curl -sS -X DELETE "http://${IP}:${PORT}/movies/1"

header "Listing all Movies"
curl -sS -X GET "http://${IP}:${PORT}/movies" | jq


header "Deleting all Movies"
curl -sS -X DELETE "http://${IP}:${PORT}/movies"

header "Listing all Movies"
curl -sS -X GET "http://${IP}:${PORT}/movies" | jq
