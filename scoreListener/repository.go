package scoreListener

type ScoreRepository interface {
	GetStudents() []string
	GetExams() []int64
	GetStudentScores(studentId string) Student
	GetExam(examId int64) Exam
}

type MemoryScoreRepository struct {
	StudentToExams StudentToExamsMap
	ExamToScores   ExamToScoresMap
}

func NewScoreRepository() MemoryScoreRepository {
	store := MemoryScoreRepository{
		make(map[string][]StudentExamScore),
		make(map[int64][]ExamScore)}

	channel := make(chan ScoreEvent)
	go Subscribe(channel)
	go fillData(store, channel)

	return store
}

func fillData(store MemoryScoreRepository, channel chan ScoreEvent) {
	for {
		score := <-channel
		studentId := score.StudentId

		exam := StudentExamScore{score.ExamId, score.Score}
		store.StudentToExams[studentId] = append(store.StudentToExams[studentId], exam)

		studentScore := ExamScore{studentId, score.Score}
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
	total := 0.0

	for _, e := range store.StudentToExams[studentId] {
		total += e.Score
	}

	average := total / float64(len(store.StudentToExams[studentId]))
	return Student{studentId, average, store.StudentToExams[studentId]}
}

func (store *MemoryScoreRepository) GetExam(examId int64) Exam {
	total := 0.0

	for _, e := range store.ExamToScores[examId] {
		total += e.Score
	}

	average := total / float64(len(store.ExamToScores[examId]))
	return Exam{store.ExamToScores[examId], average}
}
