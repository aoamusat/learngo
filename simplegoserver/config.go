package main

import "os"

func getServerPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		return "8080"
	}
	return port
}

func getDatabaseURI() string {
	uri := os.Getenv("DB_URL")
	if uri == "" {
		return "mysql://localhost:3306"
	}
	return uri
}

func getConfig() map[string]any {
	conf := make(map[string]any)
	conf["port"] = getServerPort()
	conf["database_uri"] = getDatabaseURI()
	return conf
}
