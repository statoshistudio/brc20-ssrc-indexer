package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// "github.com/ByteGum/go-ssrc/pkg/core/indexer"
	"github.com/ByteGum/go-ssrc/pkg/core/indexer"
	sql_mod "github.com/ByteGum/go-ssrc/pkg/core/sql"
	"github.com/ByteGum/go-ssrc/utils"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var ctx context.Context
var cfg utils.Configuration

func init() {
	cfg = utils.Config
	ctx = context.Background()

	ctx = context.WithValue(ctx, utils.ConfigKey, &cfg)

}

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/accounts", getAccounts)
	r.HandleFunc("/accounts-token-balances", getAccountTokenBalances)
	r.HandleFunc("/tokens", getTokens)
	r.HandleFunc("/tokens/{address}", getAccountTokens)
	r.HandleFunc("/inscriptions", getInscriptions)
	r.HandleFunc("/generic-inscriptions", getGenericInscriptions)
	r.HandleFunc("/generic-inscriptions/{inscriptionId}", getUnitGenericInscription)
	r.HandleFunc("/pending-transactions", getPendingTransactions)
	r.HandleFunc("/callback", handleCallback)
	// http.Handle("/", r)

	// http.HandleFunc("/accounts", getAccounts)
	// http.HandleFunc("/tokens", getTokens)
	// http.HandleFunc("/inscriptions", getInscriptions)
	utils.Logger.Infof("Ordinal API listening on %s", cfg.OrdinalApiServer)
	err := http.ListenAndServe(cfg.OrdinalApiServer, r)
	if err != nil {
		// log.Fatal(err)
		panic(err)
	}

}

func perPageParams(current string, perPage string) (int, int, error) {
	_current, err := strconv.Atoi(current)
	if err != nil {
		return 0, 0, err
	}
	_perPage, err := strconv.Atoi(perPage)
	if err != nil {
		return 0, 0, err
	}
	return _current, _perPage, err
}

func getAccounts(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}
	result, err := sql_mod.GetAllAccounts(sql_mod.SqlDB, _current, _perPage)
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}
func getAccountTokenBalances(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}
	fmt.Println("--------")
	fmt.Println(r.URL.Query().Get("address"))
	result, err := sql_mod.GetAllAccountTokenBalances(sql_mod.SqlDB, _current, _perPage, r.URL.Query().Get("address"), r.URL.Query().Get("token"))
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}
func getTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}

	result, err := sql_mod.GetAllBrc20Tokens(sql_mod.SqlDB, _current, _perPage, "")
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func getAccountTokens(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)

	address := urlParams["address"]
	fmt.Println(urlParams)
	fmt.Println(address)
	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}

	result, err := sql_mod.GetAllBrc20Tokens(sql_mod.SqlDB, _current, _perPage, address)
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func getInscriptions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}

	result, err := sql_mod.GetAllInscriptions(sql_mod.SqlDB, _current, _perPage)
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func getGenericInscriptions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}

	result, err := sql_mod.GetAllGenericInscriptions(sql_mod.SqlDB, _current, _perPage)
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func getUnitGenericInscription(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	urlParams := mux.Vars(r)

	inscriptionId := urlParams["inscriptionId"]
	fmt.Println(urlParams)
	fmt.Println(inscriptionId)

	result, err := sql_mod.GetUnitGenericInscription(sql_mod.SqlDB, inscriptionId)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		message := make(map[string]string)
		message["message"] = err.Error()
		message["inscriptionId"] = inscriptionId
		fmt.Println("--------")
		fmt.Println(message)
		json.NewEncoder(w).Encode(message)
		return
	}
	fmt.Println("--------")
	fmt.Println(result.InscriptionBody)

	json.NewEncoder(w).Encode(indexer.GetOrdStructure(result))

}

func getPendingTransactions(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	message := make(map[string]string)
	_current, _perPage, err := perPageParams(r.URL.Query().Get("current"), r.URL.Query().Get("perPage"))
	if err != nil {
		message["message"] = err.Error()
		message["param"] = "current"
		json.NewEncoder(w).Encode(message)
		return
	}
	result, err := sql_mod.GetAllPendingTransactions(sql_mod.SqlDB, _current, _perPage)
	if err != nil {
		fmt.Println("--------")
		fmt.Println(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func handleCallback(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	message := make(map[string]string)
	query := r.URL.Query()
	txId := query.Get("txId")
	index := r.URL.Query().Get("index")
	offset := r.URL.Query().Get("offset")
	inscription_id := r.URL.Query().Get("inscription_id")
	values := r.URL.Query()
	for k, v := range values {
		fmt.Println(k, " => ", v)
	}

	if len(inscription_id) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		message["message"] = "inscription_id is required"
		json.NewEncoder(w).Encode(message)
		return
	}

	_, err := sql_mod.GetUnitGenericInscription(sql_mod.SqlDB, inscription_id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusOK)
			message["message"] = "New inscription"
			json.NewEncoder(w).Encode(message)
			return
		}
		w.WriteHeader(http.StatusBadGateway)
		message["message"] = err.Error()
		json.NewEncoder(w).Encode(message)
		return
	}
	_, err = sql_mod.SaveUpdatedInscription(sql_mod.SqlDB, inscription_id, fmt.Sprintf("%s:%s:%s", txId, index, offset))
	if err != nil {
		w.WriteHeader(http.StatusBadGateway)
		message["message"] = err.Error()
		json.NewEncoder(w).Encode(message)
		return
	}
	w.WriteHeader(http.StatusOK)
	message["message"] = inscription_id
	json.NewEncoder(w).Encode(message)
	return

	// inscription, err := indexer.GetUnitDataByIdFromServer(&ctx, inscription_id)
	// if err != nil {
	// 	w.WriteHeader(http.StatusInternalServerError)

	// 	message["message"] = err.Error()
	// 	message["inscriptionId"] = inscription_id
	// 	fmt.Println("--------")
	// 	fmt.Println(message)
	// 	json.NewEncoder(w).Encode(message)
	// 	return
	// }

	// go func() {
	// 	for i := 0; i < 5; i++ {
	// 		time.Sleep(4 * time.Second)
	// 		inscription, err = indexer.GetUnitDataByIdFromServer(&ctx, inscription_id)
	// 		if err != nil {
	// 			w.WriteHeader(http.StatusInternalServerError)

	// 			message["message"] = err.Error()
	// 			message["inscriptionId"] = inscription_id
	// 			fmt.Println("--------")
	// 			fmt.Println(message)
	// 			json.NewEncoder(w).Encode(message)
	// 			continue
	// 		}
	// 		_, err := indexer.HandleCallback(sql_mod.SqlDB, *inscription)
	// 		if err != nil {
	// 			w.WriteHeader(http.StatusNotFound)

	// 			message["message"] = err.Error()
	// 			message["inscriptionId"] = inscription_id
	// 			fmt.Println("--------")
	// 			fmt.Println(message)
	// 			json.NewEncoder(w).Encode(message)
	// 			continue
	// 		}
	// 		break
	// 	}

	// }()

	//json.NewEncoder(w).Encode(inscription)

}
