package postgres

import (
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"reflect"
)

// DB - интерфейс для обёртки
type DB interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
}

// Wrapper - структура, которая содержит пул соединений с базой данных
type Wrapper struct {
	pool *pgxpool.Pool
}

// NewWrapper - функция для создания новой обёртки с пулом соединений
func NewWrapper(pool *pgxpool.Pool) *Wrapper {
	return &Wrapper{pool: pool}
}

// QueryRow - обёртка для метода QueryRow
func (w *Wrapper) QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row {
	return w.pool.QueryRow(ctx, sql, args...)
}

// Query - обёртка для метода Query
func (w *Wrapper) Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error) {
	return w.pool.Query(ctx, sql, args...)
}

// Exec - обёртка для метода Exec
func (w *Wrapper) Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error) {
	return w.pool.Exec(ctx, sql, args...)
}

// Get - функция для выполнения запроса, который возвращает одну строку, и сканирования ее в переданную структуру
func (w *Wrapper) Get(ctx context.Context, dest interface{}, sql string, args ...interface{}) error {
	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Struct {
		return errors.New("dest must be a pointer to a struct")
	}

	destElem := destVal.Elem()
	numFields := destElem.NumField()
	scanArgs := make([]interface{}, numFields)

	for i := 0; i < numFields; i++ {
		field := destElem.Field(i)
		scanArgs[i] = field.Addr().Interface()
	}

	row := w.pool.QueryRow(ctx, sql, args...)

	return row.Scan(scanArgs...)
}

// Select - функция для получения нескольких результатов и их сканирования в слайс
func (w *Wrapper) Select(ctx context.Context, dest interface{}, sqlStr string, args ...interface{}) error {
	rows, err := w.pool.Query(ctx, sqlStr, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	destVal := reflect.ValueOf(dest)
	if destVal.Kind() != reflect.Ptr || destVal.Elem().Kind() != reflect.Slice {
		return errors.New("dest must be a pointer to a slice")
	}

	sliceVal := destVal.Elem()
	elemType := sliceVal.Type().Elem()

	ptrToStruct := false
	if elemType.Kind() == reflect.Ptr && elemType.Elem().Kind() == reflect.Struct {
		ptrToStruct = true
		elemType = elemType.Elem()
	} else if elemType.Kind() != reflect.Struct {
		return errors.New("slice elements must be structs or pointers to structs")
	}

	fieldsDescriptions := rows.FieldDescriptions()
	columns := make([]string, len(fieldsDescriptions))
	for i, fd := range fieldsDescriptions {
		columns[i] = string(fd.Name)
	}

	for rows.Next() {

		elemPtr := reflect.New(elemType)

		fields, err := structFieldsPointers(elemPtr.Interface(), columns)
		if err != nil {
			return err
		}

		if err := rows.Scan(fields...); err != nil {
			return err
		}

		if ptrToStruct {
			sliceVal.Set(reflect.Append(sliceVal, elemPtr))
		} else {
			sliceVal.Set(reflect.Append(sliceVal, elemPtr.Elem()))
		}
	}

	if err := rows.Err(); err != nil {
		return err
	}

	return nil
}

// Вспомогательная функция для создания слайса структур
func structFieldsPointers(strct interface{}, columns []string) ([]interface{}, error) {
	v := reflect.ValueOf(strct)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	if v.Kind() != reflect.Struct {
		return nil, errors.New("input must be a struct or pointer to struct")
	}

	t := v.Type()

	fieldMap := make(map[string]int)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		tag := field.Tag.Get("db")
		if tag == "" {
			tag = field.Name
		}
		fieldMap[tag] = i
	}

	fields := make([]interface{}, len(columns))
	for i, col := range columns {
		idx, ok := fieldMap[col]
		if !ok {
			return nil, fmt.Errorf("no matching struct field found for column %s", col)
		}
		fields[i] = v.Field(idx).Addr().Interface()
	}

	return fields, nil
}
