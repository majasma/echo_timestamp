# Assignment for Disruptive Technologies
## Server
Written in Golang. Mainly uses the net/http package. 

### How it works
The ListenAndServe method starts a server at a given port. As the handler is set to nil, it uses the handler function added by HandleFunc. This function is defined as ServerHandler. ServerHandler sets the response header to StatusOK (200) and writes the request as a respond. 

### Requirements:
 - go compiler

### Todo:
 - [ ] Find correct ports
 - [ ] Decide format
 - [x] Implement a simple server
 - [ ] Create a shell script ect to run script at login  

## Client script

### How it works

### Todo:
