#!/bin/bash
# get timestamp
time=$(date +%s)
msg=$(date)


printf "time in seconds $time\n"

# check if an hour has passed
if  [ $(echo "$time % 3600" | bc) -eq 0 ];
then
    msg="an hour has passed..."
fi

printf "message $msg \n"

# set url
#url=http://localhost:port/ 
url=http://localhost:54321/

# create http request using curl
(curl -X POST ${url} -H 'Content-type: application/json' --data "{\"text\": \"${msg}\"}" ) >> /home/majasma/echo_timestamp/client/client.log
echo "" >> /home/majasma/echo_timestamp/client/client.log


