package resources

import "arkanoid/lib/components"

// Resources contains reference to data not related to any entity
type Resources struct {
	ScreenDimensions *ScreenDimensions
	Controls         *Controls
	InputHandler     *InputHandler
	SpriteSheets     *map[string]components.SpriteSheet
}

// InitResources initializes resources
func InitResources() *Resources {
	return &Resources{}
}
