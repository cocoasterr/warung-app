## Init Your Go Module
```bash
go mod init github.com/cocoasterr/warungapp
```
## Download Module and create Go binary file
```bash
make build
```
## Docker Build
```bash
docker build -t name-image:tag .
```
Example like this:
 ```bash 
 - docker build -t warung-app:1.0 .
```
## Docker run
```bash
docker run -p outer-port:inner-port nama-image:tag
```
Notes:
- Inner-port using for golang app
- and outer-port using for browser to acces the golang app

example like this:
 ```bash 
docker run -p 7777:8888 nama-image:tag
```
- my golang app using port 8888
- and i wanna access golang app using 7777 
