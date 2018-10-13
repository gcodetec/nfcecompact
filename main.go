package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/gcodetec/nfcecompact/nfcecompact"
	"github.com/spf13/viper"
)

func main() {
	fmt.Println("listing files")
	now := time.Now().Local()
	currentYear := now.Year()
	currentMonth := int(now.Month())

	configFile := flag.String("f", "config", "seta o arquivo de configuração")
	competenceYear := flag.Int("year", currentYear, "o diretorio dos arquivos")
	competenceMonth := flag.Int("month", currentMonth, "o diretorio dos arquivos")
	flag.Parse()

	viper.SetConfigName(*configFile)
	viper.AddConfigPath(".")

	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	path := viper.GetString("path")
	copyPath := viper.GetString("copy_path")
	nfcecompact.CopyAllFilesToPath(copyPath, path)
	nfcecompact.CompactFilesByCompetence(path, *competenceYear, *competenceMonth)
}
