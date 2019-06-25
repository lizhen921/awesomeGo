package jw

type Worker struct {
	WorkPool chan chan Job
	JobPool chan Job
}

