package vtils

import "github.com/AlecAivazis/survey/v2"

// Confirm  promptUi Confirm
func Confirm(msg string) bool {
	name := false
	prompt := &survey.Confirm{
		Message: msg,
	}
	_ = survey.AskOne(prompt, &name)

	return name
}

// Int32ToString  int32 to string
func Int32ToString(n int32) string {
	buf := [11]byte{}
	pos := len(buf)
	i := int64(n)
	signed := i < 0
	if signed {
		i = -i
	}
	for {
		pos--
		buf[pos], i = '0'+byte(i%10), i/10
		if i == 0 {
			if signed {
				pos--
				buf[pos] = '-'
			}
			return string(buf[pos:])
		}
	}
}
