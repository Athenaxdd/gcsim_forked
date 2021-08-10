package generic

import (
	"github.com/genshinsim/gsim/pkg/combat"
	"github.com/genshinsim/gsim/pkg/core"
)

func init() {
	combat.RegisterWeaponFunc("rust", weapon)
}

func weapon(c core.Character, s core.Sim, log core.Logger, r int, param map[string]int) {
	m := make([]float64, core.EndStatType)
	inc := .3 + float64(r)*0.1
	c.AddMod(core.CharStatMod{
		Key: "rust",
		Amount: func(a core.AttackTag) ([]float64, bool) {
			if a == core.AttackTagNormal {
				m[core.DmgP] = inc
				return m, true
			}
			if a == core.AttackTagExtra {
				m[core.DmgP] = -0.1
				return m, true
			}
			return nil, false
		},
		Expiry: -1,
	})
}
