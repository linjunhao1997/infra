package utils

import (
	"strings"
)

type TagOptions []string
type OptionValues []string
type TagValues []string

type Value struct {
	Key string
	Val string
}

type TagSetting struct {
	Name  string
	Value string
}

type TagSettings []*TagSetting

func (settings TagSettings) GetSettingValue(name string, defaultValue string) string {
	for _, setting := range settings {
		if setting.Name == name {
			return setting.Value
		}
	}
	return defaultValue
}

func (settings TagSettings) Contains(name string) bool {
	for _, setting := range settings {
		if setting.Name == name {
			return true
		}
	}
	return false
}

// 获取tag的设置，例如"a:av;b:bv;c:cv"
func ParseTagSetting(str string, sep string) TagSettings {
	settings := TagSettings{}
	names := strings.Split(str, sep)

	for i := 0; i < len(names); i++ {
		j := i
		if len(names[j]) > 0 {
			for {
				if names[j][len(names[j])-1] == '\\' {
					i++
					names[j] = names[j][0:len(names[j])-1] + sep + names[i]
					names[i] = ""
				} else {
					break
				}
			}
		}
		values := strings.Split(names[j], ":")
		k := strings.TrimSpace(values[0])
		setting := TagSetting{}
		if k != "" {
			setting.Name = k
		}
		if len(values) >= 2 {
			setting.Value = strings.Join(values[1:], ":")
		}
		settings = append(settings, &setting)
	}

	return settings
}
