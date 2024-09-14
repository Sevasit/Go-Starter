// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package entity

import (
	"context"
	"database/sql"

	"gorm.io/gorm"

	"gorm.io/gen"

	"gorm.io/plugin/dbresolver"
)

var (
	Q          = new(Query)
	Country    *country
	Department *department
	Employee   *employee
	Region     *region
)

func SetDefault(db *gorm.DB, opts ...gen.DOOption) {
	*Q = *Use(db, opts...)
	Country = &Q.Country
	Department = &Q.Department
	Employee = &Q.Employee
	Region = &Q.Region
}

func Use(db *gorm.DB, opts ...gen.DOOption) *Query {
	return &Query{
		db:         db,
		Country:    newCountry(db, opts...),
		Department: newDepartment(db, opts...),
		Employee:   newEmployee(db, opts...),
		Region:     newRegion(db, opts...),
	}
}

type Query struct {
	db *gorm.DB

	Country    country
	Department department
	Employee   employee
	Region     region
}

func (q *Query) Available() bool { return q.db != nil }

func (q *Query) clone(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Country:    q.Country.clone(db),
		Department: q.Department.clone(db),
		Employee:   q.Employee.clone(db),
		Region:     q.Region.clone(db),
	}
}

func (q *Query) ReadDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Read))
}

func (q *Query) WriteDB() *Query {
	return q.ReplaceDB(q.db.Clauses(dbresolver.Write))
}

func (q *Query) ReplaceDB(db *gorm.DB) *Query {
	return &Query{
		db:         db,
		Country:    q.Country.replaceDB(db),
		Department: q.Department.replaceDB(db),
		Employee:   q.Employee.replaceDB(db),
		Region:     q.Region.replaceDB(db),
	}
}

type queryCtx struct {
	Country    ICountryDo
	Department IDepartmentDo
	Employee   IEmployeeDo
	Region     IRegionDo
}

func (q *Query) WithContext(ctx context.Context) *queryCtx {
	return &queryCtx{
		Country:    q.Country.WithContext(ctx),
		Department: q.Department.WithContext(ctx),
		Employee:   q.Employee.WithContext(ctx),
		Region:     q.Region.WithContext(ctx),
	}
}

func (q *Query) Transaction(fc func(tx *Query) error, opts ...*sql.TxOptions) error {
	return q.db.Transaction(func(tx *gorm.DB) error { return fc(q.clone(tx)) }, opts...)
}

func (q *Query) Begin(opts ...*sql.TxOptions) *QueryTx {
	tx := q.db.Begin(opts...)
	return &QueryTx{Query: q.clone(tx), Error: tx.Error}
}

type QueryTx struct {
	*Query
	Error error
}

func (q *QueryTx) Commit() error {
	return q.db.Commit().Error
}

func (q *QueryTx) Rollback() error {
	return q.db.Rollback().Error
}

func (q *QueryTx) SavePoint(name string) error {
	return q.db.SavePoint(name).Error
}

func (q *QueryTx) RollbackTo(name string) error {
	return q.db.RollbackTo(name).Error
}
