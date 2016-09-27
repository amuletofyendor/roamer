package universe

import (
  "../interop"
)

func (u *Universe) HandleDeath(com interop.Command) bool {
  return true
}
