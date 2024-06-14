package config_test

import (
	"testing"

	"github.com/ebisu2018/xBlog/config"
)

func TestLoadFromToml(t *testing.T) {
	err := config.LoadFromTomlFile()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(config.ReadConfig())
}