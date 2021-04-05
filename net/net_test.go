package net

import (
	"os"
	"testing"
)

func TestGetIP(t *testing.T) {
	t.Run("get IP from default env", func(t *testing.T) {
		want := "123456"
		env := defaultOption().env
		os.Setenv(env, want)

		defer os.Setenv(env, "")

		got := GetIP()
		if got != want {
			t.Errorf("GetIP() = %v, want %v", got, want)
		}
	})

	t.Run("get IP from custom env", func(t *testing.T) {
		want := "12345678"
		env := "CATCHZENG_HOST"
		os.Setenv(env, want)
		defer os.Setenv(env, "")

		got := GetIP(WithEnv(env))
		if got != want {
			t.Errorf("GetIP() = %v, want %v", got, want)
		}
	})
}
