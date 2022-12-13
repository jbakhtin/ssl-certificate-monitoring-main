package models

type StirShakenRegionServer struct {
	Id uint
	Enabled bool
	RegionKey string
	Host string
	Country string
	CountryCode string `json:"countryCode" gorm:"column:countryCode" form:"countryCode"`
	RegionName string `json:"regionName" gorm:"column:regionName" form:"regionName"`
	City string
	Color string
	CreatedAt []uint8
	UpdatedAt []uint8
}
