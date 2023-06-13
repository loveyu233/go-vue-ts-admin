package dao

import (
	"database/sql"
	"log"
	"server/domain/dto"
	"server/domain/model"
	"server/domain/vo"
)

const (
	selectById                    = "select * from User where id = ?"
	verifyUsernamePassword        = "select * from User where username = ? and password = ?"
	selectUserRoutes              = "SELECT r.name FROM User u JOIN UserRoute ur ON u.id = ur.user_id JOIN Route r ON ur.route_id = r.id WHERE u.id = ?"
	selectButtons                 = "SELECT r.name FROM User u JOIN UserButton ur ON u.id = ur.user_id JOIN Button r ON ur.id = r.id WHERE u.id = ?"
	selectRoles                   = "SELECT r.name FROM User u JOIN UserRole ur ON u.id = ur.user_id JOIN Role r ON ur.id = r.id WHERE u.id = ?"
	updateUserInfo                = "update user set username=?,password=?,avatar=? where id=?"
	updateUserInfoExcludePassword = "update user set username=?,avatar=? where id=?"
)

func SelectById(id int) (*model.User, error) {
	log.Println(mysqlClient)
	var user = new(model.User)
	prepare, err := mysqlClient.Prepare(selectById)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer prepare.Close()
	row := prepare.QueryRow(id)
	err = ScanStruct(row, user)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func VerifyUsernamePassword(username, password string) (*model.User, error) {
	var user = new(model.User)
	prepare, err := mysqlClient.Prepare(verifyUsernamePassword)
	if err != nil {
		log.Println(err)
		return user, err
	}
	defer prepare.Close()
	query := prepare.QueryRow(username, password)

	if err := ScanStruct(query, user); err != nil {
		return nil, err
	}

	return user, nil
}

func SelectRoutes(uid int) ([]vo.RouteVo, error) {
	var routes = make([]vo.RouteVo, 0)
	prepare, err := mysqlClient.Prepare(selectUserRoutes)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	rows, err := prepare.Query(uid)
	if err != nil {
		return nil, err
	}
	err = ScanRows(rows, &routes)
	if err != nil {
		return nil, err
	}
	return routes, nil
}

func SelectRoles(uid int) ([]vo.RoleVo, error) {
	var roles = make([]vo.RoleVo, 0)
	prepare, err := mysqlClient.Prepare(selectRoles)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	rows, err := prepare.Query(uid)
	if err != nil {
		return nil, err
	}
	err = ScanRows(rows, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func SelectButtons(uid int) ([]vo.ButtonVo, error) {
	var buttons = make([]vo.ButtonVo, 0)
	prepare, err := mysqlClient.Prepare(selectButtons)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	rows, err := prepare.Query(uid)
	if err != nil {
		return nil, err
	}
	err = ScanRows(rows, &buttons)
	if err != nil {
		return nil, err
	}
	return buttons, nil
}

func UpdateUserInfo(user dto.UserUpdateDto) (int64, error) {
	var err error
	var prepare *sql.Stmt
	var exec sql.Result
	if user.Password == "" {
		prepare, err = mysqlClient.Prepare(updateUserInfoExcludePassword)
		if err != nil {
			return 0, err
		}
		exec, err = prepare.Exec(user.Username, user.Icon, user.Uid)
	} else {
		prepare, err = mysqlClient.Prepare(updateUserInfo)
		if err != nil {
			return 0, err
		}
		exec, err = prepare.Exec(user.Username, user.Password, user.Icon, user.Uid)
	}
	if err != nil {
		return 0, err
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		return 0, err
	}
	return affected, nil
}
