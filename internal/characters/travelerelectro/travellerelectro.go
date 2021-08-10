package travelerelectro

import (
	"github.com/genshinsim/gsim/pkg/character"
	"github.com/genshinsim/gsim/pkg/combat"
	"github.com/genshinsim/gsim/pkg/core"
	"go.uber.org/zap"
)

type char struct {
	*character.Tmpl
}

func init() {
	combat.RegisterCharFunc("travelerelectro", NewChar)
}

func NewChar(s core.Sim, log *zap.SugaredLogger, p core.CharacterProfile) (core.Character, error) {
	c := char{}
	t, err := character.NewTemplateChar(s, log, p)
	if err != nil {
		return nil, err
	}
	c.Tmpl = t
	c.Energy = 80
	c.EnergyMax = 80
	c.Weapon.Class = core.WeaponClassSword
	c.BurstCon = 3
	c.SkillCon = 5
	c.NormalHitNum = 5

	return &c, nil
}
