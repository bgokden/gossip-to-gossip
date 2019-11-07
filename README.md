# gossip-to-gossip

A Secure Gossip Protocol for GRPC

This is an extension project for veri

Veri is a decentralised and distributed Feature Store.

Node synchronisation and discovery is done over gossip.

I had some problems when I needed tls certificate rotation.

Gossip to gossip has two services running on different ports:

* Register (TLS)

* Gossip (Mutual TLS)

Register allows the first registration of another node.

A node can join in server mode or client mode.

There are different set of tokens to join as server mode or client mode.
Server and client tokens are unique per node and can be synced/queried from a node which is already in network.
To join the network, a new node should be nominated by another node.

When a server registered, the registering server will connect back and exchange other peer data.

Gossip allows synchronisation of new connection info (certificate rotation), service list and peer list.

Veri uses a weak network structure where nodes doesn't know and doesn't care about the full network as long as network share at least one node. This makes it virtually infinitely scalable.


It is expected that every node updates its own self signed certificate and send reachable interfaces public key to other peers. Every node is responsible for updating its client set with new certificates. It is possible that a node is running on multiple interfaces so every node can support multiple addresses. This is called broadcast addresses. Every join request includes this broadcast addresses and certificates information. It is also possible that a node doesn't know its external ip address, that is why every join response includes the extracted ip address of the node. So a node can ask other node about its external ip address. A node is responsible for updating its broadcast address list with new ip address information if possible.

Gossip to Gossip has concept of services and peers. It is very common that peers share a common service. As an example in kubernetes, all nodes can share a service and connect to service periodically. Although this service connection may be connecting itself. Over time, all nodes will be synchronised. This is a general node discovery system without using any api. This works across different network providers, container schedulers and cloud providers.

For known peers, the periodic join calls are run and other peer data is exchanged. Peers who haven't exchanged data for a long period are removed from peer list.

Since different grpc services can share same port and connection.
3rd party services can run on the server where gossip service is running and elevate its secure channels.


Connections will be stored in a connection map by client id to client wrapper.
Client wrapper keeps a grpc client and connection statistic like latency.

In big data applications, it is very important to have rack awareness but it is very hard to insert this data in our current work where nodes are running on virtual machines where we don't have enough information on actual network structure. Also there there are multiple layers rack, network, region, country ...

Instead of defining a hardcoded rack information, gossip-to-gossip use network latency as a pointer of distance.

By grouping peers by network latency, it is possible to predict cluster of peers.
Veri distributes queries to other peers in a way that queries reaches to most nodes more quickly. Most logical way is to query a peer from the most distant group first then a peer the least distant group then repeat this operation until all known peers are queried.

Gossip creates a virtual query list of clients with semi-optimum query order.

First peers grouped by average latency where groups are labeled by average latency of its elements. Where each group is implemented as stack where pop is implemented.

Groups are ordered by latency.
Let's assume there are n groups
Client order is created by
```
While groups have elements:
  let i = 0
    pop one client from group i if there is an element add to virtual client list
    pop one client from group (n-i) if there is an element and add it to virtual client list
    increment i
    until i reaches n/2
```

It is possible have a shorter virtual client list than number of available peers.

This virtual client list will be periodically updated that gives best possible query reach for veri distributed queries where reaching more varied nodes is important. This also allows automated client side load balancing where nodes are encouraged to query varied nodes.
