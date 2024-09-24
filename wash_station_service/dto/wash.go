package dto

type WashPackageDataCompact struct {
	ID       uint32
	Name     string
	Category uint32
	Price    float64
}

type DetailingPackageDataCompact struct {
	ID          uint32
	Name        string
	Description string
	Price       float64
}
