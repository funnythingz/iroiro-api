package repositories

import (
	"github.com/funnythingz/iroiro-api/domain"
	"github.com/funnythingz/iroiro-api/factories"
	"github.com/funnythingz/iroiro-api/mapper"
	_ "github.com/k0kubun/pp"
)

type IroRepository struct{}

func (r *IroRepository) Store(iro *domain.Iro) {
	mi := mapper.Iro{}
	mi.Map(iro)
	mi.Commit()
	iro.Id = mi.Id
}

func (r *IroRepository) Update(iro *domain.Iro) {
	//TODO
}

func (r *IroRepository) Resolve(id int) domain.Iro {
	mi := mapper.Iro{}
	mi.Fetch(id)
	return factories.IroFact.CreateIro(mi)
}

func (r *IroRepository) ResolveList(permit int, page int) []domain.Iro {
	mii := mapper.IroIro{}
	mii.Fetch(permit, page)
	return factories.IroFact.CreateIroIro(mii)
}

var (
	IroRepo = &IroRepository{}
)
