package env

import (
	"fmt"
	"os"
	"strconv"
	"time"

	jsoniter "github.com/json-iterator/go"

	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func GetDbUri() string   { return dburi }
func GetWebAddr() string { return webAddr }

func Load() {
	if err := godotenv.Load(); err != nil {
		errAndExit("[ERROR] load env", err)
	}

	dburi = get("DB_URI", "")
	if dburi == "" {
		errAndExit("[ERROR] load env. Отсутствует DB_URI")
	}

	webAddr = get("WEB_ADDR", "")
	if webAddr == "" {
		errAndExit("[ERROR] load env. Отсутствует WEB_ADDR")
	}

}

var (
	dburi   string
	webAddr string
)

func get(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

func getDoubleSliceSlice(s string) ([][]string, error) {
	pairs := get(s, "")

	if len(pairs) == 0 {
		return nil, errors.New("no data")
	}
	var slice [][]string
	if err := jsoniter.UnmarshalFromString(pairs, &slice); err != nil {
		return nil, err
	}
	return slice, nil
}

func getBool(key string, defaultVal bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		data, err := strconv.ParseBool(value)
		if err != nil {
			return defaultVal
		}
		return data
	}
	return defaultVal
}

func getInt(key string, defaultVal int) int {
	if value, exists := os.LookupEnv(key); exists {
		data, err := strconv.Atoi(value)
		if err != nil {
			return defaultVal
		}
		return data
	}
	return defaultVal
}

func getFloat64(key string, defaultVal float64) float64 {
	if value, exists := os.LookupEnv(key); exists {
		data, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return defaultVal
		}
		return data
	}
	return defaultVal
}

func getDuration(key string, defaultVal time.Duration) time.Duration {
	if value, exists := os.LookupEnv(key); exists {
		data, err := time.ParseDuration(value)
		if err != nil {
			return defaultVal
		}
		return data
	}
	return defaultVal
}
func errAndExit(msg ...any) {
	_, _ = fmt.Fprintln(os.Stderr, msg...)
	os.Exit(1)
}
