package lessoncodebasic

import (
	"errors"
	"time"
)

type nonCriticalError struct{}

func (e nonCriticalError) Error() string {
	return "validation error"
}

var (
	errBadConnection = errors.New("bad connection")
	errBadRequest    = errors.New("bad request")
)

const unknownErrorMsg = "unknown error"

/* Реализуйте функцию GetErrorMsg(err error) string, которая возвращает текст ошибки,
 если она критичная. В случае неизвестной ошибки возвращается строка unknown error: */
func GetErrorMsg(err error) string {
if errors.Is(err,  nonCriticalError{} ) {
	return ""
} else if errors.Is(err, errBadConnection) {
return errBadConnection.Error()
} else if errors.Is(err, errBadRequest) {
return errBadRequest.Error()
}
	return unknownErrorMsg
}

//___________________________________________________________________________________________

type MergeDictsJob struct {
	Dicts      []map[string]string
	Merged     map[string]string
	IsFinished bool
}

// errors
var (
	errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
	errNilDict        = errors.New("nil dictionary")
)

/*
Реализуйте функцию ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error),

	которая выполняет джобу MergeDictsJob и возвращает ее. Алгоритм обработки джобы следующий:

- перебрать по порядку все словари job.Dicts и записать каждое ключ-значение в результирующую мапу job.Merged
- если в структуре job.Dicts меньше 2-х словарей, возвращается ошибка
errNotEnoughDicts = errors.New("at least 2 dictionaries are required")
- если в структуре job.Dicts встречается словарь в виде нулевого значения nil,
то возвращается ошибка errNilDict = errors.New("nil dictionary")
- независимо от успешного выполнения или ошибки в возвращаемой структуре MergeDictsJob
поле IsFinished должно быть заполнено как true
*/
func ExecuteMergeDictsJob(job *MergeDictsJob) (*MergeDictsJob, error) {

	job.IsFinished = true
	if len(job.Dicts) < 2 {
		return job, errNotEnoughDicts
	}
	job.Merged = make(map[string]string)
	for i := 0; i < len(job.Dicts); i++ {
		if job.Dicts[i] == nil {
return job, errNilDict
		}
		for key, vol := range job.Dicts[i] {
			job.Merged[key] = vol 
		}
	}

	return job, nil
}
//____________________________________________________________________________________________
/*
Реализуйте функцию MaxSum(nums1, nums2 []int) []int из прошлого задания,
 используя горутины для расчета каждой суммы слайса.
Не забудьте использовать функцию time.Sleep(100 * time.Millisecond),
 чтобы сумма успела посчитаться. В настоящих приложениях используются специальные инструменты,
чтобы дожидаться исполнения асинхронного кода, но для простоты здесь будет использоваться обычный сон.
*/
func MaxSum(nums1, nums2 []int) []int {
	r1, r2 := 0, 0
	go Sum(nums1, &r1)
	go Sum(nums2, &r2)
	time.Sleep(100 * time.Millisecond)
	if r2 > r1 {
		return nums2
	}
return nums1
}

func Sum(sl []int, r *int)  {
	for _, vol := range sl {
	*r += vol
	}
}
//___________________________________________________________________________________________________
/*
Реализуйте функцию-воркера SumWorker(numsCh chan []int, sumCh chan int), 
которая суммирует переданные числа из канала numsCh и передает результат в канал sumCh:
*/
func SumWorker(numsCh chan []int, sumCh chan int) {
	var result int
	for _, vol := range <-numsCh {
		result += vol
	}
	sumCh <- result
	}
//________________________________________________________________________________________________