# lulupod
a micro file sharing and streaming http server to run across your local network.<br />
instead of using third party services (like *google drive, one-drive, icloud*, etc.), <br />
***lulupod*** lets you start up a simple client to share, download and watch files across your local network.<br />
of course, other then being isolated from the public internet, this also makes the sharing of a files significantly faster!

## usage (with docker; recommeded)
this example will spawn a *lulupod* service on your local ip adress, in port 8000 with **persistant storage**<br />
do the following commends:
```bash
git clone https://github.com/noam-g4/lulupod
cd lulupod 
docker build -t lulupod .
cd ~ 
mkdir lulupod-data
```
now to spin-up a container:
```bash
docker run -d --name lulupod \
-p 8000:8000 \
-e LULUPOD_URL=http://[your local ip]:8000 \
-v $HOME/lulupod-data:/app/public \
lulupod
```
from this point you're good to go, </br>
now you can use any device in your network and open your web-browser in `http://[the container host local ip]:8000`</br>
and start upload, download and watch files in your local network.<br />
since we've mounted a volume (`-v`) to this container, it means that we now have persistant storage. <br />
so even when the container is down, the data is in the host machine and running a new container with the same volume<br />
will persist to the new container. 

## use without docker (unrecommended)
you can build the binaries yourself (or run with `go run main.go`) <br />
in any way you need to have **go** installed on your machine
```bash
git clone https://github.com/noam-g4/lulupod
cd lulupod 
go mod tidy
go build -o main .
export LULUPOD_URL=http://[your local ip]:8000
./main
``` 
this will run lulupod directly on your machine, and teh storage will go to the `public` directory of this repo. <br />
if you're choosing to run it without docker, consider running it as a background process. 
