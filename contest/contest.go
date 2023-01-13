package contest

type Eater interface {
	Eat()
}

func DoEatingContest(contestants []Eater) {
	for _, contestant := range contestants {
		contestant.Eat()
	}
}
