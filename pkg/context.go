package pkg

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

		namedContext := GetContextByName(nc.Name, newConfig)
		if namedContext != nil {
			DefaultConfig.Contexts = append(DefaultConfig.Contexts, *namedContext)
		}

		namedCluster := GetClusterByName(nc.Name, newConfig)
		if namedCluster != nil {
			DefaultConfig.Clusters = append(DefaultConfig.Clusters, *namedCluster)
		}

		namedUser := GetUserByName(nc.Name, newConfig)
		if namedUser != nil {
			DefaultConfig.Users = append(DefaultConfig.Users, *namedUser)
		}
		Info("Added context [" + nc.Name + "]")
	}
}

func DeleteContext(deleteContext string) bool {
	Info("Removing context [" + deleteContext + "]")
	removed := false
	for i, nc := range DefaultConfig.Contexts {
		if nc.Name == deleteContext {
			DefaultConfig.Contexts = append(DefaultConfig.Contexts[:i], DefaultConfig.Contexts[i+1:]...)
			removed = true
		}
	}

	for i, nc := range DefaultConfig.Clusters {
		if nc.Name == deleteContext {
			DefaultConfig.Clusters = append(DefaultConfig.Clusters[:i], DefaultConfig.Clusters[i+1:]...)
			removed = true
		}
	}

	for i, nu := range DefaultConfig.Clusters {
		if nu.Name == deleteContext {
			DefaultConfig.Users = append(DefaultConfig.Users[:i], DefaultConfig.Users[i+1:]...)
			removed = true
		}
	}

	if removed == false {
		Warn("Context [" + deleteContext + "] not found in the config file, typo? 😇")
	}

	return removed
}

func GetContextByName(cName string, config *Config) *NamedContext {
	for _, nc := range config.Contexts {
		if nc.Name == cName {
			return &nc
		}
	}

	return nil
}

func GetClusterByName(cName string, config *Config) *NamedCluster {
	for _, nc := range config.Clusters {
		if nc.Name == cName {
			return &nc
		}
	}

	return nil
}

func GetUserByName(cName string, config *Config) *NamedUser {
	for _, nu := range config.Users {
		if nu.Name == cName {
			return &nu
		}
	}

	return nil
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
