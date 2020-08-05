package helpers

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/log"
	"time"
)

// LoggingMiddleware wraps the logs for incoming requests
type LoggingMiddleware struct {
	Logger log.Logger
	Next EncryptService
}
// Encrypt logs the encyption requests
func (mw LoggingMiddleware) Encrypt(ctx context.Context, key string, text string) (output string, err error) {
	fmt.Println("Log called")
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "encrypt",
			"key", key,
			"text", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	fmt.Println("Log called Encrypt start")
	output, err = mw.Next.Encrypt(ctx, key, text)
	fmt.Println("Log called Encrypt End")
	fmt.Println("Log Ends")
	return
}
// Decrypt logs the encyption requests
func (mw LoggingMiddleware) Decrypt(ctx context.Context, key string,
	text string) (output string, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "decrypt",
			"key", key,
			"message", text,
			"output", output,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output, err = mw.Next.Decrypt(ctx, key, text)
	return
}
