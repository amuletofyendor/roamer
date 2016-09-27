package creature

import (
	"../interop"
	"../item"
	"../util"
	"../dice"
	"fmt"
	"math/rand"
)

func (c *Creature) Init() {
	c.hp = c.MaxHp()
}

func (c *Creature) MaxHp() int {
	return int(float64(c.baseStats.maxHp) * (1.0 + (float64(c.level) / 10.0)))
}

func (c *Creature) MaxStr() int {
	return int(float64(c.baseStats.maxStr) * (1.0 + (float64(c.level) / 15.0)))
}

func (c *Creature) EquippedWeapon() interop.InventoryItem {
	return c.equippedWeapon
}

func (c *Creature) Vit() int {
	return util.Max(c.vitBoost,
		util.Min(c.baseStats.maxVit,
			int(c.baseStats.vitMultiplier*float64(c.level))))
}

func (c *Creature) Spd() int {
	return util.Min(c.baseStats.maxSpd, int(c.baseStats.spdMultiplier*float64(c.level)))
}

func (c *Creature) IsAlive() bool {
	return c.hp > 0
}

func (c *Creature) AddExp(exp, currentTick, currentSubTick int) {
	c.exp += exp
	c.RecordExpGain(exp, currentTick, currentSubTick)
	oldLevel := c.level
	c.level = 1 + (c.exp / 1000)
	if c.level > oldLevel {
		c.RecordLevelGain(c.level-oldLevel, currentTick, currentSubTick)
	}
}

func (c *Creature) AlterHp(amount, currentTick, currentSubTick int) {
	if amount < 0 {
		c.RecordDamage(-amount, currentTick, currentSubTick)
	} else if amount > 0 {
		c.RecordHeal(amount, currentTick, currentSubTick)
	}

	c.hp += amount

	if c.hp <= 0 {
		c.hp = 0
		c.RecordDeath(currentTick, currentSubTick)
	} else if c.hp > c.MaxHp() {
		c.hp = c.MaxHp()
	}
}

func (c *Creature) AcceptAttack(dmg, spd, currentTick, currentSubTick int) {
	c.RecordAcceptAttack(currentTick, currentSubTick)
	if dice.Oned8() == 8 {
		c.RecordEvadeAttack(currentTick, currentSubTick)
	} else {
		c.AlterHp(-dmg, currentTick, currentSubTick)
	}
}

func (c *Creature) MakeAttack(currentTick, currentSubTick int) (int, int) {
	wi := c.EquippedWeapon()

	if wi != nil {
		weapon := wi.(item.Weapon)
		c.RecordMakeAttack(weapon.Name(), currentTick, currentSubTick)
		return weapon.BaseDamage() + c.MaxStr(), c.Spd()
	} else {
		c.RecordMakeAttack("nothing", currentTick, currentSubTick)
		return 1, c.Spd()
	}
}

func (c *Creature) AwardExp() int {
	return int(c.baseStats.expMultiplier * float64(c.level))
}

func (c *Creature) AwardExpFor(other interop.ExpProvider, currentTick, currentSubTick int) {
	c.AddExp(other.AwardExp(),
		currentTick,
		currentSubTick)
}

func (c *Creature) LookAt(objectName string, currentTick, currentSubTick int) {
	c.RecordLookAt(objectName, currentTick, currentSubTick)
}

func (c *Creature) Take(i interop.InventoryItem, currentTick, currentSubTick int) {
	itemRune := c.Add(i)
	c.RecordTake(fmt.Sprintf("%c - %v", itemRune, i.Name()), currentTick, currentSubTick)
}

func (c *Creature) Wield(weapon interop.InventoryItem, currentTick, currentSubTick int) {
	c.equippedWeapon = weapon
	c.RecordEquip(weapon.Name(), currentTick, currentSubTick)
}

func (c *Creature) StatTick(currentTick, currentSubTick int) {
	if c.IsAlive() &&
		c.hp < c.MaxHp() &&
		rand.Intn((100-util.Min(100, c.Vit()))+1) == 0 {
		if c.satiation > 25 {
			c.AlterHp(c.MaxHp()/10, currentTick, currentSubTick)
		} else {
			c.RecordHungerPreventedHealing(currentTick, currentSubTick)
		}
	}

	if c.vitBoost > 0 {
		c.vitBoost = c.vitBoost - 5
	}

	if c.baseStats.metabolism > 0 {
		if rand.Intn((100-util.Min(100, c.baseStats.metabolism))+1) == 0 {
			c.satiation--
			if c.satiation <= 0 {
				c.satiation = 0
			}
		}
	}
}

func (c *Creature) AsString() string {
	return fmt.Sprintf("Lvl %v Exp %v HP %v/%v Vit %v/%v Spd %v/%v %v",
		c.level,
		c.exp,
		c.hp,
		c.MaxHp(),
		c.Vit(),
		c.baseStats.maxVit,
		c.Spd(),
		c.baseStats.maxSpd,
		c.SatiationAsString())
}

func (c *Creature) SatiationAsString() string {
	switch {
	case c.satiation >= 100:
		return "about to burst"
	case c.satiation > 90:
		return "very full"
	case c.satiation > 75:
		return "satiated"
	case c.satiation > 50:
		return ""
	case c.satiation > 40:
		return "peckish"
	case c.satiation > 25:
		return "hungry"
	default:
		return "starved"
	}
}

func (c *Creature) Consume(foodItem interop.InventoryFoodItem,
	currentTick, currentSubTick int) {
	if foodItem.Edible() {
		c.satiation += foodItem.SatiationValue()
		c.RecordSatiation(currentTick, currentSubTick)
	} else {
		c.satiation -= foodItem.SatiationValue()
		c.RecordGasticDistress(currentTick, currentSubTick)
	}
}
