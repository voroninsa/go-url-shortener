package configs

import "testing"

func TestConfigs(t *testing.T) {
	expected := Config{"inmemory", "short_urls", "5432", "postgres_container", "root", ""}
	InitConfig()
	actual := CFG

	if expected != actual {
		t.Errorf("Result was incorrect, got: %s, want: %s.", actual, expected)
	}
}
