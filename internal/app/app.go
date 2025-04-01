package app

import (
	"bwanews/config"

	"github.com/rs/zerolog/log"
)

func RunServer() {
	cfg := config.NewConfig() 
	_, err := cfg.ConnectionPostgress()

	if err != nil {
		log.Fatal().Msgf("Failed connecting to database : %v", err)
		return 
	}
}