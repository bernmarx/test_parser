package main

import (
	"fmt"
	"github.com/adepte-myao/test_parser/internal/handlers"
	"github.com/adepte-myao/test_parser/internal/storage"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"os"

	"github.com/adepte-myao/test_parser/internal/config"
	"github.com/adepte-myao/test_parser/internal/server"
	"gopkg.in/yaml.v2"
)

func main() {
	f, err := os.Open("config/config.yaml")
	if err != nil {
		fmt.Println("[ERROR]: ", err)
		return
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		fmt.Println("[ERROR]: ", err)
		return
	}

	logger := logrus.New()
	router := mux.NewRouter()
	store := storage.NewStore(&cfg.Store, logger)

	server := server.NewServer(&cfg, logger, router)

	server.RegisterHandler("/link", handlers.NewLinksHandler(logger, cfg.Server.BaseLink, store).Handle)
	server.RegisterHandler("/solution", handlers.NewSolutionHandler(logger, cfg.Server.BaseLink, store).Handle)

	err = server.Start()
	if err != nil {
		fmt.Println("[ERROR]: ", err)
		return
	}
}
