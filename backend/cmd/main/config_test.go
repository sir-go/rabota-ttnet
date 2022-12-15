package main

import (
	"reflect"
	"testing"
	"time"

	"rabota_ttnet/backend/internal/server"
	"rabota_ttnet/backend/internal/store"
)

func Test_unmarshal(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *Config
		wantErr bool
	}{
		{"empty", args{nil}, nil, true},
		{"ok", args{[]byte(`
[service]
    host        = "localhost"
    port        = 4563
    secret      = "-uuid-"
    run_mode    = "release"
    [service.timeouts]
        write   = "5m"
        read    = "8m42s"
        idle    = "10m"
[db]
    host        = "mysql-server"
    port        = 3306
    user        = "admin-user"
    password    = "admin-password"
    dbname      = "db-name"
    timeout     = "10s"
    id_len      = 10
    id_tries    = 50
`)}, &Config{
			Service: server.Config{
				Host:   "localhost",
				Port:   4563,
				Secret: "-uuid-",
				Timeouts: struct {
					Write time.Duration `toml:"write"`
					Read  time.Duration `toml:"read"`
					Idle  time.Duration `toml:"idle"`
				}{
					5 * time.Minute,
					8*time.Minute + 42*time.Second,
					10 * time.Minute,
				},
				RunMode: "release",
			},
			Db: store.Config{
				Host:     "mysql-server",
				Port:     3306,
				User:     "admin-user",
				Password: "admin-password",
				DbName:   "db-name",
				Timeout:  10 * time.Second,
				IdLen:    10,
				IdTries:  50,
			},
		}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := unmarshal(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("unmarshal() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("unmarshal() gotCfg = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
