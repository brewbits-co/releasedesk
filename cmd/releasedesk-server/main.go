package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const WelcomeMessage = `
 ____      _                     ____            _    
|  _ \ ___| | ___  __ _ ___  ___|  _ \  ___  ___| | __
| |_) / _ \ |/ _ \/ _` + "`" + ` / __|/ _ \ | | |/ _ \/ __| |/ /
|  _ <  __/ |  __/ (_| \__ \  __/ |_| |  __/\__ \   <
|_| \_\___|_|\___|\__,_|___/\___|____/ \___||___/_|\_\


Welcome to the ReleaseDesk Server.
Documentation: https://releasedesk.dev/docs

Version:                v1.0.0
Console URL:            http://localhost:8080
Portal URL:             http://localhost:9090

`

func main() {
	fmt.Print(WelcomeMessage)

	// Container for dependency injection
	container := buildContainer()

	// Create the ReleaseDesk Console
	consoleServer := buildConsole(container)

	// Create the ReleaseDesk Portal
	portalServer := buildPortal(container)

	// Channel to signal shutdown
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	// Wait group to manage goroutines
	var wg sync.WaitGroup

	// Start Console Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := consoleServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting Console at port 8080: %v\n", err)
		}
	}()

	// Start Portal Server
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := portalServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error starting Portal at port 8080: %v\n", err)
		}
	}()

	// Wait for shutdown signal
	<-signalChan

	// Create a context with timeout for graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Shutdown server 1
	if err := consoleServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down Console: %v\n", err)
	} else {
		fmt.Println("Console shut down gracefully")
	}

	// Shutdown server 2
	if err := portalServer.Shutdown(ctx); err != nil {
		fmt.Printf("Error shutting down Portal: %v\n", err)
	} else {
		fmt.Println("Portal shut down gracefully")
	}

	// Wait for all goroutines to finish
	wg.Wait()
}
