// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/Sevasit/go-crud/model"
)

func newCountry(db *gorm.DB, opts ...gen.DOOption) country {
	_country := country{}

	_country.countryDo.UseDB(db, opts...)
	_country.countryDo.UseModel(&model.Country{})

	tableName := _country.countryDo.TableName()
	_country.ALL = field.NewAsterisk(tableName)
	_country.CountryID = field.NewInt32(tableName, "country_id")
	_country.CountryName = field.NewString(tableName, "country_name")
	_country.RegionID = field.NewInt32(tableName, "region_id")

	_country.fillFieldMap()

	return _country
}

type country struct {
	countryDo

	ALL         field.Asterisk
	CountryID   field.Int32
	CountryName field.String
	RegionID    field.Int32

	fieldMap map[string]field.Expr
}

func (c country) Table(newTableName string) *country {
	c.countryDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c country) As(alias string) *country {
	c.countryDo.DO = *(c.countryDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *country) updateTableName(table string) *country {
	c.ALL = field.NewAsterisk(table)
	c.CountryID = field.NewInt32(table, "country_id")
	c.CountryName = field.NewString(table, "country_name")
	c.RegionID = field.NewInt32(table, "region_id")

	c.fillFieldMap()

	return c
}

func (c *country) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *country) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 3)
	c.fieldMap["country_id"] = c.CountryID
	c.fieldMap["country_name"] = c.CountryName
	c.fieldMap["region_id"] = c.RegionID
}

func (c country) clone(db *gorm.DB) country {
	c.countryDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c country) replaceDB(db *gorm.DB) country {
	c.countryDo.ReplaceDB(db)
	return c
}

type countryDo struct{ gen.DO }

type ICountryDo interface {
	gen.SubQuery
	Debug() ICountryDo
	WithContext(ctx context.Context) ICountryDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICountryDo
	WriteDB() ICountryDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICountryDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICountryDo
	Not(conds ...gen.Condition) ICountryDo
	Or(conds ...gen.Condition) ICountryDo
	Select(conds ...field.Expr) ICountryDo
	Where(conds ...gen.Condition) ICountryDo
	Order(conds ...field.Expr) ICountryDo
	Distinct(cols ...field.Expr) ICountryDo
	Omit(cols ...field.Expr) ICountryDo
	Join(table schema.Tabler, on ...field.Expr) ICountryDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICountryDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICountryDo
	Group(cols ...field.Expr) ICountryDo
	Having(conds ...gen.Condition) ICountryDo
	Limit(limit int) ICountryDo
	Offset(offset int) ICountryDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICountryDo
	Unscoped() ICountryDo
	Create(values ...*model.Country) error
	CreateInBatches(values []*model.Country, batchSize int) error
	Save(values ...*model.Country) error
	First() (*model.Country, error)
	Take() (*model.Country, error)
	Last() (*model.Country, error)
	Find() ([]*model.Country, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Country, err error)
	FindInBatches(result *[]*model.Country, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.Country) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICountryDo
	Assign(attrs ...field.AssignExpr) ICountryDo
	Joins(fields ...field.RelationField) ICountryDo
	Preload(fields ...field.RelationField) ICountryDo
	FirstOrInit() (*model.Country, error)
	FirstOrCreate() (*model.Country, error)
	FindByPage(offset int, limit int) (result []*model.Country, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICountryDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c countryDo) Debug() ICountryDo {
	return c.withDO(c.DO.Debug())
}

func (c countryDo) WithContext(ctx context.Context) ICountryDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c countryDo) ReadDB() ICountryDo {
	return c.Clauses(dbresolver.Read)
}

func (c countryDo) WriteDB() ICountryDo {
	return c.Clauses(dbresolver.Write)
}

func (c countryDo) Session(config *gorm.Session) ICountryDo {
	return c.withDO(c.DO.Session(config))
}

func (c countryDo) Clauses(conds ...clause.Expression) ICountryDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c countryDo) Returning(value interface{}, columns ...string) ICountryDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c countryDo) Not(conds ...gen.Condition) ICountryDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c countryDo) Or(conds ...gen.Condition) ICountryDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c countryDo) Select(conds ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c countryDo) Where(conds ...gen.Condition) ICountryDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c countryDo) Order(conds ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c countryDo) Distinct(cols ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c countryDo) Omit(cols ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c countryDo) Join(table schema.Tabler, on ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c countryDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICountryDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c countryDo) RightJoin(table schema.Tabler, on ...field.Expr) ICountryDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c countryDo) Group(cols ...field.Expr) ICountryDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c countryDo) Having(conds ...gen.Condition) ICountryDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c countryDo) Limit(limit int) ICountryDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c countryDo) Offset(offset int) ICountryDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c countryDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICountryDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c countryDo) Unscoped() ICountryDo {
	return c.withDO(c.DO.Unscoped())
}

func (c countryDo) Create(values ...*model.Country) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c countryDo) CreateInBatches(values []*model.Country, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c countryDo) Save(values ...*model.Country) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c countryDo) First() (*model.Country, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.Country), nil
	}
}

func (c countryDo) Take() (*model.Country, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.Country), nil
	}
}

func (c countryDo) Last() (*model.Country, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.Country), nil
	}
}

func (c countryDo) Find() ([]*model.Country, error) {
	result, err := c.DO.Find()
	return result.([]*model.Country), err
}

func (c countryDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.Country, err error) {
	buf := make([]*model.Country, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c countryDo) FindInBatches(result *[]*model.Country, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c countryDo) Attrs(attrs ...field.AssignExpr) ICountryDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c countryDo) Assign(attrs ...field.AssignExpr) ICountryDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c countryDo) Joins(fields ...field.RelationField) ICountryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c countryDo) Preload(fields ...field.RelationField) ICountryDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c countryDo) FirstOrInit() (*model.Country, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.Country), nil
	}
}

func (c countryDo) FirstOrCreate() (*model.Country, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.Country), nil
	}
}

func (c countryDo) FindByPage(offset int, limit int) (result []*model.Country, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c countryDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c countryDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c countryDo) Delete(models ...*model.Country) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *countryDo) withDO(do gen.Dao) *countryDo {
	c.DO = *do.(*gen.DO)
	return c
}
