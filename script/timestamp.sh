# get timestamp
time=$(date +%s)
message=date -d @$(date -u +%s) 

# set url
url=http://localhost:port/ 

# check if an hour has passed
if ! ((TIME % 60))
then
    message = "an hour has passed..."
fi

# create http request using curl
curl -X POST ${url} -H 'Content-type: application/json' --data "{\"text\": \"${message}\"}"

#conversion of dates
#https://www.codegrepper.com/code-examples/shell/bash+script+get+unix+timestamp