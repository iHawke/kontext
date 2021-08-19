package pkg

func AddNamespace() {
	ns := Args.NS
	if Args.NS == "-"{
		ns = "default"
	}
	Info("Setting default namespace to [" + ns + "]")
	addNamespace()
	Commit(false)
}
