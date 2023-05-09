package utils

// transmits messages received from other nodes in p2p to daemon
var IncomingMessagesP2P2_D_c = make(chan *ClientMessage)

// transmits messages sent through rpc, or other channels to daemon
var SentMessagesRPC_D_c = make(chan *ClientMessage)

// transmits validated messages from Daemon to P2P to be broadcasted
var OutgoingMessagesD_P2P_c = make(chan *ClientMessage)

// transmits new subscriptions from RPC to Daemon for processing
var SubscribersRPC_D_c = make(chan *Subscription)

// transmits valid subscriptions from Daemon to P2P for broadcasting
var SubscriptionD_P2P_C = make(chan *Subscription)
var ClientHandshakeC = make(chan *ClientHandshake)
var IncomingDeliveryProofsC = make(chan *DeliveryProof)
var OutgoingDeliveryProof_BlockC = make(chan *Block)
var OutgoingDeliveryProofC = make(chan *DeliveryProof)
var PubSubInputBlockC = make(chan *Block)
var PubSubInputProofC = make(chan *DeliveryProof)
var PublishedSubC = make(chan *Subscription)
