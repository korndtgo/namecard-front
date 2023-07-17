package container

import "card-service/internal/repository"

// RepositoryProvider Inject Repository
func (c *Container) RepositoryProvider() {
	//	Repository
	if err := c.Container.Provide(repository.NewRepository); err != nil {
		c.Error = err
	}
	if err := c.Container.Provide(repository.NewTimeRepository); err != nil {
		c.Error = err
	}
	if err := c.Container.Provide(repository.NewCardRepository); err != nil {
		c.Error = err
	}
}
