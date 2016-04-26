package boil

import (
	"bytes"
	"database/sql"
	"fmt"
	"strings"
)

type where struct {
	clause string
	args   []interface{}
}

type join struct {
	on   string
	args []interface{}
}

type Query struct {
	executor        Executor
	delete          bool
	update          map[string]interface{}
	selectCols      []string
	table           string
	innerJoins      []join
	outerJoins      []join
	leftOuterJoins  []join
	rightOuterJoins []join
	where           []where
	groupBy         []string
	orderBy         []string
	having          []string
	limit           int
}

func buildQuery(q *Query) (string, []interface{}) {
	var buf *bytes.Buffer
	var args []interface{}

	switch {
	case q.delete:
		buf, args = buildDeleteQuery(q)
	case len(q.update) > 0:
		buf, args = buildUpdateQuery(q)
	default:
		buf, args = buildSelectQuery(q)
	}

	return buf.String(), args
}

func buildSelectQuery(q *Query) (*bytes.Buffer, []interface{}) {
	buf := &bytes.Buffer{}

	buf.WriteString("SELECT ")

	if len(q.selectCols) > 0 {
		buf.WriteString(strings.Join(q.selectCols, ","))
	} else {
		buf.WriteByte('*')
	}

	buf.WriteString(" FROM ")
	fmt.Fprintf(buf, `"%s"`, q.table)

	return buf, []interface{}{}
}

func buildDeleteQuery(q *Query) (*bytes.Buffer, []interface{}) {
	buf := &bytes.Buffer{}

	return buf, nil
}

func buildUpdateQuery(q *Query) (*bytes.Buffer, []interface{}) {
	buf := &bytes.Buffer{}

	return buf, nil
}

// ExecQuery executes a query that does not need a row returned
func ExecQuery(q *Query) (sql.Result, error) {
	qs, args := buildQuery(q)
	return q.executor.Exec(qs, args...)
}

// ExecQueryOne executes the query for the One finisher and returns a row
func ExecQueryOne(q *Query) *sql.Row {
	qs, args := buildQuery(q)
	return q.executor.QueryRow(qs, args)
}

// ExecQueryAll executes the query for the All finisher and returns multiple rows
func ExecQueryAll(q *Query) (*sql.Rows, error) {
	qs, args := buildQuery(q)
	return q.executor.Query(qs, args)
}

func SetDelete(q *Query) {
	q.delete = true
}

func SetUpdate(q *Query, cols map[string]interface{}) {
	q.update = cols
}

func SetExecutor(q *Query, exec Executor) {
	q.executor = exec
}

func SetSelect(q *Query, columns ...string) {
	q.selectCols = append(q.selectCols, columns...)
}

func SetTable(q *Query, table string) {
	q.table = table
}

func SetInnerJoin(q *Query, on string, args ...interface{}) {
	q.innerJoins = append(q.innerJoins, join{on: on, args: args})
}

func SetOuterJoin(q *Query, on string, args ...interface{}) {
	q.outerJoins = append(q.outerJoins, join{on: on, args: args})
}

func SetLeftOuterJoin(q *Query, on string, args ...interface{}) {
	q.leftOuterJoins = append(q.leftOuterJoins, join{on: on, args: args})
}

func SetRightOuterJoin(q *Query, on string, args ...interface{}) {
	q.rightOuterJoins = append(q.rightOuterJoins, join{on: on, args: args})
}

func SetWhere(q *Query, clause string, args ...interface{}) {
	q.where = append(q.where, where{clause: clause, args: args})
}

func SetGroupBy(q *Query, clause string) {
	q.groupBy = append(q.groupBy, clause)
}

func SetOrderBy(q *Query, clause string) {
	q.orderBy = append(q.orderBy, clause)
}

func SetHaving(q *Query, clause string) {
	q.having = append(q.having, clause)
}

func SetLimit(q *Query, limit int) {
	q.limit = limit
}
