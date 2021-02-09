package messages

type Store struct {
	Link              string              `json:"link"`
	Name              string              `json:"name"`
	OwnerName         string              `json:"ownerName"`
	OwnerCastle       string              `json:"ownerCastle"`
	Kind              string              `json:"kind"`
	Mana              int                 `json:"mana"`
	Offers            []StoreOffer        `json:"offers"`
	Specialization     *map[string]int    `json:"specialization,omitempty"`
	//Specializations   StoreSpecialization `json:"specializations"`
	QualityCraftLevel int                 `json:"qualityCraftLevel"`
	MaintenanceEnabled bool               `json:"maintenanceEnabled"`
	MaintenanceCost    int                `json:"maintenanceCost"`
	GuildDiscount      int                `json:"guildDiscount"`
	CastleDiscount     int                `json:"castleDiscount"`
}

type StoreOffer struct {
	Item  string `json:"item"`
	Price int    `json:"price"`
	Mana  int    `json:"mana"`
}

type StoreSpecialization struct {
	Level  int            `json:"level"`
	Values map[string]int `json:"values"`
}
