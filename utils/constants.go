package utils

import "html/template"

const (
	VALID_HANDSHAKE_SECONDS = 15 // time interval within which to accept a handshake
)

const (
	DefaultRPCPort          string = "9521"           // time interval within which to accept a handshake
	DefaultWebSocketAddress string = "localhost:8088" // time interval within which to accept a handshake
)

const (
	RelayNodeType     uint = 0
	ValidatorNodeType      = 1
)

const MaxBlockSize = 1000

const (
	ValidMessageStore             string = "valid-messages"
	UnsentMessageStore                   = "unsent-messages"
	SentMessageStore                     = "sent-messages"
	NewChannelSubscriptionStore          = "new-channel-subscription"
	ChannelSubscriptionStore             = "channel-subscriptions"
	ChannelSubscriptionCountStore        = "channel-subscription-count"
	DeliveryProofStore                   = "delivery-proof-store"
	UnconfirmedDeliveryProofStore        = "unconfirmed-delivery-proof-store"
	DeliveryProofBlockStateStore         = "delivery-proof-block-state-store"
	SubscriptionBlockStateStore          = "subscription-block-state-store"
	DeliveryProofBlockStore              = "dp-block-store"
	SubscriptionBlockStore               = "sub-block-store"
)

// Channel Ids within main context
const (
	ConfigKey                       string = "Config"
	OutgoingMessageChId                    = "OutgoingMessageChannel"
	OutgoingMessageDP2PChId                = "OutgoingMessageDP2PChannel"
	IncomingMessageChId                    = "IncomingMessageChannel"
	PublishMessageChId                     = "PublishMessageChannel"
	SubscribeChId                          = "SubscribeChannel"
	SubscriptionDP2PChId                   = "SubscriptionDP2PChannel"
	ClientHandShackChId                    = "ClientHandshakeChannel"
	OutgoingDeliveryProof_BlockChId        = "OutgoingDeliveryProofBlockChannel"
	OutgoingDeliveryProofChId              = "OutgoinDeliveryProofChannel"
	PubSubBlockChId                        = "PubSubBlockChannel"
	PubsubDeliverProofChId                 = "PubsubProofChannel"
	PublishedSubChId                       = "PublishedSubChannel"
)

// State store key
const (
	CurrentDeliveryProofBlockStateKey string = "/df-block/current-state"
	CurrentSubscriptionBlockStateKey  string = "/sub-block/current-state"
)

type SubAction string

const (
	Broadcast SubAction = "broadcast"
	Join      SubAction = "join"
	Leave     SubAction = "leave"
)

var HomeTemplate = template.Must(template.New("").Parse(`
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<script>  
window.addEventListener("load", function(evt) {
    var output = document.getElementById("output");
    var input = document.getElementById("input");
    var ws;
    var print = function(message) {
        var d = document.createElement("div");
        d.textContent = message;
        output.appendChild(d);
        output.scroll(0, output.scrollHeight);
    };
    document.getElementById("open").onclick = function(evt) {
        if (ws) {
            return false;
        }
        ws = new WebSocket("{{.}}");
        ws.onopen = function(evt) {
            print("OPEN");
        }
        ws.onclose = function(evt) {
            print("CLOSE");
            ws = null;
        }
        ws.onmessage = function(evt) {
            print("RESPONSE: " + evt.data);
        }
        ws.onerror = function(evt) {
            print("ERROR: " + evt.data);
        }
        return false;
    };
    document.getElementById("send").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        print("SEND: " + input.value);
        ws.send(input.value);
        return false;
    };
    document.getElementById("close").onclick = function(evt) {
        if (!ws) {
            return false;
        }
        ws.close();
        return false;
    };
});
</script>
</head>
<body>
<table>
<tr><td valign="top" width="50%">
<p>Click "Open" to create a connection to the server, 
"Send" to send a message to the server and "Close" to close the connection. 
You can change the message and send multiple times.
<p>
<form>
<button id="open">Open</button>
<button id="close">Close</button>
<p><input id="input" type="text" value="Hello world!">
<button id="send">Send</button>
</form>
</td><td valign="top" width="50%">
<div id="output" style="max-height: 70vh;overflow-y: scroll;"></div>
</td></tr></table>
</body>
</html>
`))
