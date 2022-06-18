package main
 
import (
	"./binance"
    "fmt"
    // "time"
    "strings"
    "strconv"
    "os/exec"
);

func Explode(delimiter, text string) []string {
	if len(delimiter) > len(text) {
		return strings.Split(delimiter, text)
	} else {
		return strings.Split(text, delimiter)
	}
}

func gatherHistory(client *binance.Binance, impact ImpactRequest){


	/** 
	1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M
	**/

	for _, trigger := range impact.TriggerOutput {
		// sentiment trigger.PositiveSentiment
		// 2020-04-29 00:37:36 trigger.TweetTime
		// dateArray  := Explode(" ", trigger.TweetTime)
		// dateBig := Explode("-", dateArray[0])
		// dateSmall  := Explode(":", dateArray[1])

		// year  , _ := strconv.ParseInt(dateBig[0], 6, 12) 
		// month  , _ := strconv.ParseInt(dateBig[1], 6, 12)
		// day  , _ := strconv.ParseInt(dateBig[2], 6, 12) 

		// hour  , _ := strconv.ParseInt(dateSmall[0], 6, 12)
		// minute  , _ := strconv.ParseInt(dateSmall[1], 6, 12)
		// second  , _ := strconv.ParseInt(dateSmall[2], 6, 12)


		// theTime := time.Date(int(year), time.Month(int(month)), int(day), int(hour), int(minute), int(second), 0, time.UTC)
		// fmt.Println(int(year), time.Month(int(month)), int(day), int(hour), int(minute), int(second), 0)
		// fmt.Println("%s", theTime.Local())


		app := "date"
	    arg0 := "-d"
	    arg1 := trigger.TweetTime
	    arg2 := "+\"%s\""

	    cmd := exec.Command(app, arg0, arg1, arg2)
	    stdout, err := cmd.Output()

	    if err != nil {
	        fmt.Println(err.Error())
	        return
	    }
	    timestamp := strings.Replace(string(stdout), "\n", "", -1) + "000"
	    timestamp = strings.Replace(timestamp, "\"", "", -1)
	    // Print the output

	    query := binance.HistoryQuery {
	        Symbol: "BTCUSDT", 
	        OpenTime: timestamp,
	        CloseTime: timestamp,
	        WindowSize: "12h",
	    }

	    klines, err := client.GetHistory(query)  

	    if err != nil {
	    	fmt.Println(err)
			return
		}
		
		//fmt.Println(res)
		/**
			type Kline struct {
				OpenTime         int64
				Open             float64
				High             float64
				Low              float64
				Close            float64
				Volume           float64
				CloseTime        int64
				QuoteVolume      float64
				NumTrades        int64
				TakerBaseVolume  float64
				TakerQuoteVolume float64
			}
		**/
		fmt.Println("REQUEST !!!!")

		startValue := klines[0].Low;
		endValue := klines[len(klines) - 1].High;
		totalTrades := 0

		for _, kline := range klines { 
			// fmt.Println(kline.OpenTime, "OpenTime")
			// fmt.Println(kline.High, "High")
			// fmt.Println(kline.Low, "Low")
			// fmt.Println(kline.Close, "Close")
			// fmt.Println(kline.Volume, "Volume")
			// fmt.Println(kline.CloseTime, "CloseTime")
			// fmt.Println(kline.QuoteVolume, "QuoteVolume")
			// fmt.Println(kline.NumTrades, "NumTrades")
			// fmt.Println(kline.TakerBaseVolume, "TakerBaseVolume")
			// fmt.Println(kline.TakerQuoteVolume, "TakerQuoteVolume")
			totalTrades = int(kline.NumTrades) + totalTrades
		}
		/**
		    PositiveSentiment bool `json:"positiveSentiment"`
		    TweetTime string ` json:"tweetTime"`
		    Coins []string `json:"coins"`
		    IntervalMinutes int `json:"intervalMinutes"`  
		**/
		diff := endValue - startValue
		fmt.Println(fmt.Sprintf("Summary for IfElse Algorithm ::: startValue=%s&endValue=%s&totalTrades=%s&diff=%s", strconv.FormatFloat(startValue, 'f', 5, 64), strconv.FormatFloat(endValue, 'f', 5, 64), strconv.Itoa(totalTrades)), strconv.FormatFloat(diff, 'f', 5, 64))

		//decrease - increase

	}

	fmt.Println("History run")

}