// package customerimporter reads from the given customers.csv file and returns a
// sorted (data structure of your choice) of email domains along with the number
// of customers with e-mail addresses for each domain.  Any errors should be
// logged (or handled). Performance matters (this is only ~3k lines, but *could*
// be 1m lines or run on a small machine).
package customerimporter

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

func PrintCustomers() []domainCount {
	file, err := os.Open(getFileToOpen())

	if err != nil {
		log.Println(err)
	}

	return countDomains(file)
}

func getFileToOpen() string {
	_, currentFile, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal("Couldn't retrieve file, program terminated")
	}
	return filepath.Join(filepath.Dir(currentFile), "customers.csv")
}

func countDomains(f *os.File) []domainCount {
	parser := csv.NewReader(f)
	domainMap := make(map[string]int)

	for {
		record, err := parser.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		if record[2] == "email" {
			continue
			// I don't know any better way of getting first line for now :)
		}

		fillDomainMap(record[2], domainMap)

	}
	return castDomainToSortedStruct(domainMap)
}

func fillDomainMap(r string, dm map[string]int) {
	splitted := strings.Split(r, "@")
	d := splitted[1]
	v, ok := dm[splitted[1]]
	if ok {
		dm[d] = v + 1
	} else {
		dm[d] = 1
	}
}

func castDomainToSortedStruct(dm map[string]int) []domainCount {
	var sorteDomains []domainCount

	for k, v := range dm {
		sorteDomains = append(sorteDomains, domainCount{k, v})
	}

	sort.Slice(sorteDomains, func(i, j int) bool {
		return sorteDomains[i].count > sorteDomains[j].count
	})

	return sorteDomains

}
