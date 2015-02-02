package config

var (
	DatetimeFormat = "2006-01-02 15:04:05"
	DateFormat     = "2006-01-02"

	BasePath string = "/tmp/repo"
	DBPath   string = "/tmp/data"
)

func InitConfig(conf string) error {
	BasePath = "/tmp/test"
	return nil
}
