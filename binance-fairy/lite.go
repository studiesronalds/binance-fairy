package main

import (
    "os"
    "fmt"
    "log"
    "reflect"
    "./binance"
    "github.com/gorilla/mux"
    "encoding/json"
	"net/http"
)

var client *binance.Binance;

type FairyRequest struct {
	Symbol string `json:"symbol"`
	Response float64 `json:"response"`
}

func binanceClient() *binance.Binance {
	return client
}

func fairyApi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	var fairyRequest FairyRequest
	_ = json.NewDecoder(r.Body).Decode(&fairyRequest)

    query := binance.SymbolQuery {
        Symbol: fairyRequest.Symbol,
    }

    res, err := client.GetLastPrice(query)

    if err != nil {
        panic(err)
    }

    fmt.Println(res)
    fmt.Println(reflect.TypeOf(res.Price));
    fairyRequest.Response = res.Price

	json.NewEncoder(w).Encode(&fairyRequest)
}

type Trigger struct { 
    PositiveSentiment bool `json:"positiveSentiment"`
    TweetTime string ` json:"tweetTime"`
    Coins []string `json:"coins"`
    IntervalMinutes int `json:"intervalMinutes"`
}

type ImpactRequest struct {
    TriggerOutput []Trigger `json:"triggers"`
}

func impactApi(w http.ResponseWriter, r *http.Request){
    w.Header().Set("Content-Type", "application/json")

    var impactRequest ImpactRequest
    _ = json.NewDecoder(r.Body).Decode(&impactRequest)

    fmt.Println("IMPACT ::::: ", impactRequest)
    gatherHistory(client, impactRequest);

    json.NewEncoder(w).Encode(&impactRequest)
}

func main() {

    client = binance.New(os.Getenv("BINANCE_API_KEY"), os.Getenv("BINANCE_SECRET"))
    positions, err := client.GetPositions()

    if err != nil {
        panic(err)
    }

    for _, p := range positions {
        fmt.Println(p.Asset, p.Free, p.Locked)
    }
    
    r := mux.NewRouter()
    r.HandleFunc("/fairy", fairyApi).Methods("POST")
    r.HandleFunc("/api/trigger/check", impactApi).Methods("POST")
    
	port := "8886"

	if os.Getenv("DEVELOPMENT_PORT_BINANCE") != "" {
		port = os.Getenv("DEVELOPMENT_PORT_BINANCE")
	}

	fmt.Println(fmt.Sprintf("Starting Server 0.0.0.0:%s", port))
 	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), r))
}