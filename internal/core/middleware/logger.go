package middleware

import (
    "fmt"
    "log/slog"
    "net/http"
    "time"

	"github.com/MaximShekshaev/kinotowerGo/internal/core/logger"
    "github.com/go-chi/chi/v5/middleware"
)

func RequestLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

        next.ServeHTTP(ww, r)

        status := ww.Status()
        if status == 0 {
            status = 404
        }
        duration := time.Since(start)

        logFn := statusLogLevel(status)
        logFn("http",
            "method", r.Method,
            "path", r.URL.Path,
            "status", fmt.Sprintf("%d %s", status, statusEmoji(status)),
            "ms", duration.Milliseconds(),
            "id", middleware.GetReqID(r.Context()),
        )
    })
}

func statusLogLevel(status int) func(string, ...any) {
    switch {
    case status >= 500:
        return logger.Log.Error
    case status >= 400:
        return logger.Log.Warn
    default:
        return logger.Log.Info
    }
}

func statusEmoji(status int) string {
    switch {
    case status >= 500:
        return "💥"
    case status >= 400:
        return "⚠️"
    case status >= 300:
        return "↩️"
    default:
        return "✓"
    }
}

var _ = slog.LevelInfo