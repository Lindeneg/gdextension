package utils

import "fmt"

type Step struct {
	Msg      string
	Callback func() error
}

type Steps struct {
	steps  []Step
	offset int
	total  int
}

func MakeSteps() *Steps {
	return &Steps{steps: make([]Step, 0), offset: 1, total: 0}
}

func (s *Steps) Execute() error {
	for _, st := range s.steps {
		fmt.Printf("(%d/%d) %s\n", s.offset, s.total, st.Msg)
		if err := st.Callback(); err != nil {
			return err
		}
		s.offset += 1
	}
	return nil
}

func (s *Steps) Add(sts ...Step) *Steps {
	s.steps = append(s.steps, sts...)
	s.total = len(s.steps)
	return s
}
