package rate

import "github.com/go-pg/pg"

type PgRepository struct {
	Db *pg.DB
}

func (repository *PgRepository) FindRateByName(name string) (rate *Rate, err error) {
	rate = new(Rate)
	err = repository.Db.Model(rate).
		Where("rate.name = ?", name).
		Select()
	return rate, err
}
