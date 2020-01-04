package convert

import (
	. "draftDeploy/app/model"
	"regexp"
	"strings"
)

func Convert(data ReadData, err error) (ConvertData, error) {
	var converted ConvertData
	converted.OutputDir = data.OutputDir
	converted.BuildConfig = data.BuildConfig
	var res []string
	if err != nil {
		return ConvertData{}, err
	}
	for _, draft := range data.Drafts {
		b := draft.
			EachLines(trim).
			EachLines(setPlaceholders(data.BuildConfig.Placeholders)).
			EachLines(addNewLineMark).
			EachLines(addIndent)
		res = append(res, b.Data...)
	}
	converted.Data = res
	return converted, nil
}

func trim(str string) string {
	return strings.TrimSpace(str)
}

func addNewLineMark(str string) string {
	return appendPostfix(str, "\n")
}

func addIndent(str string) string {
	if !startsWithBracket(str) {
		return appendPrefix("　", str)
	} else {
		return str
	}
}

func setPlaceholders(placeholders map[string]string) func(string) string {
	return func(str string) string {
		result := str
		for key, value := range placeholders {
			r := regexp.MustCompile("#{" + key + "}")
			result = r.ReplaceAllString(result, value)
		}
		return result
	}
}

func startsWithBracket(str string) bool {
	brackets := []string{
		"「",
		"『",
		"〈",
		"［",
		"（",
		"｛",
	}
	for _, e := range brackets {
		if strings.HasPrefix(str, e) {
			return true
		}
	}
	return false
}

func appendPrefix(prefix string, str string) string {
	buf := make([]byte, 0)
	buf = append(buf, prefix...)
	buf = append(buf, str...)
	return string(buf)
}

func appendPostfix(str string, postfix string) string {
	buf := make([]byte, 0)
	buf = append(buf, str...)
	buf = append(buf, postfix...)
	return string(buf)
}
