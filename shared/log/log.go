package log

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"runtime/debug"
	"time"
)

var (
	logger  *log.Logger
	LogName string
)

//InitializeLogger ...
func InitializeLogger(logPath string) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Execption InitializeLogger err: ", err)
		}
	}()

	createLogfileObj(logPath)
}

type basicWriter struct {
	http.ResponseWriter
	WroteHeader bool
	Code        int
	Bytes       int
}

func wrapResponseWriter(w http.ResponseWriter) *basicWriter {
	bw := basicWriter{ResponseWriter: w, WroteHeader: false, Code: 0, Bytes: 0}
	return &bw
}

func (b *basicWriter) WriteHeader(code int) {
	if !b.WroteHeader {
		b.Code = code
		b.WroteHeader = true
		b.ResponseWriter.WriteHeader(code)
	}
}

func (b *basicWriter) Write(buf []byte) (int, error) {
	//b.WriteHeader(http.StatusOK)
	n, err := b.ResponseWriter.Write(buf)
	b.Bytes += n
	return n, err
}

func (b *basicWriter) Status() int {
	return b.Code
}

func (b *basicWriter) IsSuccess() bool {
	return b.Status() >= 200 && 200 <= b.Status()
}

func (b *basicWriter) BytesWritten() int {
	return b.Bytes
}

// Logger ...
func Logger(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {

		lw := wrapResponseWriter(w)
		start := time.Now()

		defer func() {
			if lw.Status() == 0 {
				lw.WriteHeader(http.StatusNotFound)
			}

			scheme := "http"
			if r.TLS != nil {
				scheme = "https"
			}

			result := fmt.Sprintf("%v  %v://%v%v %v  from %v - %v %v in %v", r.Method, scheme, r.Host, r.RequestURI, r.Proto, r.RemoteAddr, lw.Status(), lw.BytesWritten(), time.Now().Sub(start))

			if lw.IsSuccess() {
				Info(r.Context(), result)
			} else {
				Error(r.Context(), result)
			}

			recover := recover()
			if recover != nil {
				Error(r.Context(), recover, string(debug.Stack()))
			}

		}()

		h.ServeHTTP(lw, r)

	}

	return http.HandlerFunc(fn)
}

func createLogfileObj(dir string) {

	fileName := dir + "/" + LogName + ".log"
	fmt.Println("Log FileName : ", fileName)

	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("Unable to Create Log File: ", err)
		return
	}

	logger = log.New(logFile, "", log.Ldate|log.Ltime)
}

// Fatal ...
func Fatal(ctx context.Context, logMSG ...interface{}) {
	logger.Fatal(logMSG...)
}

// Println log based on Severty level
func Println(ctx context.Context, logMSG ...interface{}) {
	logger.Println(logMSG...)
}

//Debug ...
func Debug(ctx context.Context, logMSG ...interface{}) {
	logger.Println(logMSG...)
}

// Error log based on Severty level
func Error(ctx context.Context, logMSG ...interface{}) {
	logger.Println(logMSG...)
}

// Info log based on Severty level
func Info(ctx context.Context, logMSG ...interface{}) {
	logger.Println(logMSG...)
}

//Warn ...
func Warn(ctx context.Context, logMSG ...interface{}) {
	logger.Println(logMSG...)
}
