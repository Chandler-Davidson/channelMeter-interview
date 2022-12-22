package scoreListener

type ScoreRepository interface {
	GetStudents() []string
	GetExams() []int64
	GetStudentScores(studentId string) Student
	GetExam(examId int64) Exam
}

type StudentToExamsMap = map[string][]ExamScore
type ExamToScoresMap = map[int64][]StudentScore

type MemoryScoreRepository struct {
	StudentToExams StudentToExamsMap
	ExamToScores   ExamToScoresMap
}

func NewScoreRepository() MemoryScoreRepository {
	store := MemoryScoreRepository{
		make(map[string][]ExamScore),
		make(map[int64][]StudentScore)}

	channel := make(chan ScoreEvent)
	go Subscribe(channel)
	go fillData(store, channel)

	return store
}

func fillData(store MemoryScoreRepository, channel chan ScoreEvent) {
	for {
		score := <-channel
		studentId := score.StudentId

		exam := ExamScore{score.ExamId, score.Score}
		store.StudentToExams[studentId] = append(store.StudentToExams[studentId], exam)

		studentScore := StudentScore{studentId, score.Score}
		store.ExamToScores[score.ExamId] = append(store.ExamToScores[score.ExamId], studentScore)
	}
}

func (store *MemoryScoreRepository) GetStudents() []string {
	students := make([]string, 0, len(store.StudentToExams))
	for s := range store.StudentToExams {
		students = append(students, s)
	}

	return students
}

func (store *MemoryScoreRepository) GetExams() []int64 {
	exams := make([]int64, 0, len(store.ExamToScores))
	for e := range store.ExamToScores {
		exams = append(exams, e)
	}

	return exams
}

func (store *MemoryScoreRepository) GetStudentScores(studentId string) Student {
	student := store.StudentToExams[studentId]

	if student == nil {
		return Student{}
	}

	total := 0.0

	for _, e := range student {
		total += e.Score
	}

	average := total / float64(len(student))
	return Student{studentId, average, student}
}

func (store *MemoryScoreRepository) GetExam(examId int64) Exam {
	exam := store.ExamToScores[examId]

	if exam == nil {
		return Exam{}
	}

	total := 0.0

	for _, e := range exam {
		total += e.Score
	}

	average := total / float64(len(exam))
	return Exam{exam, average}
}
