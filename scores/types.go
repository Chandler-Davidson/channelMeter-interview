package scores

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
	Exams     []StudentExamScore
}

type StudentExamScore struct {
	Id    int64
	Score float64
}

type ExamScore struct {
	StudentId string
	Score     float64
}

type Exam struct {
	Scores  []ExamScore
	Average float64
}

type StudentToExamsMap = map[string][]StudentExamScore
type ExamToScoresMap = map[int64][]ExamScore

type ScoreDataStore struct {
	StudentToExams StudentToExamsMap
	ExamToScores   ExamToScoresMap
}

type ScoreRepository interface {
	GetStudents() []string
	GetExams() []int64
	GetStudentScores(studentId string) Student
	GetExam(examId int64) Exam
}
