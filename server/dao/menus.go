package dao

import (
	"log"
	"server/cache"
	"server/domain/dto"
	"server/domain/model"
	"sort"
)

const (
	getMenusInfo = "SELECT * FROM sys_base_menus ORDER BY r_id"
	Limit        = "select * from sys_base_menus limit ?,?"
	Count        = "select count(*) from sys_base_menus"
	Update       = "update sys_base_menus set title=?, image=? where id=?;"
)

func GetMenusInfo() ([]*model.SysBaseMenus, error) {
	prepare, err := mysqlClient.Prepare(getMenusInfo)
	if err != nil {
		return nil, err
	}
	defer prepare.Close()
	rows, err := prepare.Query()
	if err != nil {
		return nil, err
	}
	var rootItems []*model.SysBaseMenus

	for rows.Next() {
		var menuItem model.SysBaseMenus
		err := rows.Scan(
			&menuItem.ID,
			&menuItem.RID,
			&menuItem.MenuLevel,
			&menuItem.ParentID,
			&menuItem.Path,
			&menuItem.Name,
			&menuItem.Hidden,
			&menuItem.Component,
			&menuItem.Sort,
			&menuItem.Title,
			&menuItem.Icon,
		)
		if err != nil {
			log.Fatal(err)
		}
		rootItems = append(rootItems, &menuItem)
	}
	cache.SaveMenu(rootItems)
	var m = make([]*model.SysBaseMenus, 0)
	for i := range rootItems {
		m = list(rootItems[i], m)
	}
	defer func() {
		m = nil
	}()
	return m, nil
}

func list(item *model.SysBaseMenus, res []*model.SysBaseMenus) []*model.SysBaseMenus {
	sort.Slice(res, func(i, j int) bool {
		if res[i].Sort < res[j].Sort {
			return true
		}
		return false
	})
	if item.ParentID == 0 {
		res = append(res, item)
	} else {
		for i := range res {
			if res[i].RID == item.ParentID {
				if res[i].Children == nil {
					res[i].Children = make([]*model.SysBaseMenus, 0)
				}
				res[i].Children = append(res[i].Children, item)
				break
			} else {
				list(item, res[i].Children)
			}
		}
	}
	return res
}

func GetMenusLimit(page, pageSize int) ([]model.SysBaseMenus, int64) {
	prepare, err := mysqlClient.Prepare(Limit)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	rows, err := prepare.Query((page-1)*pageSize, pageSize)
	if err != nil {
		log.Println(err)
		return nil, -1
	}
	rootItems := make([]model.SysBaseMenus, 0)
	for rows.Next() {
		var menuItem model.SysBaseMenus
		err := rows.Scan(
			&menuItem.ID,
			&menuItem.RID,
			&menuItem.MenuLevel,
			&menuItem.ParentID,
			&menuItem.Path,
			&menuItem.Name,
			&menuItem.Hidden,
			&menuItem.Component,
			&menuItem.Sort,
			&menuItem.Title,
			&menuItem.Icon,
		)
		if err != nil {
			log.Fatal(err)
		}
		rootItems = append(rootItems, menuItem)
	}

	return rootItems, GetMenusTotal()
}

func GetMenusTotal() int64 {
	stmt, err := mysqlClient.Prepare(Count)
	if err != nil {
		return -1
	}
	row := stmt.QueryRow()
	var total int64
	row.Scan(&total)
	return total
}

func MenuUpdate(menu dto.MenuUpdate) int64 {
	prepare, err := mysqlClient.Prepare(Update)
	if err != nil {
		log.Println(err)
		return -1
	}
	exec, err := prepare.Exec(menu.Title, menu.Icon, menu.Id)
	if err != nil {
		log.Println(err)
		return -1
	}
	affected, err := exec.RowsAffected()
	if err != nil {
		log.Println(err)
		return -1
	}
	return affected
}
