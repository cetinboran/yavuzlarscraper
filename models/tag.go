package models

import (
	"strings"
)

func TagInit() *Tag {
	return &Tag{search: searchInit()}
}

func (t *Tag) SetName(name string) {
	t.name = name
}

func (t *Tag) SetClasses(classes string) {
	t.class = strings.Split(classes, ",")
}

func (t *Tag) SetAttiributes(attribute string) {
	t.attribute = strings.Split(attribute, ",")
}

func (t *Tag) SetId(id string) {
	t.id = id
}

func (t *Tag) setSearch() {
	t.search.setSearch(*t)
}

func createTag(tagStr string) *Tag {
	// .selam .title #la div böyle düz olsun

	newTag := &Tag{search: searchInit()}
	var tagName, classes, attribute, id string

	pieces := strings.Split(tagStr, " ")

	for _, v := range pieces {
		v = strings.ReplaceAll(v, " ", "")

		str, has := strings.CutPrefix(v, ".")
		if has {
			classes += str + " "
		}

		str, has = strings.CutPrefix(v, "#")
		if has {
			id = v
		}

		if strings.HasPrefix(v, "[") && strings.HasSuffix(v, "]") {
			withOutLeft := strings.ReplaceAll(v, "[", "")
			attributeStr := strings.ReplaceAll(withOutLeft, "]", "")

			attribute += attributeStr + " "
		}

		if !strings.HasPrefix(v, "#") && !strings.HasPrefix(v, ".") && !strings.HasPrefix(v, "[") && !strings.HasSuffix(v, "]") {
			tagName = v
		}
	}

	classes = strings.TrimSpace(classes)
	attribute = strings.TrimSpace(attribute)

	newTag.name = tagName
	newTag.class = strings.Split(classes, " ")
	newTag.attribute = strings.Split(attribute, " ")
	newTag.id = id

	// Eğer attribute ve classes == "" ise hiç bişi gelmemiştir
	// ben " " ile splitlediğimden hata oluyor o yüzden bunu yapıyorum şimdilik.
	if attribute == "" {
		newTag.attribute = []string{}
	}

	if classes == "" {
		newTag.class = []string{}
	}

	return newTag
}
