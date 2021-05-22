package dao_mongo

import (
	"context"
	"gin-20210518/model"

	"github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/bson"
)

type studentCollection struct {
	name string
}

var StudentColl = studentCollection{
	name: "student",
}

func (c *studentCollection) FindAll() ([]*model.Student, error) {
	result := make([]*model.Student, 0)
	filter := bson.M{}
	curs, err := MongoDBDataBase.Collection(c.name).Find(context.TODO(), filter)
	if err != nil {
		logrus.Errorf("fail to find student err:%v", err)
	}
	err = curs.All(context.TODO(), &result)
	return result, err
}

func (c *studentCollection) Find(id int64) (*model.Student, error) {
	result := new(model.Student)
	filter := bson.M{"id": id}
	err := MongoDBDataBase.Collection(c.name).FindOne(nil, filter).Decode(result)
	if err != nil {
		logrus.Errorf("find student by id:%v err:%v", id, err)
	}
	return result, err
}

func (c *studentCollection) Add(stu *model.Student) error {
	result, err := MongoDBDataBase.Collection(c.name).InsertOne(context.TODO(), stu)
	logrus.Infof("insert result:%v", result.InsertedID)
	return err
}

func (c *studentCollection) Delete(id int64) error {
	filter := bson.M{"id": id}
	one, err := MongoDBDataBase.Collection(c.name).DeleteOne(context.TODO(), filter)
	if err != nil {
		logrus.Errorf("fail to delete student id:%v err:%v", id, err)
	}
	logrus.Infof("delete data:%v", one.DeletedCount)
	return err
}
