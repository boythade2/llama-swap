package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/mostlygeek/llama-swap/proxy"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	var (
		configPath  = flag.String("config", "config.yaml", "path to configuration file")
		listenAddr  = flag.String("listen", ":11434", "address to listen on")
		showVersion = flag.Bool("version", false, "show version information")
		logLevel    = flag.String("log-level", "info", "log level: info, warn, error")
	)
	flag.Parse()

	if *showVersion {
		fmt.Printf("llama-swap %s (commit: %s, built: %s)\n", version, commit, date)
		os.Exit(0)
	}

	// Load configuration
	cfg, err := proxy.LoadConfig(*configPath)
	if err != nil {
		log.Fatalf("failed to load config from %s: %v", *configPath, err)
	}

	// Apply CLI overrides
	if *listenAddr != ":11434" || cfg.ListenAddr == "" {
		cfg.ListenAddr = *listenAddr
	}

	log.Printf("llama-swap %s starting", version)
	log.Printf("config: %s", *configPath)
	log.Printf("listen: %s", cfg.ListenAddr)
	log.Printf("log-level: %s", *logLevel)

	// Create and start the proxy server
	server, err := proxy.New(cfg)
	if err != nil {
		log.Fatalf("failed to create proxy server: %v", err)
	}

	// Handle graceful shutdown
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigCh
		log.Printf("received signal %s, shutting down...", sig)

 = server

llama-swap stopped")
change listen port to 11434 to match ollama convention