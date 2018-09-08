package zomato

type CommonAPI interface {
	Categories() ([]Category, error)
	Cities() ([]City, error)
}
