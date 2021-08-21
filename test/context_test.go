package test

import (
	"os"
	"testing"

	"github.com/ihawke/kontext/pkg"
)

func TestMerge(t *testing.T) {
	path, _ := os.Getwd()
	inputConfig := path + "/samples/config1"
	config1, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}

	inputConfig = path + "/samples/config2"
	config2, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}

	mergedOutput := path + "/samples/config1-config2-kontext-merged"
	pkg.DefaultConfigFileName = mergedOutput
	pkg.DefaultConfig = config1
	pkg.MergeContexts(config2)
	pkg.Commit(false)

	kontextMergedConfig, err := pkg.ReadConfig(mergedOutput)
	if err != nil {
		t.Fatal(err)
	}

	kubectlOutput := path + "/samples/config1-config2-kubectl-merged"
	kubectlMergeConfig, err := pkg.ReadConfig(kubectlOutput)
	if err != nil {
		t.Fatal(err)
	}

	ConfigExist(kontextMergedConfig, kubectlMergeConfig, t)
}

func TestDelete(t *testing.T) {
	path, _ := os.Getwd()
	inputConfig := path + "/samples/config1-config2-kubectl-merged"
	mergedConfig, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}

	pkg.DefaultConfig = mergedConfig
	pkg.DeleteContext("minikube")
	pkg.DeleteContext("arn:aws:eks:us-east-1:1231716704631:cluster/test")

	config1, err := pkg.ReadConfig(inputConfig)
	if err != nil {
		t.Fatal(err)
	}

	ConfigExist(config1, mergedConfig, t)
}