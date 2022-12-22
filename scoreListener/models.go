package scoreListener

type Message struct {
	Event string     `json:"event"`
	Data  ScoreEvent `json:"data"`
}

type ScoreEvent struct {
	StudentId string  `json:"studentId"`
	ExamId    int64   `json:"exam"`
	Score     float64 `json:"score"`
}

type Student struct {
	StudentId string
	Average   float64
	Exams     []ExamScore
}

type ExamScore struct {
	ExamId int64
	Score  float64
}

type StudentScore struct {
	StudentId string
	Score     float64
}

type Exam struct {
	Scores  []StudentScore
	Average float64
}
