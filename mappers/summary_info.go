package mappers


import (
	
    "fmt"
	"strconv"
	"time"
	"strings"

)

type SummaryInfo struct {
	Email string
	Total float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	MonthTransactions map[string]int
}

type SummaryInfoBodyEmail struct{
	Email string
	Total float64
	AverageDebitAmount  float64
	AverageCreditAmount float64
	MonthTransactions string
}

func GetSummaryInfoBodyEmail(s SummaryInfo) SummaryInfoBodyEmail{
	info := SummaryInfoBodyEmail{
		Email: s.Email,
		Total: s.Total,
		AverageDebitAmount: s.AverageDebitAmount,
		AverageCreditAmount: s.AverageCreditAmount,
	}

	monthTransactions := ""
	for k, v := range s.MonthTransactions {
		monthTransactions += fmt.Sprintf("Number of transactions in %s: %d </br>", k, v)
	}

	info.MonthTransactions = monthTransactions
	return info
}

func GetSummaryInfo(records [][]string) SummaryInfo{
	summaryInfo := SummaryInfo{
		MonthTransactions: make(map[string]int),
	}

	var averageDebitAmountCount, averageCreditAmountCount int
	for i := 1; i < len(records); i++ {
		res := strings.Split(records[i][1], "/")
		if len(res) != 2 {
			continue
		}

		id := records[i][0]
		month, err := getMonth(res[0])
		if err != nil {
			fmt.Printf("ID: %s  error parsing month %e\n", id, err)
			continue
		}

        amount, err2 := getAmount(records[i][2])
		if err2 != nil {
			fmt.Printf("ID: %s  error parsing amount %e\n", id, err2)
			continue
		}

		if *amount > 0 {
			summaryInfo.AverageCreditAmount +=  *amount
			averageCreditAmountCount++
		}else {
			summaryInfo.AverageDebitAmount +=  *amount
			averageDebitAmountCount++
		}
    

		v, _ := summaryInfo.MonthTransactions[*month]
		summaryInfo.MonthTransactions[*month] = v + 1
		
		summaryInfo.Total += *amount
	}


	summaryInfo.AverageCreditAmount =  summaryInfo.AverageCreditAmount / float64(averageCreditAmountCount)
	summaryInfo.AverageDebitAmount =  summaryInfo.AverageDebitAmount / float64(averageCreditAmountCount)

	return summaryInfo
}

func getAmount(str string) (*float64, error){
	amount, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nil, err
	}
	return &amount, nil
}

func getMonth(str string) (*string, error){
	m, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	month := time.Month(m).String()
	
   return &month, nil
}
