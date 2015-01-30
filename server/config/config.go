package config

var (
	BasePath string = "/tmp/repo"
	DBPath   string = "/tmp/data"
)

func InitConfig(conf string) error {
	BasePath = "/tmp/test"
	return nil
}
