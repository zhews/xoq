package domain

type Statistic interface {
	Won()
	Lost()
	Draw()
	Get() (int, int, int, int)
}
