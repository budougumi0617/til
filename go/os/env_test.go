package os

import (
	"os"
	"testing"
)

func Test_getConfigData(t *testing.T) {
	tests := []struct {
		name string
		want map[string]string
	}{
		{
			name: "normalTest",
			want: map[string]string{
				"SLACK_API_URL":        "send_api_url",
				"SLACK_CHANNEL_NAME":   "channel_name",
				"AWS_ACCESS_KEY":       "access_key",
				"AWS_SECRET_KEY":       "secret_key",
				"WBEW_DATABASE_REGION": "region",
				"WBEW_DATABASE_NAME":   "data_base",
			},
		},
	}

	oldSlackAPIURL := "old"
	_ = os.Setenv("SLACK_API_URL", oldSlackAPIURL)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			setup(t, tt.want)
			for k, want := range tt.want {
				if got := os.Getenv(k); got != want {
					t.Errorf("want %s is %s, but got %s", k, want, got)
				}
			}
		})
	}
	t.Log("check!!")
	if got := os.Getenv("SLACK_API_URL"); got != oldSlackAPIURL {
		t.Fatalf("no restore, SLACK_API_URL is: %q", got)
	}
}

func setup(t *testing.T, envs map[string]string) {
	for k, v := range envs {
		prev, ok := os.LookupEnv(k)
		if err := os.Setenv(k, v); err != nil {
			t.Fatalf("cannot set environment key: %q", k)
		}
		k := k // 束縛しておく
		if ok {
			t.Cleanup(func() {
				_ = os.Setenv(k, prev)
				got := os.Getenv(k)
				t.Logf("update %q:%q", k, got)
			})
		} else {
			t.Cleanup(func() { _ = os.Unsetenv(k) })
		}
	}
}
