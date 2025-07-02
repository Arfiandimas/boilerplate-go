#!/bin/bash
vault kv get -format=json kaj/stg/{{PROJECT_NAME}} | jq -r '.data.data | to_entries | .[] | "\(.key)=\(.value)"' >> .env