//Package message

package msg

import (
	"fmt"
	"net/http"
	"strings"
	"sync"

	"github.com/Arfiandimas/kaj-rest-engine-go/src/pkg/util"
)

var msgs map[string]*Message
var once sync.Once

// MessageConfig as messages configuration
type MessageConfig struct {
	Messages []*Message `yaml:"messages"`
}

// Message configuration structure
type Message struct {
	Name     string     `yaml:"name"`
	Code     int        `yaml:"code"`
	Contents []*Content `yaml:"contents"`
	contents map[string]*Content
}

// doMap create content map from slice
func (m *Message) doMap() *Message {
	m.contents = make(map[string]*Content, 0)
	for _, c := range m.Contents {
		l := strings.ToLower(c.Lang)
		if _, ok := m.contents[l]; !ok {
			m.contents[l] = c
			continue
		}
	}
	return m
}

// Content message content configuration structure
type Content struct {
	Lang string `yaml:"lang"`
	Text string `yaml:"text"`
}

// Setup initializes messages  from yaml file
// args:
//
//	path: path of message list definition file
//
// returns:
//
//	err: operation error
func Setup(fname string, paths ...string) (err error) {
	var mcfg MessageConfig
	once.Do(func() {
		msgs = make(map[string]*Message, 0)
		for _, p := range paths {
			f := fmt.Sprint(p, fname)
			err = util.ReadFromYAML(f, &mcfg)
			if err != nil {
				continue
			}
			err = nil
		}
	})

	if err != nil {
		err = fmt.Errorf("unable to read config from files %s", err.Error())
		return
	}
	for _, m := range mcfg.Messages {
		if _, ok := msgs[m.Name]; !ok {
			m := &Message{Name: m.Name, Code: m.Code, Contents: m.Contents}
			msgs[m.Name] = m.doMap()
		}
	}
	return
}

// Get messages by language
func Get(key, lang string, msg interface{}) (text string) {
	lang = cleanLangStr(lang)
	if msg == nil {
		if m, ok := msgs[key]; ok {
			if c, ok := m.contents[lang]; ok {
				text = c.Text
				return
			}
		}
	}
	text = util.ToString(msg)
	return
}

// GetCode messages by language
func GetCode(key string) int {
	if m, ok := msgs[key]; ok {
		return m.Code
	}
	return http.StatusUnprocessableEntity
}

func cleanLangStr(s string) string {
	return strings.ToLower(strings.Trim(s, " "))
}
