package test

import (
	"os"
	"testing"

	"github.com/ihawke/kontext/pkg"
)

func TestCommit(t *testing.T) {
	path, _ := os.Getwd()
	configFileName := path + "/samples/config"

	config, err := pkg.ReadConfig(configFileName)
	if err != nil {
		t.Fatal(err)
	}

	outFile := path + "/samples/config.out"
	pkg.DefaultConfig = config
	pkg.DefaultConfigFileName = outFile
	pkg.Commit(false)

	config2, err := pkg.ReadConfig(outFile)
	if err != nil {
		t.Fatal(err)
	}

	for i, c := range config2.Contexts {
		if c.Name != config.Contexts[i].Name {
			t.Fatal("Context 'name' mis-match at index ", i)
		}
		if c.Cluster != config.Contexts[i].Cluster {
			t.Fatal("Context 'cluster' mis-match at index ", i)
		}
	}

	for i, c := range config2.Clusters {
		if c.Name != config.Clusters[i].Name {
			t.Fatal("Cluster 'name' mis-match at index ", i)
		}
		if c.Cluster.Server != config.Clusters[i].Cluster.Server {
			t.Fatal("Cluster 'server' mis-match at index ", i)
		}
	}

	for i, u := range config2.Users {
		if u.Name != config.Users[i].Name {
			t.Fatal("User 'name' mis-match at index ", i)
		}
	}

	if config.CurrentContext != config2.CurrentContext {
		t.Fatal("Current-context mis-match!")
	}
}

func TestReadConfig(t *testing.T) {
	path, _ := os.Getwd()
	configFileName := path + "/samples/config"

	c, err := pkg.ReadConfig(configFileName)
	if err != nil {
		t.Fatal(err)
	}

	if c == nil {
		t.Fatal("Failed to read file: ", configFileName)
	}
}
