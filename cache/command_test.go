package redisgo

import (
	"context"
	"testing"
)

func TestRedisgo_Set(t *testing.T) {
	r := NewRedisgo()
	t.Log(r.Set(context.Background(), "k1", "v1"))
}

func TestRedisgo_SetNEX(t *testing.T) {
	r := NewRedisgo(WithAddr("localhost:6379"))
	t.Log(r.SetNX(context.Background(), "k1", "v1"))
}

func TestRedisgo_Set2(t *testing.T) {
//	r := NewRedisgo(WithAddr("localhost:6379"))
}
func TestRedisgo_Get(t *testing.T) {
	r := NewRedisgo()
	t.Log(r.GetString(context.Background(), "k1"))
}
