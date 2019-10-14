url=https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json
curl -s ${url} | jq --arg HERO_ID "$HERO_ID" '.[] |select (.id == ($HERO_ID|tonumber)) |.connections.relatives' | sed 's/'//g'
