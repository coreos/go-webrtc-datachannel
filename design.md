# go-webrtc-datachannel design
There are two servers needed for the system: STUN server and signalling
server. We can use Google's STUN server (or any other open STUN server) for
the NAT punching purpose. For signaling server, we had better build one by
ourself.

### Signaling Server
At first, we should design the protocol between signaling servers and clients.

It might be something looking like DNS.

I prefer to use HTTP protocol and use similar the API as go-etct.

I think we can use memcache to store these temporary information, and save
username and password to database like MySQL.

It supports several operations. (This is a draft protocol. We might revise it
accordingly.)

###### Update NAT punching information
Client sends `http://localhost:7446/set/?token=${TOKEN}&&host=${HOST}&&timeout=${TIMEOUT}`, receives "OK".

###### Request other server's IP and port
Client sends `http://localhost:7446/get/?token=${TOKEN}&&uid=${UID}`, receives a json object with uid, host, and timeout.

### WebRTC Client
We need to implement the protocol stack of WebRTC.

Main steps of WebRTC client (browser) include
* Request IP and port of the remote server
* Connect to the remote server

Main steps of WebRTC client (go-server) include
* Login to signaling server
* Get IP and port for NAT punching from STUN server
* Periodically update IP and port to signaling server
* Update timeout to signalling server if changed

The rest thing is to design the framework of WebRTC protocol stack.

### Userful links
* https://code.google.com/p/webrtc/
* http://dev.w3.org/2011/webrtc/editor/webrtc.html

### Backup links
* https://gist.github.com/garyburd/1316852
* import "code.google.com/p/go.net/websocket"

