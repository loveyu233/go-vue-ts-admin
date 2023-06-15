package dao

import (
	"log"
	"server/domain/model"
	"strings"
)

const (
	getCarouselsSql       = "select * from carousel"
	getShowCarouselSql    = "select * from carousel where is_show = 1"
	addCarouselSql        = "insert into carousel(url, is_show) values (?,?)"
	delCarouselSql        = "delete from carousel where id = ?"
	UpdateCarouselSql     = "update carousel set url = ?, is_show = ? where id = ?"
	UpdateCarouselShowSql = "update carousel set is_show = ? where id = ?"
	DeleteCarouselInSql   = "delete from carousel where id in  (?)"
)

func GetCarousels() ([]*model.Carousel, error) {
	prepare, err := mysqlClient.Prepare(getCarouselsSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query, err := prepare.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var cs []*model.Carousel
	for query.Next() {
		c := model.Carousel{}
		err := query.Scan(&c.ID, &c.Url, &c.IsShow)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cs = append(cs, &c)
	}
	return cs, nil
}

func GetShowCarousels() ([]*model.Carousel, error) {
	prepare, err := mysqlClient.Prepare(getShowCarouselSql)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	query, err := prepare.Query()
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var cs []*model.Carousel
	for query.Next() {
		c := model.Carousel{}
		err := query.Scan(&c.ID, &c.Url, &c.IsShow)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cs = append(cs, &c)
	}
	return cs, nil
}

func AddCarousel(carousel model.Carousel) (int64, error) {
	prepare, err := mysqlClient.Prepare(addCarouselSql)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	exec, err := prepare.Exec(carousel.Url, carousel.IsShow)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return exec.RowsAffected()
}

func DeleteCarousel(id string) (int64, error) {
	prepare, err := mysqlClient.Prepare(delCarouselSql)
	if err != nil {
		if err != nil {
			log.Println(err)
			return 0, err
		}
	}
	exec, err := prepare.Exec(id)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return exec.RowsAffected()
}

func UpdateCarousel(carousel model.Carousel) (int64, error) {
	exec, err := mysqlClient.Exec(UpdateCarouselSql, carousel.Url, carousel.IsShow, carousel.ID)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return exec.RowsAffected()
}

// TODO DeleteCarouselIn
func DeleteCarouselIn(ids []string) (int64, error) {
	exec, err := mysqlClient.Exec(DeleteCarouselInSql, strings.Join(ids, ","))
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return exec.RowsAffected()
}

func UpdateCarouselShow(op int, id string) (int64, error) {
	prepare, err := mysqlClient.Prepare(UpdateCarouselShowSql)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	exec, err := prepare.Exec(op, id)
	if err != nil {
		return 0, err
	}
	return exec.RowsAffected()
}
