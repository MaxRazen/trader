package storage

type IStorage interface {
	New()
	GetOrders()
	Migrate()
}
