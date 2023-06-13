package dao

import (
	"database/sql"
	"fmt"
	"reflect"
)

var (
	ErrScanFailed = fmt.Errorf("scan failed")
)

// ScanStruct 将数据库查询结果扫描和填充到指定的结构体中
//
// 参数：
//   - row: *sql.Row 对象，表示一行数据库查询结果
//   - dest: 目标结构体的指针，用于接收查询结果
//
// 返回值：
//   - error: 如果扫描和填充过程中发生错误，返回对应的错误信息；如果查询结果为空，返回 sql.ErrNoRows 错误
//
// 注意事项：
//   - 目标结构体的字段必须具有 `db` 标签，用于指定对应的列名
func ScanStruct(row *sql.Row, dest interface{}) error {
	// 检查目标参数是否为指针类型并非空
	destValue := reflect.ValueOf(dest)
	if destValue.Kind() != reflect.Ptr || destValue.IsNil() {
		return fmt.Errorf("invalid destination")
	}

	// 获取目标结构体的类型信息和字段个数
	destType := destValue.Elem().Type()
	numFields := destType.NumField()

	// 创建列名和值的切片
	columns := make([]string, numFields)
	values := make([]interface{}, numFields)

	// 遍历目标结构体的字段，获取列名和值的指针
	for i := 0; i < numFields; i++ {
		field := destType.Field(i)
		columns[i] = field.Tag.Get("db") // 获取字段的 db 标签值

		// 创建值的指针，并将其存储在 values 切片中
		valuePtr := reflect.New(field.Type)
		values[i] = valuePtr.Interface()
	}

	// 执行查询结果的扫描
	err := row.Scan(values...)
	if err != nil {
		// 处理扫描过程中的错误
		if err == sql.ErrNoRows {
			return err // 返回 sql.ErrNoRows 错误，表示查询结果为空
		}
		return fmt.Errorf("%w: %v", ErrScanFailed, err) // 返回自定义错误，表示扫描失败
	}

	// 将扫描结果填充到目标结构体的字段中
	for i := 0; i < numFields; i++ {
		// 获取值的指针，并将其赋值给目标结构体的字段
		valuePtr := reflect.ValueOf(values[i])
		destValue.Elem().Field(i).Set(valuePtr.Elem())
	}

	return nil
}

func ScanRows(rows *sql.Rows, dest interface{}) error {
	// 获取传入结构体数组的类型信息
	destType := reflect.TypeOf(dest)
	if destType.Kind() != reflect.Ptr || destType.Elem().Kind() != reflect.Slice {
		return ErrScanFailed
	}

	// 获取结构体类型
	structType := destType.Elem().Elem()

	// 创建一个新的切片作为结果集
	result := reflect.MakeSlice(destType.Elem(), 0, 0)

	// 遍历查询结果
	for rows.Next() {
		// 创建一个新的结构体实例
		structValue := reflect.New(structType).Elem()

		// 获取结构体字段的数量
		numFields := structType.NumField()

		// 创建一个字段值切片，用于传递给 rows.Scan
		fieldValues := make([]interface{}, numFields)

		// 遍历结构体字段
		for i := 0; i < numFields; i++ {
			// 获取字段值的指针
			fieldValues[i] = structValue.Field(i).Addr().Interface()
		}

		// 使用 rows.Scan 将查询结果扫描到结构体字段中
		err := rows.Scan(fieldValues...)
		if err != nil {
			return err
		}

		// 将结构体添加到结果集中
		result = reflect.Append(result, structValue)
	}

	// 将结果集赋值给传入的结构体数组
	reflect.ValueOf(dest).Elem().Set(result)

	return nil
}

func createNewInstance(input interface{}) interface{} {
	// 获取传入结构体的类型信息
	inputType := reflect.TypeOf(input)

	// 创建对应类型的新实例
	newInstance := reflect.New(inputType).Elem()

	// 返回新实例的值
	return newInstance.Interface()
}
