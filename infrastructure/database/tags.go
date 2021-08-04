package database

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
	"rest-api-endpoints/domain/entity"
)

type TagRepo struct {
	db *sql.DB
}

func NewTagRepository(dbc *sql.DB) *TagRepo {
	if dbc == nil {
		return nil
	}

	return &TagRepo{dbc}
}

func (tg *TagRepo) FetchByProduct(c *gin.Context) (entity.TagsByProduct, error){

	result := entity.TagsByProduct{}

	rows, err := tg.db.Query("SELECT wptr.nmid,array_agg(wt.tagid) as tags FROM wb_tags as wt, wb_product_tags_relations as wptr WHERE wptr.tagid=wt.tagid AND accountid = "+ c.Query("accountid") +" GROUP BY wptr.nmid" )

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		
		var nmID int64
		var tags pq.Int64Array

		err = rows.Scan(&nmID, &tags)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, entity.ProductTags{NmID: nmID, Tags: []int64(tags)})
	}


	return result, nil
}

func (tg *TagRepo) Fetch(c *gin.Context) (entity.Tags, error) {

	result := entity.Tags{}
	//if c.Request.URL.Path == "/alltags" {
	//
	//}
	rows, err := tg.db.Query("SELECT tagid, tagname, created, updated, accountid FROM wb_tags WHERE accountid = " + c.Query("accountid"))

	if err != nil {
		return nil, err
	}


	for rows.Next() {
		tag := entity.Tag{}

		err = rows.Scan(&tag.ID, &tag.TagName, &tag.Created, &tag.Updated, &tag.AccountID)
		if err != nil {
			panic(err.Error())
		}

		result = append(result, tag)
	}

	if len(result) == 0 {
		return nil, errors.New("User not found")
	}
	return result, nil
}

func (tg *TagRepo) Create(t *entity.CreateTag) error {
	// TODO SQL to DB
	return nil
}
