package logs

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

// var logFile untuk menyimpan file log
var logFile *os.File

// InitLogger untuk inisialisasi log file
func InitLogger() {
	logFileName := fmt.Sprintf("logs/log-%s.log", time.Now().Format("2006-01-02"))
	var err error
	logFile, err = os.OpenFile(logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
}

// LogInfo untuk mencatat log informasi
func LogInfo(r *http.Request, message string, statusCode int, responseTime time.Duration) {
	if logFile == nil {
		InitLogger()
	}

	ip := r.RemoteAddr
	agent := r.UserAgent()
	endpoint := r.URL.Path
	method := r.Method // Menambahkan method HTTP

	logEntry := fmt.Sprintf(
		"[%s] INFO: Method=%s, IP=%s, User-Agent=%s, Endpoint=%s, Status=%d, ResponseTime=%s, Message=%s\n",
		time.Now().Format("2006-01-02 15:04:05"), method, ip, agent, endpoint, statusCode, responseTime, message,
	)

	// Tulis log ke file
	_, err := logFile.WriteString(logEntry)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}

	// Tampilkan ke terminal juga
	fmt.Print(logEntry)
}

// LogError untuk mencatat log error
func LogError(r *http.Request, message string, statusCode int, responseTime time.Duration) {
	if logFile == nil {
		InitLogger()
	}

	ip := r.RemoteAddr
	agent := r.UserAgent()
	endpoint := r.URL.Path
	method := r.Method // Menambahkan method HTTP

	logEntry := fmt.Sprintf(
		"[%s] ERROR: Method=%s, IP=%s, User-Agent=%s, Endpoint=%s, Status=%d, ResponseTime=%s, Error=%s\n",
		time.Now().Format("2006-01-02 15:04:05"), method, ip, agent, endpoint, statusCode, responseTime, message,
	)

	// Tulis log ke file
	_, err := logFile.WriteString(logEntry)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}

	// Tampilkan ke terminal juga
	fmt.Print(logEntry)
}

// LogInfoSimple untuk mencatat log informasi tanpa memerlukan *http.Request
func LogInfoSimple(message string) {
	if logFile == nil {
		InitLogger()
	}

	logEntry := fmt.Sprintf(
		"[%s] INFO: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), message,
	)

	// Tulis log ke file
	_, err := logFile.WriteString(logEntry)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}

	// Tampilkan ke terminal juga
	fmt.Print(logEntry)
}

// LogErrorSimple untuk mencatat log error tanpa memerlukan *http.Request
func LogErrorSimple(message string) {
	if logFile == nil {
		InitLogger()
	}

	logEntry := fmt.Sprintf(
		"[%s] ERROR: %s\n",
		time.Now().Format("2006-01-02 15:04:05"), message,
	)

	// Tulis log ke file
	_, err := logFile.WriteString(logEntry)
	if err != nil {
		fmt.Println("Error writing to log file:", err)
	}

	// Tampilkan ke terminal juga
	fmt.Print(logEntry)
}

// CloseLogger untuk menutup file log
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
	}
}
