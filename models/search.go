package models

import (
	"fmt"
	"regexp"
	"strings"
)

func SearchInit() *Search {
	return &Search{}
}

func (s *Search) setSearch(t Tag) {
	s.createRegex(t)
	s.setEnd(t)
}

func (s *Search) createRegex(t Tag) {
	// <div[^>]*class="selam"[^>]*id="3"[^>]*> working regex

	regex := fmt.Sprintf(`<%v`, t.Name)

	if len(t.class) > 0 {
		classes := strings.Join(t.class, " ")

		regex += `\s+[^>]*class="`
		regex += classes
		regex += `"`
	}

	if t.id != "" {
		regex += `[^>]*`
		regex += fmt.Sprintf(`id="%v"`, t.id)
	}

	regex += `[^>]*>`

	s.StartReg = regex
}

func (s *Search) RegexCheck(t Tag, data string) bool {
	// Regex'i işliyorum
	reg := regexp.MustCompile(s.StartReg)

	// gelen Data'da regex var ise true yok ise false dönüyor.
	return reg.MatchString(data)
}

func (s *Search) setEnd(t Tag) {
	end := fmt.Sprintf("/%v", t.Name)
	s.End = end
}
