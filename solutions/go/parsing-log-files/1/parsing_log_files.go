package parsinglogfiles

import "regexp"

func IsValidLine(text string) bool {
	re := regexp.MustCompile(`^(\[TRC\]|\[DBG\]|\[INF\]|\[WRN\]|\[ERR\]|\[FTL\])`)
	return re.MatchString(text)
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=\-]*>`)
	return re.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)"[^"]*password[^"]*"`)
	count := 0
	for _, line := range lines {
		if re.MatchString(line) {
			count += 1
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line+(\d*)`)
	return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User +(\S+)`)
	result := []string{}
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match == nil {
			result = append(result, line)
		} else {
			userName := match[1]
			result = append(result, "[USR] "+userName+" "+line)
		}
	}
	return result
}
