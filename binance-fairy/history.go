package main
 
import (
	"./binance"
    "fmt"
    // "time"
    "math"
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


type IfElseMagic struct {
	Predicted bool `json:"predicted"`
	MagicPredictedPercentage float64 `json:"percentage"`
	Klines []binance.Kline `json:"data"`
	Coin string `json:"coin"`
}

func gatherHistory(client *binance.Binance, impact ImpactRequest){

	var elseIfMagic []IfElseMagic

	for _, trigger := range impact.TriggerOutput {

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

		/**
		    IntervalMinutes int `json:"intervalMinutes"`
		    1m 3m 5m 15m 30m 1h 2h 4h 6h 8h 12h 1d 3d 1w 1M  
		**/
		WindowSize := "4h"
		if trigger.IntervalMinutes > 1 {
			WindowSize = "3m"
		} else if trigger.IntervalMinutes > 3 && trigger.IntervalMinutes < 5 {
			WindowSize = "5m"
		} else if trigger.IntervalMinutes > 5 && trigger.IntervalMinutes < 15 {
			WindowSize = "15m"
		} else if trigger.IntervalMinutes > 15 && trigger.IntervalMinutes < 30 {
			WindowSize = "30m"
		} else if trigger.IntervalMinutes > 30 && trigger.IntervalMinutes < 60 {
			WindowSize = "1h"
		} else if trigger.IntervalMinutes > 60 && trigger.IntervalMinutes < 120 {
			WindowSize = "2h"
		} else if trigger.IntervalMinutes > 120 && trigger.IntervalMinutes < 240 {
			WindowSize = "4h"
		} else if trigger.IntervalMinutes > 240 && trigger.IntervalMinutes < 360 {
			WindowSize = "6h"
		} else if trigger.IntervalMinutes > 360 && trigger.IntervalMinutes < 480 {
			WindowSize = "8h"
		} else if trigger.IntervalMinutes > 480 && trigger.IntervalMinutes < 720 {
			WindowSize = "12h"
		} else if trigger.IntervalMinutes > 720 && trigger.IntervalMinutes < 1440 {
			WindowSize = "1d"
		} else if trigger.IntervalMinutes > 1440 && trigger.IntervalMinutes < 4320 {
			WindowSize = "3d"
		} else if trigger.IntervalMinutes > 4320 {
			WindowSize = "1w"
		}


	    for _, coin := range trigger.Coins {
		    query := binance.HistoryQuery {
		        Symbol: fmt.Sprintf("%sUSDT", coin), 
		        OpenTime: timestamp,
		        CloseTime: timestamp,
		        WindowSize: WindowSize,
		    }

		    klines, err := client.GetHistory(query)  

		    if err != nil {
		    	fmt.Println(err)
				return
			}
			
			fmt.Println("REQUEST !!!!")

			startValue := klines[0].Low;
			endValue := klines[len(klines) - 1].High;
			totalTrades := 0

			for _, kline := range klines { 
				totalTrades = int(kline.NumTrades) + totalTrades
			}

			diff := endValue - startValue
			fmt.Println(fmt.Sprintf("Summary for IfElse Algorithm ::: startValue=%s&endValue=%s&totalTrades=%s&diff=%s", strconv.FormatFloat(startValue, 'f', 5, 64), strconv.FormatFloat(endValue, 'f', 5, 64), strconv.Itoa(totalTrades)), strconv.FormatFloat(diff, 'f', 5, 64))

			var magic IfElseMagic
			magic.Predicted = true
			magic.Coin = coin
			//decrease - increase
			if (trigger.PositiveSentiment && diff < 0) || (!trigger.PositiveSentiment && diff > 0)  {
				magic.Predicted = false
			}

			if magic.Predicted && diff < 0 {
				magic.MagicPredictedPercentage = ((math.Abs(diff)/startValue) * 100)
			} else if magic.Predicted && diff > 0 {
				magic.MagicPredictedPercentage = ((math.Abs(diff)/endValue) * 100)
			}

			elseIfMagic = append(elseIfMagic, magic)
	    }
	}

	fmt.Println("Final Verdict")
	fmt.Println(elseIfMagic)

	// simple artificial If - if more then half are !predicted , then your strategy just sucks 
	// count of failed/success versions 
	verdict := false
	totalImpactCandidates := len(elseIfMagic)
	impactedCount := 0
	// each !predicted makes percent 0
	totalPercent := 0
	// average of all predicted
	averagePercent := 0
	// get the best of all ....
	bestPredictionPercent := 0
	var bestPredictionKlinkes []binance.Kline
	var bestCoin string

	for _, verdictCollect := range elseIfMagic { 
		if verdictCollect.Predicted {
			impactedCount = impactedCount + 1
			totalPercent = int(verdictCollect.MagicPredictedPercentage)
			if bestPredictionPercent < int(verdictCollect.MagicPredictedPercentage) {
				bestPredictionPercent = int(verdictCollect.MagicPredictedPercentage)
				bestPredictionKlinkes = verdictCollect.Klines
				bestCoin = verdictCollect.Coin
			}
		}
	}

	averagePercent = totalPercent / totalImpactCandidates;
	if averagePercent > 60 &&  impactedCount > (totalImpactCandidates/2) {
		verdict = true
	}

	var verdictResponse IfElseMagic
	verdictResponse.Klines = bestPredictionKlinkes
	verdictResponse.Predicted = verdict
	verdictResponse.MagicPredictedPercentage = float64(averagePercent)
	verdictResponse.Coin = bestCoin

	fmt.Println("Final Verdict SummUp")
	fmt.Println(verdictResponse)
}

