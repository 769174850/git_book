// package dao 定义了数据库操作的函数
package dao

import (
	"becourse7/models"
	"becourse7/util"
	"log"
)

// GetMessages 获取所有留言
func GetMessages() ([]models.Message, error) {
	// 定义一个切片，用于存储查询到的留言
	var messages []models.Message
	// 定义一个SQL语句，用于查询未删除的留言，按照创建时间降序排列
	sql := "SELECT id, detail, create_at, update_at FROM message WHERE is_deleted IS FALSE ORDER BY create_at DESC"
	// 执行SQL语句，返回一个*sql.Rows对象，用于遍历查询结果
	rows, err := DB.Query(sql)
	if err != nil {
		// 返回错误
		log.Println(err)
		return nil, err
	}
	// 关闭*sql.Rows对象，释放资源
	defer rows.Close()
	// 遍历查询结果，将每一条留言追加到切片中
	for rows.Next() {
		// 定义一个Message变量，用于存储一条留言
		var m models.Message
		// 将查询结果的每一列的值扫描到Message变量的对应字段中
		err := rows.Scan(&m.ID, &m.Detail, &m.CreateAt, &m.UpDateAt)
		if err != nil {
			// 返回错误
			log.Println(err)
			return nil, err
		}
		// 将Message变量追加到切片中
		messages = append(messages, m)
	}
	// 检查遍历过程中是否有错误
	err = rows.Err()
	if err != nil {
		// 返回错误
		log.Println(err)
		return nil, err
	}
	// 返回留言切片和nil错误
	return messages, nil
}

// AddMessage 添加一条留言
func AddMessage(m models.Message) error {
	// 定义一个SQL语句，用于插入一条留言
	sql := "INSERT INTO message (detail) VALUES (?)"
	// 执行SQL语句，返回一个*sql.Result对象，用于获取执行结果
	result, err := DB.Exec(sql, m.Detail)
	if err != nil {
		// 返回错误s
		log.Println(err)
		return err
	}
	// 获取插入的留言的ID
	id, err := result.LastInsertId()
	if err != nil {
		// 返回错误
		return err
	}
	// 将插入的留言的ID赋值给Message结构体的ID字段
	m.ID = id
	// 返回nil错误
	return nil
}

// DeleteMessage 删除一条留言
func DeleteMessage(id int64) error {
	// 定义一个SQL语句，用于更新一条留言的删除时间
	sql := "UPDATE message SET is_deleted = ? WHERE id = ?"
	// 执行SQL语句，返回一个*sql.Result对象，用于获取执行结果
	result, err := DB.Exec(sql, true, id)
	if err != nil {
		// 返回错误
		return err
	}
	// 获取影响的行数
	rows, err := result.RowsAffected()
	if err != nil {
		// 返回错误
		return err
	}
	// 判断影响的行数是否为0，如果为0，表示没有找到要删除的留言
	if rows == 0 {
		// 返回一个自定义的错误，表示资源不存在
		return util.NoRecordExistError
	}
	// 返回nil错误
	return nil
}
