package monsters

import "eldoria/inventory"

type Monster struct {
	Name       string
	MaxHP      int
	HP         int
	Attack     int
	Initiative int

	XPReward   int
	GoldReward int
	IsBoss     bool
	Loot       []inventory.Item
}
