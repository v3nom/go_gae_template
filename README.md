go_gae-template - Golang project template for Google App Engine
===============
Golang import path resolution might be quite confusing to newcomers and it gets even more so when you try to run on Google App Engine. This project template and setup guide should help programmers to start coding and deploying from day one.

#Golang template project checklist
* runs some code
* scalable structure
* unit testing

#Setup

#Runing
cd to app folder and execute:
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
