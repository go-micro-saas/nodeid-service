package main

import "flag"

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Version is the version of the compiled software.
	Version string

	// flagconf is the config flag.
	flagconf  string
	sourceDir string
	storeDir  string
)

func init() {
	flag.StringVar(&flagconf, "conf", "../../configs", "config path, eg: -conf config.yaml")
	flag.StringVar(&sourceDir, "source_dir", "", "store source path, eg: -source_dir path/to/source_dir")
	flag.StringVar(&storeDir, "store_dir", "", "custom store path, eg: -store_dir project_name/service_name/store_dir")
}

// go run ./cmd/store-configuration/... -conf=./configs
func main() {
	if !flag.Parsed() {
		flag.Parse()
	}
}
