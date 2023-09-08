package rpc

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/rpc"

	// "net/rpc/jsonrpc"

	"github.com/ByteGum/go-ssrc/utils"
)

type JsonRpc struct {
	JsonRpcVersion string            `json:"jsonrpc"`
	Id             int               `json:"id"`
	Method         string            `json:"method"`
	Params         []json.RawMessage `json:"params"`
}

type HttpService struct {
	Ctx       *context.Context
	Cfg       *utils.Configuration
	rpcClient *rpc.Client
}

func NewHttpService(mainCtx *context.Context) *HttpService {
	cfg, _ := (*mainCtx).Value(utils.ConfigKey).(*utils.Configuration)
	return &HttpService{
		Ctx: mainCtx,
		Cfg: cfg,
	}
}

func (p *HttpService) sendHttp(w http.ResponseWriter, r *http.Request) {

	var jsonrpc JsonRpc
	err := json.NewDecoder(r.Body).Decode(&jsonrpc)

	payload := jsonrpc.Params
	// var params interface{}
	var reply RpcResponse

	err = p.rpcClient.Call("RpcService."+jsonrpc.Method, payload[0], &reply)

	if err != nil {
		reply = RpcResponse{
			Data:   err,
			Status: "failure",
		}
	}

	jData, err := json.Marshal(reply)
	if err != nil {
		utils.Logger.Errorf("marshal json error:: %s", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

func (p *HttpService) Start() error {

	host := fmt.Sprintf("%s:%s", p.Cfg.RPCHost, p.Cfg.RPCPort)
	client, err := rpc.DialHTTP("tcp", host)

	if err != nil {
		utils.Logger.Errorf("Rpc Error:: %s", err.Error())
		return err
	}
	utils.Logger.Info("RPC Http client dial to rpc successful!")
	p.rpcClient = client
	http.HandleFunc("/", p.sendHttp)
	// http.HandleFunc("/rpcendpoint", p.serveJSONRPC)
	err = http.ListenAndServe(":"+p.Cfg.RPCHttpPort, nil)
	return err
}
