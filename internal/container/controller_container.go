package container

import controller_v1 "card-service/internal/controller/v1"

// ControllerProvider Inject Controller
func (c *Container) ControllerProvider() {
	// Controller
	if err := c.Container.Provide(controller_v1.NewController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controller_v1.NewPingController); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(controller_v1.NewCardController); err != nil {
		c.Error = err
	}
	if err := c.Container.Provide(controller_v1.NewCardInfoController); err != nil {
		c.Error = err
	}

}
