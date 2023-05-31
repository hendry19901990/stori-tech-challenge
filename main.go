package main

import (
	
    "fmt"
    "log"
    "os"
	"github.com/hendry19901990/stori-tech-challenge/repository"
	"github.com/hendry19901990/stori-tech-challenge/mappers"
	"github.com/hendry19901990/stori-tech-challenge/db"
)

func main()  {
	args := os.Args[1:] 
	if len(args) != 4 {
		fmt.Println("Usage: go run main.go <email> <txsFile> <senderEmail> <senderPassword>")
		return
	}

	fmt.Println("Args: ", args)

	userEmail := args[0]
	txsFile := args[1]

	senderEmail := args[2]
	password := args[3]
	

	auth := &repository.Auth{
		Email: senderEmail,
		Password: password,
	}

	records, err := repository.ReadCsvFile("/data/" + txsFile)
	if err != nil {
        log.Fatal(err)
    }


    summaryInfo := mappers.GetSummaryInfo(*records)
	summaryInfo.Email = userEmail
	fmt.Println(summaryInfo)

	infoBodyEmail := mappers.GetSummaryInfoBodyEmail(summaryInfo)
	fmt.Println(infoBodyEmail)

	r := repository.NewRequest([]string{userEmail}, "Transactions Summary", "", auth)
	if err := r.ParseTemplate("/templates/tx-summary.html", &infoBodyEmail); err != nil {
		log.Fatal(err)
	}

	//if err := r.SendEmail(); err != nil {
	//	log.Fatal(err)
	//}

	if err := db.SaveSummaryInfo(summaryInfo); err != nil {
		log.Fatal(err)
	}
}

