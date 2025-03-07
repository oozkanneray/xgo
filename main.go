package main

import (
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

func openBrowser(url string) error {
	var err error
	switch runtime.GOOS {
	case "windows":
		err = exec.Command("cmd", "/c", "start", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	return err
}

func main() {
	// Start the server in a goroutine
	go func() {
		server := ServeAPI(":3001")
		server.Run()
	}()

	// Wait a bit for the server to start
	time.Sleep(500 * time.Millisecond)

	// Open browser
	url := "http://localhost:3001"
	if err := openBrowser(url); err != nil {
		fmt.Printf("Error opening browser: %v\n", err)
	}

	// Keep the main goroutine running
	select {}
}
