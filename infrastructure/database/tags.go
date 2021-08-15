package database

import (
	"database/sql"
	"errors"
	"rest-api-endpoints/domain/entity"

	"github.com/lib/pq"
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

func (tg *TagRepo) FetchByProduct(accountID int64) (entity.TagsByProduct, error) {

	result := entity.TagsByProduct{}

	rows, err := tg.db.Query("SELECT wptr.nmid,array_agg(wt.tagid) as tags FROM tags as wt, product_tags_relations as wptr WHERE wptr.tagid=wt.tagid AND accountid=$1  GROUP BY wptr.nmid", accountID)

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

func (tg *TagRepo) Fetch(accountID int64) (entity.Tags, error) {

	result := entity.Tags{}
	rows, err := tg.db.Query("SELECT tagid, tagname, created, updated, accountid FROM tags WHERE accountid = $1", accountID)

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
		return nil, errors.New("user not found")
	}
	return result, nil
}
func (tg *TagRepo) Update(t *entity.CreateTag) error {

	sSql := "UPDATE tags SET tagname=$1, updated=NOW() WHERE tagid=$2 AND accountid=$3"
	res, err := tg.db.Exec(sSql, t.TagName, t.ID, t.AccountID)

	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	if t.ProductID != 0 {
		// проверка на принадлежность тега аккаунту происходит выше, тут можно не заморачиваться
		// проверка на принадлежность nmid аккаунту остается на совести фронта
		sSql := "INSERT INTO product_tags_relations(tagid, nmid) VALUES($1, $2) ON CONFLICT DO NOTHING"
		_, err := tg.db.Exec(sSql, t.ID, t.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (tg *TagRepo) Create(t *entity.CreateTag) error {

	var lastInsertId int64
	sSql := "INSERT INTO tags(tagname, accountid) VALUES($1, $2) ON CONFLICT DO NOTHING RETURNING tagid;"
	err := tg.db.QueryRow(sSql, t.TagName, t.AccountID).Scan(&lastInsertId)

	if err != nil {
		return err
	}

	if lastInsertId == 0 {
		return errors.New("failed to create tag")
	}

	t.ID = lastInsertId

	if t.ProductID != 0 {
		// проверка на принадлежность nmid аккаунту остается на совести фронта
		sSql := "INSERT INTO product_tags_relations(tagid, nmid) VALUES($1, $2) ON CONFLICT DO NOTHING"
		_, err := tg.db.Exec(sSql, t.ID, t.ProductID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (tg *TagRepo) Delete(t *entity.CreateTag) error {

	if t.ProductID == 0 {
		sSql := "DELETE FROM tags WHERE tagid=$2 AND accountid=$1"
		_, err := tg.db.Exec(sSql, t.AccountID, t.ID)
		if err != nil {
			return err
		}
	}

	// Если ProductID указан удаояем только связку с товаром
	sSql := "DELETE FROM product_tags_relations WHERE tagid=$2 AND nmid=$1"
	_, err := tg.db.Exec(sSql, t.ProductID, t.ID)
	if err != nil {
		return err
	}

	return nil
}
