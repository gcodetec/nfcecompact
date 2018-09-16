package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gcodetec/nfcecompact/nfcecompact"
)

func main() {
	fmt.Println("listing files")
	now := time.Now().Local()
	currentYear := now.Year()
	currentMonth := int(now.Month())
	path := flag.String("path", "./data", "o diretorio dos arquivos")
	competenceYear := flag.Int("year", currentYear, "o diretorio dos arquivos")
	competenceMonth := flag.Int("month", currentMonth, "o diretorio dos arquivos")
	flag.Parse()

	nfcecompact.CompactFilesByCompetence(*path, *competenceYear, *competenceMonth)
}
