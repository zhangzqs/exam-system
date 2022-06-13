package main

import (
	"encoding/json"
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

type QuestionMenu struct {
	Xx        string `json:"xx"`
	XxContent string `json:"xxContent"`
}
type SrcQuestion struct {
	Answer string         `json:"ans"`
	Menu   []QuestionMenu `json:"retMenu"`
	Title  string         `json:"title"`
}

func GetSrcQuestions() []SrcQuestion {
	var src []map[string]string
	bs, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}
	err = json.Unmarshal(bs, &src)
	if err != nil {
		log.Fatalln(err)
	}
	var sqs []SrcQuestion
	for _, m := range src {
		qs := m["content"]
		var sq SrcQuestion
		err := json.Unmarshal([]byte(qs), &sq)
		if err != nil {
			log.Fatalln(err)
		}
		sqs = append(sqs, sq)
	}
	return sqs
}

var (
	Single   = "single"
	Multiple = "multiple"
	Fill     = "fill"
	Judge    = "judge"
)

func (r *SrcQuestion) GetAns() (string, string) {
	ss := strings.Split(r.Answer, "：")
	for i := 0; i < len(ss); i++ {
		ss[i] = strings.TrimSpace(ss[i])
	}
	return ss[0], ss[1]
}
func (r *SrcQuestion) GetType() string {
	s1, s2 := r.GetAns()
	switch s1 {
	case "正确答案":
		switch s2 {
		case "×":
			return Judge
		case "√":
			return Judge
		default:
			switch len(s2) {
			case 1:
				return Single
			default:
				return Multiple
			}
		}
	case "第一空":
		return Fill
	default:
		log.Fatalln(s1)
	}
	return ""

}

type SingleQuestion struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  int      `json:"answer"`
}

func (r *SrcQuestion) ConvToSingle() *SingleQuestion {
	return &SingleQuestion{
		Title: r.Title,
		Options: func() []string {
			var ss []string
			for _, menu := range r.Menu {
				ss = append(ss, menu.XxContent)
			}
			return ss
		}(),
		Answer: func() int {
			_, s := r.GetAns()
			return int(s[0] - 'A')
		}(),
	}
}

type MultipleQuestion struct {
	Title   string   `json:"title"`
	Options []string `json:"options"`
	Answer  []int    `json:"answer"`
}

func (r *SrcQuestion) ConvToMultiple() *MultipleQuestion {
	return &MultipleQuestion{
		Title: r.Title,
		Options: func() []string {
			var ss []string
			for _, menu := range r.Menu {
				ss = append(ss, menu.XxContent)
			}
			return ss
		}(),
		Answer: func() []int {
			_, s := r.GetAns()
			var is []int
			for _, c := range s {
				is = append(is, int(c-'A'))
			}
			return is
		}(),
	}
}

type FillQuestion struct {
	Title  string   `json:"title"`
	Answer []string `json:"answer"`
}

func (r *SrcQuestion) ConvToFill() *FillQuestion {
	return &FillQuestion{
		Title: r.Title,
		Answer: func() []string {
			_, s := r.GetAns()
			return []string{s}
		}(),
	}
}

type JudgeQuestion struct {
	Title  string `json:"title"`
	Answer bool   `json:"answer"`
}

func (r *SrcQuestion) ConvToJudge() *JudgeQuestion {
	return &JudgeQuestion{
		Title: r.Title,
		Answer: func() bool {
			_, s := r.GetAns()
			return s == "√"
		}(),
	}
}

type TargetQuestion struct {
	Type    string `json:"type"`
	Content any    `json:"content"`
}

func (r *SrcQuestion) ConvToTarget() *TargetQuestion {
	return &TargetQuestion{
		Type: r.GetType(),
		Content: func() any {
			switch r.GetType() {
			case Single:
				return r.ConvToSingle()
			case Multiple:
				return r.ConvToMultiple()
			case Fill:
				return r.ConvToFill()
			case Judge:
				return r.ConvToJudge()
			default:
				return nil
			}
		}(),
	}
}

func main() {
	var qts []TargetQuestion
	for _, sq := range GetSrcQuestions() {
		qts = append(qts, *sq.ConvToTarget())
	}
	jsonStr, err := json.Marshal(qts)
	if err != nil {
		log.Fatalln(err)
	}
	ioutil.WriteFile("output.json", jsonStr, fs.ModePerm)
}
