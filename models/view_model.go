package models

import (
	"connect_db_exc/entities"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type ViewModel struct {
	DB         *mgo.Database
	Collection string
}

func (viewModel ViewModel) FindAllViews() (views []entities.View, err error) {
	err = viewModel.DB.C(viewModel.Collection).Pipe([]bson.M{
		{"$match": bson.M{"date": "11-08-2021"}},
		{"$group": bson.M{"_id": "$userID", "total": bson.M{"$sum": "$count"}}},
	}).All(&views)
	return
}

func (viewModel ViewModel) FindView(id string) (view entities.View, err error) {
	err = viewModel.DB.C(viewModel.Collection).FindId(bson.ObjectIdHex(id)).One(&view)
	return
}

func (viewModel ViewModel) CreateView(view *entities.View) (err error) {
	err = viewModel.DB.C(viewModel.Collection).Insert(&view)
	return
}

func (viewModel ViewModel) UpdateView(view entities.View) (err error) {
	err = viewModel.DB.C(viewModel.Collection).UpdateId(view.ID, &view)
	return
}

/*func (viewModel ViewModel) UpdateView(id bson.ObjectId, view *entities.View, clickCount int) (err error) {
	err = viewModel.DB.C(viewModel.Collection).Update(
		bson.M{
			"_id":    id,
			"UserID": view.UserID,
			"Date":   view.Date,
		},
		bson.M{
			"$inc": bson.M{
				"count": clickCount,
			},
		})
	return
}*/
