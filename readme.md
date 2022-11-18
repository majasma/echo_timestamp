# Assignment for Disruptive Technologies
## Server

### How it works
The ListenAndServe method starts a server at a given port. As the handler is set to nil, it uses the handler function added by HandleFunc. This function is defined as ServerHandler. ServerHandler sets the response header to StatusOK (200) and writes the request as a respond.

to echo it use http://localhost:port/

The server is set run on localhost, if it is set up on another device use http://"ipaddr":"port"/

### Requirements:
 - go compiler

### Todo:
 - [x] Find correct ports
 - [x] Decide format of response
 - [x] Implement a simple server
 - [ ] Create a shell script ect to run script at login  
 - [ ] Shut down server at logout

### How to make it start at login?
adding the command to either ~/.profile or ~/.bashrc. bashrc would make it run every time a bash shell is started (probably not the solution). .profile is only run by interactive login shells. The /etc/profile contains Linux system wide environment and other startup scripts. Usually the default command line prompt is set in this file. It is used for all users logging in to the bash, ksh, or sh shells.

bash equivalent to "go run main.go"

### ..and stop at logout?
maybe https://www.cyberciti.biz/faq/linux-logout-user-howto/

## Client script
Linux does not use the concept of executables, so it is only necessary to run a shell script or similar. Crontab can be used to create a cron schedule. Commands
crontab -e 
*/1 * * * * /home/YOU/backup.sh (edit path)

By default the script does not have the right permissions, this needs to be changed. 
this is changed with chmod. 
"chmod +x timestamp.sh"
This changes the permissions so that any user could execute it without any additional permissions. 


### Todo:
[ ] create script to schedule execution
[x] retrieve unix time and date
[x] implement check for full hour
[x] send http request using curl
[x] store response to log-file
[x] alter execute permissions
[ ] check that permissions are changed for other devices as well

Challenges:
- Was not able to figure out how cron worked properly
- When adding a startupscript I destroyed my entire virtual machine, with all of my work on, lesson learned:)
- When adding the interpreter comment on the client script, the modulo operation does not work as this is c-based