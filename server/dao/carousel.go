package dao

import (
	"log"
	"server/domain/model"
)

const (
	getCarousels = "select * from carousel"
)

func GetCarousels() ([]*model.Carousel, error) {
	prepare, err := mysqlClient.Prepare(getCarousels)
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
		err := query.Scan(&c.ID, &c.Url)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		cs = append(cs, &c)
	}
	return cs, nil
}
