package data

type Safety int

const (
	TRUSTED Safety = iota
	DEFAULT Safety = iota
	DANGER  Safety = iota
	ALL     Safety = iota
)
