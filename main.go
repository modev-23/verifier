package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("domain, hasMx, hasSpf, spfRec, hasDmarc,dmarcREc\n")
	for scanner.Scan() {
		checkDomain(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Unable to read from input : %v\n", err)
	}
}
func checkDomain(domain string) {
	var hasMx, hasDmarc, hasSpf bool
	var spfRec, dmarcREc string

	mxRecords, err := net.LookupMX(domain)
	if err != nil {
		log.Printf("Errr : %v\n", err)
	}
	if len(mxRecords) > 0 {
		hasMx = true
	}

	txtRecords, err := net.LookupTXT(domain)
	if err != nil {
		log.Printf("Errr : %v\n", err)
	}

	for _, record := range txtRecords {
		if strings.HasPrefix(record, "v=spf1") {
			hasSpf = true
			spfRec = record
			break
		}
	}

	dmarcRecords, err := net.LookupTXT("_dmarc." + domain)
	if err != nil {
		log.Printf("Errr : %v\n", err)
	}
	for _, record := range dmarcRecords {
		if strings.HasPrefix(record, "v=DMARC1") {
			hasDmarc = true
			dmarcREc = record
			break
		}
	}

	fmt.Printf("Domain : %v\n has Mx : %v\n has SPF record : %v\n spf Record : %v\n has Dmarc : %v\n DMARC Record : %v\n", domain, hasMx, hasSpf, spfRec, hasDmarc, dmarcREc)
}
