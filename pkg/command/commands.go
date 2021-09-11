package command

type task struct {
	title string
}

func CurrentTime() (*task, error) {
	return &task{"aaa"}, nil
}
