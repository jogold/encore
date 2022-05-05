package sqldb

import (
	"testing"
	_ "unsafe" // for go:linkname

	"encore.dev/runtime/config"
)

func TestDBConf(t *testing.T) {
	tests := []struct {
		DB       *config.SQLDatabase
		Host     string
		Port     uint16
		MaxConns uint32
	}{
		{
			DB: &config.SQLDatabase{
				EncoreName:     "ignore",
				DatabaseName:   "dbname",
				Host:           "/cloudsql/foo",
				User:           "user",
				Password:       "password",
				MaxConnections: 10,
			},
			Host:     "/cloudsql/foo",
			Port:     5432,
			MaxConns: 10,
		},
		{
			DB: &config.SQLDatabase{
				EncoreName:     "ignore",
				DatabaseName:   "dbname",
				Host:           "test:123",
				User:           "user",
				Password:       "password",
				MaxConnections: 0,
			},
			Host:     "test",
			Port:     123,
			MaxConns: 30,
		},
		{
			DB: &config.SQLDatabase{
				EncoreName:     "ignore",
				DatabaseName:   "dbname",
				Host:           "hostname",
				User:           "user",
				Password:       "password",
				MaxConnections: 100,
			},
			Host:     "hostname",
			Port:     5432,
			MaxConns: 100,
		},
	}

	for i, test := range tests {
		cfg, err := dbConf(test.DB)
		if err != nil {
			t.Fatalf("test %d: unexpected error: %v", i, err)
		}

		if cfg.ConnConfig.Host != test.Host {
			t.Fatalf("test %d: got host %s, want %q", i, cfg.ConnConfig.Host, test.Host)
		} else if cfg.ConnConfig.Port != test.Port {
			t.Fatalf("test %d: got port %d, want %d", i, cfg.ConnConfig.Port, test.Port)
		} else if cfg.ConnConfig.Database != test.DB.DatabaseName {
			t.Fatalf("test %d: got db %s, want %q", i, cfg.ConnConfig.Database, test.DB.DatabaseName)
		} else if cfg.ConnConfig.User != test.DB.User {
			t.Fatalf("test %d: got user %s, want %q", i, cfg.ConnConfig.User, test.DB.User)
		} else if cfg.ConnConfig.Password != test.DB.Password {
			t.Fatalf("test %d: got password %s, want %q", i, cfg.ConnConfig.Password, test.DB.Password)
		}
	}
}

//go:linkname loadConfig encore.dev/runtime/config.loadConfig
func loadConfig() (*config.Config, error) { return nil, nil }