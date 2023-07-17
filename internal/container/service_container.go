package container

import "card-service/internal/service"

// ServiceProvider Inject Service
func (c *Container) ServiceProvider() {
	//	Service
	if err := c.Container.Provide(service.NewService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(service.NewCardService); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(service.NewCardInfoService); err != nil {
		c.Error = err
	}
}
