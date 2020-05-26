package postgres

import "regexp"

func removeUserPassword(rawInput string) string {
	byteInput := []byte(rawInput)

	userRegex, err := regexp.Compile(`(user=).(\w)+.`)
	if err != nil {
		return ""
	}
	passwordRegex, err := regexp.Compile(`(password=).(\w)+.`)
	if err != nil {
		return ""
	}

	for {
		match := userRegex.FindIndex(byteInput)
		if match == nil {
			break
		}

		byteInput = append(byteInput[:match[0]], byteInput[match[1]:]...)
	}

	for {
		match := passwordRegex.FindIndex(byteInput)
		if match == nil {
			break
		}

		byteInput = append(byteInput[:match[0]], byteInput[match[1]:]...)
	}

	return string(byteInput)
}
