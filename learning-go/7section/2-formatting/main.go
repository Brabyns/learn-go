package main

import "fmt"

type ConfigItem struct {
	Key   string
	Value interface{}
	IsSet bool
}

func (c ConfigItem) String() string {
	return fmt.Sprintf("Key: %s, Value: %s, IsSet: %t", c.Key, c.Value, c.IsSet)
}
func main() {
	appName := "EnvParser"
	version := 1.2
	port := 8080
	isEnabled := true

	status := fmt.Sprintf("Application: %s (Version: %.2f) running on port %d. Enabled: %t", appName, version, port, isEnabled)
	fmt.Println(status)
}