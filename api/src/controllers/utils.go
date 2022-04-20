package controllers

import (
	"context"
	"time"
)

func InitContext() (context.Context, context.CancelFunc) {
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    return ctx, cancel
}
