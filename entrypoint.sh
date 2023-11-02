#!/bin/sh

# runs the binary file with hot reload that created 
# in Dockerfile which starts the server 
CompileDaemon -log-prefix=false -build="go build -o gotemplate" -command="./gotemplate"
