package main
import (
	"fmt"
	"time"
)

func main() {
	// Debug := false
	// LogLevel := "info"
	// startUpTime := time.Now()
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)
}