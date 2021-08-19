package pkg

func ContextExist(contextName string) (errMsg string) {
	contexts, _ := ListDefaultContexts()

	found := false
	for _, v := range contexts {
		if v == contextName {
			found = true
		}
	}
	if found == false {
		errMsg = "Context [" + contextName + "] not found in the config file, typo? 😇"
	}
	return
}

func isNewContext(contextName string) (errMsg string) {
	contexts, _ := ListDefaultContexts()

	found := false
	for _, v := range contexts {
		if v == contextName {
			found = true
		}
	}
	if found {
		errMsg = "Context [" + contextName + "] already in the config file.  Please use [kontext -r " + contextName + "] to remove it first. 😇"
	}
	return
}

func ListDefaultContexts() (contexts []string, currentContext string) {
	c := make([]string, len(DefaultConfig.Contexts))
	for i, context := range DefaultConfig.Contexts {
		c[i] = context.Name
	}
	contexts = c
	currentContext = DefaultConfig.CurrentContext
	return
}

func MergeContexts(newConfig *Config) {
	for _, nc := range newConfig.Contexts {

		errMsg := isNewContext(nc.Name)
		if errMsg != "" {
			Warn(errMsg)
			continue
		}

		namedContext := getContextByName(nc.Name, newConfig)
		if namedContext != (NamedContext{}) {
			DefaultConfig.Contexts = append(DefaultConfig.Contexts, namedContext)
		}

		namedCluster := getClusterByName(nc.Name, newConfig)
		if namedCluster != (NamedCluster{}) {
			DefaultConfig.Clusters = append(DefaultConfig.Clusters, namedCluster)
		}

		namedUser := getUserByName(nc.Name, newConfig)
		if namedUser.Name != "" {
			DefaultConfig.Users = append(DefaultConfig.Users, namedUser)
		}
		Info("Added context [" + nc.Name + "]")
	}
}

func deleteContext(deleteContext string) {
	Info("Removing context [" + deleteContext + "]")
	for index, nc := range DefaultConfig.Contexts {
		if nc.Name == deleteContext {
			DefaultConfig.Contexts = append(DefaultConfig.Contexts[:index], DefaultConfig.Contexts[index+1:]...)
			break
		}
	}

	for index, nc := range DefaultConfig.Clusters {
		if nc.Name == deleteContext {
			DefaultConfig.Clusters = append(DefaultConfig.Clusters[:index], DefaultConfig.Clusters[index+1:]...)
			break
		}
	}

	for index, nu := range DefaultConfig.Clusters {
		if nu.Name == deleteContext {
			DefaultConfig.Users = append(DefaultConfig.Users[:index], DefaultConfig.Users[index+1:]...)
			break
		}
	}
}

func getContextByName(cName string, config *Config) NamedContext {
	for _, nc := range config.Contexts {
		if nc.Name == cName {
			return nc
		}
	}

	return NamedContext{}
}

func getClusterByName(cName string, config *Config) NamedCluster {
	for _, nc := range config.Clusters {
		if nc.Name == cName {
			return nc
		}
	}

	return NamedCluster{}
}

func getUserByName(cName string, config *Config) NamedUser {
	for _, nu := range config.Users {
		if nu.Name == cName {
			return nu
		}
	}

	return NamedUser{}
}

func addNamespace() {
	currentContext := DefaultConfig.CurrentContext
	if currentContext == "" {
		Fatal("Could not find a current-context!  Set a default context by running: [Kontext] 😇")
	}

	for i, nc := range DefaultConfig.Contexts {
		if nc.Name == currentContext {
			if Args.NS == "-" {
				DefaultConfig.Contexts[i].Namespace = ""
			} else {
				DefaultConfig.Contexts[i].Namespace = Args.NS
			}
			break
		}
	}
	return
}
