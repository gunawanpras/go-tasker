package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type (
	Server struct {
		Port string `yaml:"port"`
	}

	BackProductionDb struct {
		Interval int `yaml:"interval"`
	}

	CleanUp struct {
		Interval int `yaml:"interval"`
	}

	CacheEviction struct {
		Interval int `yaml:"interval"`
	}

	GenerateReport struct {
		Interval int `yaml:"interval"`
	}

	TaskScheduler struct {
		BackProductionDb BackProductionDb `yaml:"backup_production_db"`
		CleanUp          CleanUp          `yaml:"clean_up"`
		CacheEviction    CacheEviction    `yaml:"cache_eviction"`
		GenerateReport   GenerateReport   `yaml:"generate_report"`
	}

	Config struct {
		Server        Server        `yaml:"server"`
		TaskScheduler TaskScheduler `yaml:"task_scheduler"`
	}
)

func LoadConfig(filePath string) Config {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Error opening config file:", err)
	}
	defer file.Close()

	var config Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("Error decoding config file:", err)
	}

	return config
}
