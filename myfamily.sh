curl  https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq --arg ID_num "$Hero_ID" '.[] | select( .id ==($ID_num|tonumber)) | .connections.relatives' | sed 's/"//g'
