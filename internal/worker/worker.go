package worker

import "example/internal/business_logic"

// совсем простой пример как можно использовать бизнес логику еще

type Worker struct {
	logic business_logic.Logic
}

func (w Worker) Work() {
	for {
		w.logic.Get("id")
	}
}

func New(
	logic business_logic.Logic,
) Worker {
	return Worker{logic}
}
