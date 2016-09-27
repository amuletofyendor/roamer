package behavior

import (
	"../../interop"
)

type CreatureBehavior interface {
	Tick(e interop.Entity, com interop.Command, l interop.Level)
}
