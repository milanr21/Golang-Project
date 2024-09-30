package config

type HTTPServer struct {
	Addr string
}

type Config struct {
	Env         string `yaml: "env" env: "ENV" env-required: "true"`
	StoragePath string `yaml: "storage-path" env-required: "true"`
	HTTPServer  `yaml: "http-server"`
}


function MustLoad() {
	var configPath string
	
	configPath = os.Getenv("CONFIG_PATH")

	if configPath == "" {
		flags := flag.String("Config", "", "path")
		flag.Parse()


		if configPath == "" {
			log.Fatal("Config path not provided")
		}
	}

	

	
}