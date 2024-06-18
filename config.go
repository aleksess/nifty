package nifty

type Config struct {
	Port           uint8  `json:"port"`
	DatabaseUrl    string `json:"databaseUrl"`
	DatabaseEngine string `json:"databaseEngine"`
	SessionSecret  string `json:"sessionSecret"`
}
