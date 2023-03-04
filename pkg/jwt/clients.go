package jwt

var clients map[string]*JWT

// InitClient init client
func InitClient(name string, config *Config) {
	clients[name] = New(config)
}

// InitClients init clients
func InitClients(configs map[string]*Config) {
	for name, config := range configs {
		InitClient(name, config)
	}
}

// GetClient get client
func GetClient(name string) *JWT {
	client, ok := clients[name]
	if !ok {
		return nil
	}

	return client
}
