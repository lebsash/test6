package fibb

import (
	"sync"
	"time"
)

var result, isDone sync.Map

// Calc - calculate
func Calc(jobID string, n int64) {
	var i, j int64
	res := make([]int64, 0)
	for i, j = 0, 1; j < n; i, j = i+j, i {
		time.Sleep(1000) // Тормозим обработку для наглядности
		res = append(res, i)
		result.Store(jobID, res)

	}
	isDone.Store(jobID, true)
}

func progressResult(jobID string) []int64 {
	res, ok := result.Load(jobID)
	if !ok {
		return nil
	}
	n, ok := res.([]int64)
	if !ok {
		return nil
	}
	return n
}

func isCalcDone(jobID string) bool {
	res, ok := isDone.Load(jobID)
	if !ok {
		return false
	}
	done, ok := res.(bool)
	if !ok {
		return false
	}
	return done
}

// Status - get status of worker
func Status(jobID string) ([]int64, bool) {
	return progressResult(jobID), isCalcDone(jobID)
}
