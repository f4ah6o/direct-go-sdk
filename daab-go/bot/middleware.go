// Package bot provides a hubot-like bot framework for direct.
package bot

import (
	"context"
	"fmt"
	"log"
	"time"
)

// Middleware wraps a Handler to add cross-cutting behavior.
// Middleware are executed in reverse order (last added executes first).
type Middleware func(Handler) Handler

// Use adds a middleware to the bot's middleware chain.
// Middlewares are applied to message handlers in the order they are added.
// For example, if you call Use(m1).Use(m2), m1 will wrap m2,
// so m1 executes before m2 when a message is handled.
func (r *Robot) Use(m Middleware) {
	if r.middlewares == nil {
		r.middlewares = make([]Middleware, 0)
	}
	r.middlewares = append(r.middlewares, m)
}

// applyMiddlewares applies the middleware chain to a handler.
func (r *Robot) applyMiddlewares(h Handler) Handler {
	// Apply in reverse order so the first middleware wraps all others
	for i := len(r.middlewares) - 1; i >= 0; i-- {
		h = r.middlewares[i](h)
	}
	return h
}

// LoggingMiddleware logs handler execution with timing information.
func LoggingMiddleware(logger *log.Logger) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			start := time.Now()
			msg := res.Message

			logger.Printf("[START] Handler for message: %q from user: %s in room: %s",
				msg.Text, msg.UserID, msg.TalkID)

			defer func() {
				if p := recover(); p != nil {
					logger.Printf("[PANIC] Handler recovered: %v", p)
					panic(p) // Re-panic after logging
				}
				logger.Printf("[END] Handler completed in %v", time.Since(start))
			}()

			next(ctx, res)
		}
	}
}

// RecoveryMiddleware recovers from panics in handlers and logs them.
// This middleware prevents panics from crashing the bot.
func RecoveryMiddleware(logger *log.Logger) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			defer func() {
				if p := recover(); p != nil {
					logger.Printf("[RECOVER] Panic in handler: %v\nMessage: %q from: %s",
						p, res.Message.Text, res.Message.UserID)
				}
			}()
			next(ctx, res)
		}
	}
}

// FilterMiddleware allows selective message processing based on a filter function.
// If the filter returns false, the handler is not executed.
func FilterMiddleware(filter func(context.Context, Response) bool) Middleware {
	return func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			if !filter(ctx, res) {
				log.Printf("[FILTER] Message filtered: %q", res.Message.Text)
				return
			}
			next(ctx, res)
		}
	}
}

// RateLimitMiddleware limits how often a handler can execute for a given user.
// Uses a simple in-memory tracker with the given cooldown duration.
func RateLimitMiddleware(cooldown time.Duration) Middleware {
	// Track last execution time per user
	type userTracker struct {
		lastTime time.Time
	}
	trackers := make(map[string]*userTracker)

	return func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			userID := res.UserID()

			// Check if user is in cooldown
			if tracker, ok := trackers[userID]; ok {
				if time.Since(tracker.lastTime) < cooldown {
					log.Printf("[RATELIMIT] User %s is in cooldown, skipping", userID)
					return
				}
			}

			// Update tracker
			trackers[userID] = &userTracker{lastTime: time.Now()}

			next(ctx, res)
		}
	}
}

// ContextMiddleware adds values to the context before handler execution.
func ContextMiddleware(keyVal ...interface{}) Middleware {
	if len(keyVal)%2 != 0 {
		panic("ContextMiddleware requires an even number of arguments (key-value pairs)")
	}

	return func(next Handler) Handler {
		return func(ctx context.Context, res Response) {
			// Add all key-value pairs to context
			for i := 0; i < len(keyVal); i += 2 {
				key := fmt.Sprintf("%v", keyVal[i])
				val := keyVal[i+1]
				ctx = context.WithValue(ctx, contextKey(key), val)
			}
			next(ctx, res)
		}
	}
}

// contextKey is a custom type to avoid context key collisions.
type contextKey string

// GetContextValue retrieves a value from the context that was set by ContextMiddleware.
func GetContextValue(ctx context.Context, key string) interface{} {
	return ctx.Value(contextKey(key))
}

// Chain creates a middleware chain from multiple middlewares.
// The middlewares are composed in the order they are provided.
func Chain(middlewares ...Middleware) Middleware {
	return func(next Handler) Handler {
		// Apply in reverse order so first middleware wraps the rest
		for i := len(middlewares) - 1; i >= 0; i-- {
			next = middlewares[i](next)
		}
		return next
	}
}
