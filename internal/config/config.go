package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GitHubToken string
	GitHubOwner string
	GitHubRepo  string
}

func Load() Config {
	err := godotenv.Load()

	if err != nil {
		log.Println("No .env file found")
	} else {
		log.Println(".env file loaded")
	}

	cfg := Config{
		GitHubToken: os.Getenv("GITHUB_TOKEN"),
		GitHubOwner: os.Getenv("GITHUB_OWNER"),
		GitHubRepo:  os.Getenv("GITHUB_REPO"),
	}

	if cfg.GitHubToken == "" {
		log.Fatal("GITHUB_TOKEN is required")
	}

	log.Printf("Owner: %s", cfg.GitHubOwner)
	log.Printf("Repo: %s", cfg.GitHubRepo)
	log.Printf("Token Present: %t", cfg.GitHubToken != "")

	return cfg
}
