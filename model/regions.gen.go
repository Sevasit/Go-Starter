// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameRegion = "regions"

// Region mapped from table <regions>
type Region struct {
	RegionID   int32  `gorm:"column:region_id;primaryKey;autoIncrement:true" json:"region_id"`
	RegionName string `gorm:"column:region_name" json:"region_name"`
}

// TableName Region's table name
func (*Region) TableName() string {
	return TableNameRegion
}
