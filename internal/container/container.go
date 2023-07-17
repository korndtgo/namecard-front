package container

import (
	"card-service/internal/config"
	"card-service/internal/infrastructure/database"
	"card-service/internal/infrastructure/http"
	"card-service/internal/infrastructure/server"
	"card-service/internal/util"
	"card-service/internal/util/databasehelp"
	"card-service/internal/util/log"
	"sync"

	"go.uber.org/dig"
)

// Container ...
type Container struct {
	logger    *log.Logger
	Container *dig.Container
	Error     error
}

// Configure ...
func (c *Container) Configure() {
	c.Container = dig.New()

	//Config
	if err := c.Container.Provide(config.NewConfiguration); err != nil {
		c.Error = err
	}

	//Logger
	if err := c.Container.Provide(log.NewLogger); err != nil {
		c.Error = err
	}

	//Infrastructure
	if err := c.Container.Provide(database.NewServerDB); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(http.NewHTTPClient); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(server.NewServer); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(databasehelp.NewDatabaseHelp); err != nil {
		c.Error = err
	}

	if err := c.Container.Provide(util.NewUtils); err != nil {
		c.Error = err
	}

	c.ControllerProvider()
	c.ServiceProvider()
	c.RepositoryProvider()
}

// Run ...
func (c *Container) Run() *Container {
	if err := c.Container.Invoke(func(g *server.Server) {
		wg := new(sync.WaitGroup)
		wg.Add(1)

		//go func() {
		//	if err := g.Start(); err != nil {
		//		c.logger.Println("[Container: Run] Start is panic: ", err)
		//		panic(err)
		//	}
		//	wg.Done()
		//}()
		go func() {
			if err := g.Start(); err != nil {
				c.logger.Println("[Container: Run] StartRestful is panic: ", err)
				panic(err)
			}
			wg.Done()
		}()

		wg.Wait()
	}); err != nil {
		panic(err)
	}

	return c
}

// NewContainer ...
func NewContainer() *Container {
	c := &Container{}

	c.Configure()

	return c
}
