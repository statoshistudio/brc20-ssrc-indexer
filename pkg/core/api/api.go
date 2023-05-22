package api

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	sql_mod "github.com/ByteGum/go-ssrc/pkg/core/sql"
	"github.com/gorilla/mux"
)

var db *sql.DB

func init() {

}

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/accounts", getAccounts)
	r.HandleFunc("/accounts-token-balances", getAccountTokenBalances)
	r.HandleFunc("/tokens", getTokens)
	r.HandleFunc("/tokens/{address}", getAccountTokens)
	r.HandleFunc("/inscriptions", getInscriptions)
	r.HandleFunc("/pending-transactions", getPendingTransactions)
	// http.Handle("/", r)

	// http.HandleFunc("/accounts", getAccounts)
	// http.HandleFunc("/tokens", getTokens)
	// http.HandleFunc("/inscriptions", getInscriptions)

	log.Fatal(http.ListenAndServe(":8088", r))

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
