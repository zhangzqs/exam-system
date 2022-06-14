package service

import "github.com/zhangzqs/exam-system/repository"

type PaperRequestBody struct {
	Type      string        `json:"type"`
	Title     string        `json:"title"`
	Questions []QidAndScore `json:"questions"`
}
type QidAndScore struct {
	Qid   int     `json:"qid"`
	Score float64 `json:"score"`
}

// AddPaperManual 手动添加试卷
func AddPaperManual(uid int, paper *PaperRequestBody) (pid int, err error) {
	pid, err = repository.AddPaper(uid, paper.Title)
	if err != nil {
		return 0, err
	}
	for _, q := range paper.Questions {
		err := repository.PutPaperQuestion(q.Qid, pid, q.Score)
		if err != nil {
			return 0, err
		}
	}
	return pid, nil
}

type PaperQuestion struct {
	Question Question `json:"question"`
	Score    float64  `json:"score"`
}
type PaperContent struct {
	Pid       int             `json:"pid"`
	Title     string          `json:"title"`
	CreatedBy int             `json:"createdBy"`
	Questions []PaperQuestion `json:"questions"`
}

func GetPaper(uid int, pid int) (*PaperContent, error) {
	pi, err := repository.GetPaperInfo(pid)
	if err != nil {
		return nil, err
	}
	pqs, err := repository.GetPaperQuestions(pid)
	if err != nil {
		return nil, err
	}

	var qs []PaperQuestion
	for _, pq := range pqs {
		var spq PaperQuestion
		spq.Score = pq.Score
		q, err := GetQuestion(uid, pq.Qid)
		if err != nil {
			return nil, err
		}
		spq.Question = *q

		qs = append(qs, spq)
	}
	return &PaperContent{
		Pid:       pi.Pid,
		Title:     pi.Title,
		CreatedBy: pi.CreatedBy,
		Questions: qs,
	}, nil
}

func GetUserPapers(uid int) ([]PaperContent, error) {
	pis, err := repository.GetPaperInfosByUid(uid)
	if err != nil {
		return nil, err
	}
	var pcs []PaperContent
	for _, pi := range pis {
		pcs = append(pcs, PaperContent{
			Pid:       pi.Pid,
			Title:     pi.Title,
			CreatedBy: pi.CreatedBy,
			Questions: nil,
		})
	}
	return pcs, nil
}
