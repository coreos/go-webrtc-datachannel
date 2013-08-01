go-webrtc-datachannel
=====================

This is a planning / brainstorming area for building services using webrtc-datachannel and golang. The end result is ideally a library that satisfies the following scenerios:

#### Allowing a javascript client to communicate directly with a NAT'd server

This could be done with regular AJAX style request, or with websockets. However, if we implemented this using the same protocols as webrtc-datachannels, we would be able to do NAT punching. This would be useful, for example, if a go-server is running on a virtual machine on a users laptop, but being controlled by javascript client in the browser. 

Ideal outcome: A hosted javascript client that can communicate with a go-server that is NAT'd, such as running in a Vagrant image on a users laptop.

For a proof of concept, supporting Chome or Firefox's (ideally both) datachannel implementation would be ideal. As of writing, Chrome and Firefox have different implementations. 

#### Allowing two NAT'd go servers to communicate directly

At this point we really just care about NAT punching using STUN/ICE. However, webrtc-datachannel wraps up which part of these should be supported, and google hosts a STUN server that we could use. 

Ideal outcome: Two go-servers, that can be running anywhere, but are able to establish an bi-directional communication channel. 

#### Other notes
* We do not care about the video/voice aspects of webrtc, only the datachannel
* This library should be as idomic Go as possible. 
* Should be relatively easy to make this work with existing go services, such as etcd

#### Additional reading

Some resources on the protocols:

* http://www.webrtc.org/firefox#TOC-DataChannels
* http://www.webrtc.org/chrome#TOC-Data-Channels-
* http://www.w3.org/TR/webrtc/#datachannel
* https://hacks.mozilla.org/2013/03/webrtc-data-channels-for-great-multiplayer/
* http://peerjs.com/
* https://github.com/denis-beurive/GoStun/blob/master/src/stun/doc.go
* http://www.webrtc.org/reference/architecture

Please contibute by sending pull requests for ideas or research, or better yet, CODE! :D


