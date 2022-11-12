# Assignment for Disruptive Technologies
## Server
Written in Golang. Mainly uses the net/http package. 

### How it works
The ListenAndServe method starts a server at a given port. As the handler is set to nil, it uses the handler function added by HandleFunc. This function is defined as ServerHandler. ServerHandler sets the response header to StatusOK (200) and writes the request as a respond.

to echo it use http://localhost:port/

### Requirements:
 - go compiler

### Todo:
 - [ ] Find correct ports
 - [ ] Decide format of response
 - [x] Implement a simple server
 - [ ] Create a shell script ect to run script at login  
 - [ ] Shut down server at logout

### How to make it start at login?
adding the command to either ~/.profile or ~/.bashrc. bashrc would make it run every time a bash shell is started (probably not the solution). .profile is only run by interactive login shells. The /etc/profile contains Linux system wide environment and other startup scripts. Usually the default command line prompt is set in this file. It is used for all users logging in to the bash, ksh, or sh shells.

bash equivalent to "go run main.go"

put together the hppt request with curl?

### ..and stop at logout?
maybe https://www.cyberciti.biz/faq/linux-logout-user-howto/
## Client script
Linux does not use the concept of executables, so it is only necessary to run a shell script or similar. Crontab can be used to create a cron schedule. Commands
crontab -e 
*/1 * * * * /home/YOU/backup.sh (edit path)

### How it works
get unix timestamp
check if it is a full hour
create http request
listen for response
store response to log file

### Todo:
