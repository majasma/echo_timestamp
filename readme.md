# HTTP Echo server and client 

This repository containg the source code and executables for an HTTP echo server and client running on Ubuntu 22.04. The client script is activated every minute on the minute using systemd services. It uses curl to send an HTTP request containing a timestamp. The server echoes this back, and the client stores the body of the response in a log-file. 

## How to run this project?
1. clone repo
   ```
   git clone https://github.com/majasma/echo_timestamp
   ```
2. move client script from within the /echo_timestamp/client folder
   ```
   sudo mv -i timestamp.sh /usr/local/bin
   ```
3. create and enable service file for client, see [link](#how-to-make-it-run-every-minute)
4. move server executable from within the /echo_timestamp/server folder
   ```
   sudo mv -i echo_server /usr/local/bin
   ```
5. create and enable service, see [link](#how-to-make-it-start-at-login)
6. check log
   ```
   cat /var/log/client.log
   ```

## Server

### How it works
The ListenAndServe method starts a server at a given port. As the handler is set to nil, it uses the handler function added by HandleFunc. This function is defined as ServerHandler.

to echo it use http://localhost:port/

The server is set run on localhost, if it is set up on another device use http://"ipaddr":"port"/

### How to make it start at login?
Systemd is used to run the executable at login. In order for it to run you need to add a service.

```sudo vim /etc/systemd/system/server_echo.service```
and add the content of ./server/service_file.txt

In order to start the service:
```
systemctl daemon-reload
systemctl enable server_echo_timestamp.service
```

To check if the service is running: ```systemctl status server_echo_timestamp```

To avoid path corrections, we want the script to run from a generic location. Run
```sudo mv -i echo_server /usr/local/bin```

### ..and stop at logout?
cat etc/systemd/logind.conf,
KillUserProcesses should be set to "yes". 

```
sudo vim /etc/systemd/logind.conf
```

## Client script
Linux does not use the concept of executables, so it is only necessary to run a shell script or similar. 

By default the script does not have the right permissions, this needs to be changed. 
```chmod +x timestamp.sh```
This changes the permissions so that any user could execute it without any additional permissions. 

### How to make it run every minute?
Systemd is used to run the script every minute. In order for it to run you need to add a service and a timer.

For the service:
```sudo vim /etc/systemd/system/client_echo.service```
and add the content of ./client/service_file.txt

For the timer:
```sudo vim /etc/systemd/system/client_echo.timer```
and add the content of ./client/timer_file.txt

In order to start the service:

```
systemctl daemon-reload
systemctl enable client_echo.timer
systemctl enable client_echo.service
```

To check if the service is running: journalctl -fu client_echo.service

To avoid path corrections, we want the script to run from a generic location. Run
```sudo mv -i timestamp.sh /usr/local/bin```
to have it moved to /usr/local/bin. Log files are stored to /var/log/client.log

## Requirements
- A linux distro

### Server
- Golang (if you want to alter the source code)

### Client
- Curl package


## Challenges:
- Was not able to figure out how launch scripts using Cron. 
- When adding a startup script, my boot process stalled forever. Lesson learned about booting in recovery mode. 
- When using shebang on the client script, the modulo operation I had implemented using C arithmetic did not work. The shebang was necesarry so changed the modulo operation. 
- Was not able to get the services to stop on user logout or add the server startup on connection. Tried using specific users and groups in the systemd service files. For some reason this disables the server. 


## The strech goal
My plan was to implement a socket file called server_echo-http.socket and connect it with the server_echo service. Content of file:

```
[Socket]
ListenStream = 54321
BindIPv6Only = both
FileDescriptorName = http
Service=server_echo.service

[Install]
WantedBy = sockets.target                        
```
Error server_echo-http.socket: Failed with result 'service-start-limit-hit'.


[Unit]
Description=Run the server timestamp echo service
After=network.target server_echo-http.socket
Requires=server_echo-http.socket

[Service]
Type=oneshot
Group=majasma
ExecStart=/usr/local/bin/echo_server

[Install]
WantedBy=multi-user.target
