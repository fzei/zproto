package zproto

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

//(^#[\s\w\.@\p{Han}]*?\n)?(\w+?)
//var msg_pat = `(\w+?) (\w+?) \{([\[\]\s\w\.#@\p{Han}]*?)*?\}`
//(#[\w\t ]+?\n)?
var msg_pat = `(#[\w\p{Han}\n]*?)?(\w+?) (\w+?) \{([\[\]\s\w\.#@\p{Han}]*?)*?\}`
var item_pat = `[\t ]*?(\w*?)[\t ]+?(\[\])?(\w*?)([\t ]*?@[\d\.}]*)?([\t ]*?#[\w\p{Han}]*)?(?:[\t ]*?\n)`

//([\t ]+?#\w*)?[\t ]*?
func MatchTest(text string) {
	reg := regexp.MustCompile(`([\t ]*?#\w*)`)
	matches := reg.FindAllStringSubmatch(text, -1)
	fmt.Printf("matches] %+v\n", matches)
}

//([\t ]+?#[\w]*?)
func MatchItems(text string) {
	reg := regexp.MustCompile(item_pat)
	matches := reg.FindAllStringSubmatch(text, -1)

	//fmt.Printf("item] %d\n", len(matches))
	for id, match := range matches {
		fmt.Printf("  item %d: %+v\n", id, match[1:])
		if match[1] == "" {
			fmt.Printf("E] no var @%d\n", id)
		}
	}
}

func MatchMessage(text string) {
	reg := regexp.MustCompile(`[\r]`)
	text = reg.ReplaceAllString(text, "")
	//fmt.Printf("text: %s\n", text)

	reg = regexp.MustCompile(msg_pat)
	matches := reg.FindAllStringSubmatch(text, -1)

	for id, match := range matches {
		//fmt.Printf("  item %d-%d: %+v\n", id, len(match[1:]), match[1:])
		fmt.Printf("%d %s: %s\n", id, match[2], match[3])
		fmt.Printf("pre] %+v\n", match[1])
		if match[2] == "msg" {
			//fmt.Printf("msg] %+v\n", match[3])
			MatchItems(match[4])
		}
		if match[2] == "enum" {
			MatchItems(match[4])
		}
	}
}

func ReadFile(name string) string {
	b, err := ioutil.ReadFile(name)
	if err != nil {
		return ""
	}
	return string(b)
}
