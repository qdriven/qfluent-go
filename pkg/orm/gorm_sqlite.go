package orm

import (
	"github.com/qdriven/qfluent-go/config"
	"path/filepath"
)

type Sqlite struct {
	config.GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	return filepath.Join(s.Path, s.Dbname+".db")
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
