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
		utils.Logger.Errorf("marshal json error::", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(jData)

}

// func (p *HttpService) serveJSONRPC(w http.ResponseWriter, req *http.Request) {
// 	// if req.Method != "CONNECT" {
// 	// 	http.Error(w, "method must be connect", 405)
// 	// 	return
// 	// }
// 	conn, _, err := w.(http.Hijacker).Hijack()
// 	if err != nil {
// 		http.Error(w, "internal server error", 500)
// 		return
// 	}
// 	defer conn.Close()
// 	io.WriteString(conn, "HTTP/1.0 Connected\r\n\r\n")
// 	jsonrpc.ServeConn(conn)
// }

func (p *HttpService) Start() error {

	host := fmt.Sprintf("%s:%s", p.Cfg.RPCHost, p.Cfg.RPCHttpPort)
	client, err := rpc.DialHTTP("tcp", host)

	if err != nil {
		utils.Logger.Errorf("Rpc Error:: %s", err.Error())
		return err
	}
	utils.Logger.Info("Dial to rpc successful!")
	p.rpcClient = client
	http.HandleFunc("/", p.sendHttp)
	// http.HandleFunc("/rpcendpoint", p.serveJSONRPC)
	err = http.ListenAndServe(":"+p.Cfg.RPCHttpPort, nil)
	return err
}
