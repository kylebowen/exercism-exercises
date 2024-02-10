package expenses

import "errors"

// Record represents an expense record.
type Record struct {
	Day      int
	Amount   float64
	Category string
}

// DaysPeriod represents a period of days for expenses.
type DaysPeriod struct {
	From int
	To   int
}

// Filter returns the records for which the predicate function returns true.
func Filter(in []Record, predicate func(Record) bool) []Record {
	answer := []Record{}

	for _, record := range in {
		if predicate(record) {
			answer = append(answer, record)
		}
	}

	return answer
}

// ByDaysPeriod returns predicate function that returns true when
// the day of the record is inside the period of day and false otherwise.
func ByDaysPeriod(period DaysPeriod) func(Record) bool {
	return func(record Record) bool {
		return period.From <= record.Day && record.Day <= period.To
	}
}

// ByCategory returns predicate function that returns true when
// the category of the record is the same as the provided category
// and false otherwise.
func ByCategory(category string) func(Record) bool {
	return func(record Record) bool {
		return record.Category == category
	}
}

// TotalByPeriod returns total amount of expenses for records
// inside the period p.
func TotalByPeriod(in []Record, period DaysPeriod) float64 {
	var total float64

	for _, record := range in {
		if ByDaysPeriod(period)(record) {
			total += record.Amount
		}
	}

	return total
}

// CategoryExpenses returns total amount of expenses for records
// in category c that are also inside the period p.
// An error must be returned only if there are no records in the list that belong
// to the given category, regardless of period of time.
func CategoryExpenses(in []Record, period DaysPeriod, category string) (float64, error) {
	var total float64

	inCategory := Filter(in, ByCategory(category))
	if len(inCategory) == 0 {
		return total, errors.New("no records")
	}

	inPeriod := Filter(inCategory, ByDaysPeriod(period))
	for _, record := range inPeriod {
		total += record.Amount
	}

	return total, nil
}
