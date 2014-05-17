Golang project template for Google App Engine
===============
Golang import path resolution might be quite confusing to newcomers and it gets even more so when you try to run on Google App Engine. This project template and setup guide should help programmers to start coding and deploying from day one.

#Golang template project checklist
* runs some code
* scalable structure
* unit testing

#Setup
  I assume that Google App Engine Go SDK & Go SDK are installed and required folders are added to your PATH.
  * [http://golang.org/doc/install](http://golang.org/doc/install)
  * [https://developers.google.com/appengine/downloads](https://developers.google.com/appengine/downloads)

## Folder structure
  TODO

## App name
Open app.yaml and change "application: test" to "application: myAppName"

#Runing
cd to app folder (folder with app.yaml file) and execute:
```bash
goapp serve
```
#Testing
cd to project folder and execute:
```bash 
goapp test ./... 
```
#Deploying
cd to app folder (folder with app.yaml file) and execute:
```bash
goapp deploy
```
