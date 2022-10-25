package main

const DefaultConfFile = "/opt/config.toml"

//var RxpNumbersOnly *regexp.Regexp

//noinspection GoUnusedExportedFunction
func StringSliceContains(sl []string, item string) bool {
	for _, i := range sl {
		if i == item {
			return true
		}
	}
	return false
}
