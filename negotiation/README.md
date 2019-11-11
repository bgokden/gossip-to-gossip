# Negotiation

Decentralised communication channel requires a pre-negotiation where both of the nodes don't trust each other but the communication is possible.

For this, I added a negotiation package where users try to validate each other before exchanging certificates and uplifting communication to the actual gossip channel.

#### Structures

For this nodes share some pre-shared knowledge with each other:
* A RSA Private Key
* A RSA Public Key
* A byte array prefix

This pre-shared data is called *PredefinedKey* in negotiation package.

Negotiation is handled with three phase handshake so a negotiation session is created.

*NegotiationSession* has two parameter in negotiation package:
* A PredefinedKey
* A byte array message


Assuming Node 1 and Node 2 has the same PredefinedKey.

- Node 1 creates a NegotiationSession with random test as session message.
- Node 1 encrypts and sends as message 1.
- Node 2 decrypts message 1, add prefix to the session message and send it back as message 2.
- Node 1 decrypts message 2 and compares the result to prefix prepended to the session message.

#### Security:
* All messages are encrypted with random data, so it is not possible to crack the communication with brute force attack.
* Each session has a different session message so communication is totally different.
* Request and response has different data inside so it is not possible to create, even if the node 2 is a fake server, node 2 can not try replaying the first message to another node.
