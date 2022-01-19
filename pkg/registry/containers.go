package registry

import (
	"github.com/SeoSoojin/TKG_Kards/pkg/domain/services"
	"github.com/SeoSoojin/TKG_Kards/pkg/gateway/neo4jHandler"
	"github.com/SeoSoojin/TKG_Kards/pkg/usecases"
	"github.com/neo4j/neo4j-go-driver/v4/neo4j"
	"github.com/sarulabs/di"
)

type Container struct {
	ctn di.Container
}

//Get method implemented for this struct
func (containter *Container) Get(name string) interface{} {

	return containter.ctn.Get(name)
}

//Clean method implemented for this struct
func (container *Container) Clean() error {

	return container.ctn.Clean()
}

//Container creator, receives an []di.Def as param and returns an address of Container or a error
func NewContainer(defs []di.Def) (*Container, error) {

	builder, err := di.NewBuilder()
	if err != nil {
		return nil, err
	}
	err = builder.Add(defs...)
	if err != nil {
		return nil, err
	}
	return &Container{ctn: builder.Build()}, nil
}

func NewCardGetterContainer(driver *neo4j.Driver) (*Container, error) {

	defs := []di.Def{
		{
			Name: "CardGetter",
			Build: func(ctn di.Container) (interface{}, error) {

				return usecases.NewUCCardGetter(services.NewCardGetterService(neo4jHandler.NewNeo4jGetter(driver))), nil

			},
		},
	}
	return NewContainer(defs)
}
