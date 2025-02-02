package main

import "ms-products/src/config"

func main() {
	dependencies := config.NewResolveDependencies()
	config.RunServer(dependencies.ApplicationConfigs.App.Port, dependencies)
}
