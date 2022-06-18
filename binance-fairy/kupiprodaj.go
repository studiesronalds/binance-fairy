package main

import (
	"./binance"
    "fmt"
)

func kupiExactSymbol (symbol string, amount float64, price float64) {
    order := binance.LimitOrder {
        Symbol:      symbol,
        Side:        "BUY",
        Type:        "LIMIT",
        TimeInForce: "GTC",
        Quantity:    amount,
        Price:       price,
    }

    client := binanceClient();
    res, err := client.PlaceLimitOrder(order)
    
    if err != nil {
    	fmt.Println(err)
    }
    
    fmt.Println(res)
}

func kupiSymbol (symbol string, amount float64) {
	client := binanceClient();

	order := binance.MarketOrder {
        Symbol:   symbol,
        Side:     "BUY",
        Type:     "MARKET",
        Quantity: amount,
    }

    res, err := client.PlaceMarketOrder(order)
    
    if err != nil {
    	fmt.Println(err)
    }

    fmt.Println(res)	

}

func kupi (amount float64 ) {
	kupiSymbol("MATICBTC", amount);
}

func proday (amount float64) {
	kupiSymbol("BTCMATIC", amount);
}