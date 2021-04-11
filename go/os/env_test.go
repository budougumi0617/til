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

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer setup(tt.want)()
			for k, want := range tt.want {
				if got := os.Getenv(k); got != want {
					t.Errorf("want %s is %s, but got %s", k, want, got)
				}
			}
		})
	}
}

func setup(envs map[string]string) func() {
	prevs := map[string]string{}
	for k, v := range envs {
		if prev, exist := os.LookupEnv(k); exist {
			// exist == trueの時は何かしら環境変数が存在しているので記憶しておく
			prevs[k] = prev
		}
		_ = os.Setenv(k, v)
	}
	return func() {
		// 引数でもらって設定した環境変数を軒並みunsetしていく
		for k := range envs {
			_ = os.Unsetenv(k)
			if v, ok := prevs[k]; ok {
				// 以前の設定があったとき
				_ = os.Setenv(k, v)
			}
		}
	}
}
