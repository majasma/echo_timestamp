# get timestamp
time=$(date +%s)
msg=$(date)

printf "time in seconds $time\n"

# check if an hour has passed
if ! (($time % 3600))
then
    msg="an hour has passed..."
fi

printf "message $msg \n"

# set url
#url=http://localhost:port/ 
url=http://localhost:7893/

# create http request using curl
curl -X POST ${url} -H 'Content-type: application/json' --data "{\"text\": \"${msg}\"}"

$SHELL