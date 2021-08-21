package test

import (
	"os"
	"testing"

	"github.com/ihawke/kontext/pkg"
)

func TestDir(t *testing.T) {

	// Make default
	path, _ := os.Getwd()
	inputConfig := path + "/samples/config1"
	config1, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}
	pkg.DefaultConfig = config1

	dirName := path + "/samples/"
	newConfigs := pkg.ReadDir(dirName)

	// Merge the configs
	for _, newConfig := range newConfigs {
		pkg.MergeContexts(newConfig)
	}

	kubectlOutput := path + "/samples/config1-config2-kubectl-merged"
	kubectlMergeConfig, err := pkg.ReadConfig(kubectlOutput)
	if err != nil {
		t.Fatal(err)
	}

	ConfigExist(pkg.DefaultConfig, kubectlMergeConfig, t)
}
