# Assignment for Disruptive Technologies
## Server

### How it works
The ListenAndServe method starts a server at a given port. As the handler is set to nil, it uses the handler function added by HandleFunc. This function is defined as ServerHandler. .

to echo it use http://localhost:port/

The server is set run on localhost, if it is set up on another device use http://"ipaddr":"port"/

### Todo:
 - [x] Find correct ports
 - [x] Decide format of response
 - [x] Implement a simple server
 - [ ] Create a shell script ect to run script at login  
 - [ ] Shut down server at logout

### How to make it start at login?
Systemd is used to run the executable at login. In order for it to run you need to add a service.

run "sudo vim /etc/systemd/system/server_echo.service"
and add the content of ./server/service_file.txt

In order to start the service:
- sudo systemctl daemon-reload
- systemctl enable server_echo_timestamp.service

To check if the service is running: systemctl status server_echo_timestamp

To avoid path corrections, we want the script to run from a generic location. Run
sudo mv -i echo_server /usr/local/bin

### ..and stop at logout?
cat etc/systemd/logind.conf,
KillUserProcesses should be set to "yes". 

sudo vim /etc/systemd/logind.conf


## Client script
Linux does not use the concept of executables, so it is only necessary to run a shell script or similar. 

By default the script does not have the right permissions, this needs to be changed. "chmod +x timestamp.sh"
This changes the permissions so that any user could execute it without any additional permissions. 

Systemd is used to run the script every minute. In order for it to run you need to add a service and a timer.

For the service:
run "sudo vim /etc/systemd/system/client_echo.service"
and add the content of ./client/service_file.txt

For the timer
run "sudo vim /etc/systemd/system/client_echo.timer"
and add the content of ./client/timer_file.txt

In order to start the service:
- systemctl daemon-reload
- systemctl enable client_echo.timer
- systemctl enable client_echo.service

To check if the service is running: journalctl -fu client_echo.service

To avoid path corrections, we want the script to run from a generic location. Run
sudo mv -i timestamp.sh /usr/local/bin
to have it moved to /usr/local/bin. Log files are stored to /var/log/client.log

### Todo:
[x] create script to schedule execution
[x] retrieve unix time and date
[x] implement check for full hour
[x] send http request using curl
[x] store response to log-file
[x] alter execute permissions
[ ] check that permissions are changed for other devices as well
[x] alter paths
[ ] change timing to start on the minute

Challenges:
- Was not able to figure out how cron worked properly
- When adding a startupscript I destroyed my entire virtual machine, with all of my work on, lesson learned:)
- When adding the interpreter comment on the client script, the modulo operation does not work as this is c-based
- When running with bash-spcific it cant interpret the c-translation

## Requirements
### Server
- golang?

### Client
- curl package
- added files in the systemd-folder