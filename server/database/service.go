package database

import (
	"chatserver/json"
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Service = &service{
	url:      os.Getenv("DB_URL"),
	account:  os.Getenv("DB_ACCOUNT"),
	password: os.Getenv("DB_PASSWORD"),
	dbName:   "chat",
	db:       nil,
}

type service struct {
	url      string
	account  string
	password string
	dbName   string
	db       *mongo.Database
}

func (s *service) Start() {
	url := "mongodb+srv://" + s.account + ":" + s.password + "@" + s.url

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(url))
	if err != nil {
		log.Println("MongoDB Connection Error:", err)
		os.Exit(1)
	}
	s.db = client.Database(s.dbName)
}

func (s *service) AddUser(name, password string) bool {
	data := bson.D{
		{"name", name},
		{"password", password},
	}

	return s.insert(UserTable, data)
}

func (s *service) AddMessage(roomName, userName, text string, time int64) bool {
	data := json.Message{
		Name: userName,
		Text: text,
		Time: time,
	}

	return s.insert(roomName, data)
}

func (s *service) FindUser(name string) (User, bool) {
	filter := bson.D{
		{"name", name},
	}

	var user User
	if result := s.find(UserTable, filter); result != nil {
		if err := result.Decode(&user); err == nil {
			return user, true
		}
	}

	return user, false
}

func (s *service) FindHistory(roomName string, limit int) (history []json.Message) {
	filter := bson.D{}
	opts := options.Find().SetSort(bson.D{{"time", -1}}).SetLimit(int64(limit))

	if cursor := s.findMany(roomName, filter, opts); cursor != nil {
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var message json.Message
			err := cursor.Decode(&message)
			if err != nil {
				return nil
			}

			history = append(history, message)
		}

		return history
	}

	return nil
}

func (s *service) FindUserList(name string) (userList []string) {
	filter := bson.D{
		{"name", primitive.Regex{Pattern: name}},
	}
	opts := options.Find().SetSort(bson.D{{"name", 1}})

	if cursor := s.findMany(UserTable, filter, opts); cursor != nil {
		defer cursor.Close(context.TODO())
		for cursor.Next(context.TODO()) {
			var user User
			err := cursor.Decode(&user)
			if err != nil {
				return nil
			}

			userList = append(userList, user.Name)
		}

		return userList
	}

	return nil
}

func (s *service) AddFriend(name, friend string) bool {
	filter := bson.D{
		{"name", name},
	}
	pullData := bson.D{
		{"$pull", bson.D{{"friend", friend}}},
	}
	pushData := bson.D{
		{"$push", bson.D{{"friend", friend}}},
	}

	if s.update(UserTable, filter, pullData) {
		if s.update(UserTable, filter, pushData) {
			// exchange
			filter = bson.D{
				{"name", friend},
			}
			pullData = bson.D{
				{"$pull", bson.D{{"friend", name}}},
			}
			pushData = bson.D{
				{"$push", bson.D{{"friend", name}}},
			}
			if s.update(UserTable, filter, pullData) {
				if s.update(UserTable, filter, pushData) {
					return true
				}
			}
		}
	}

	return false
}

func (s *service) AddHint(from, to string) bool {
	filter := bson.D{
		{"name", from},
	}
	pullData := bson.D{
		{"$pull", bson.D{{"hint", to}}},
	}
	pushData := bson.D{
		{"$push", bson.D{{"hint", to}}},
	}

	if s.update(UserTable, filter, pullData) {
		if s.update(UserTable, filter, pushData) {
			return true
		}
	}

	return false
}

func (s *service) DeleteHint(from, to string) bool {
	filter := bson.D{
		{"name", from},
	}
	pullData := bson.D{
		{"$pull", bson.D{{"hint", to}}},
	}

	if s.update(UserTable, filter, pullData) {
		return true
	}

	return false
}

func (s *service) insert(table string, data interface{}) bool {
	collection := s.db.Collection(table)
	_, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("Insert Fail", table, data, err)
		return false
	}

	return true
}

func (s *service) find(table string, filter interface{}) *mongo.SingleResult {
	collection := s.db.Collection(table)
	result := collection.FindOne(context.TODO(), filter)
	if err := result.Err(); err != nil {
		log.Println("Find Many Fail", table, filter, err)
		return nil
	}

	return result
}

func (s *service) findMany(table string, filter interface{}, opts *options.FindOptions) *mongo.Cursor {
	collection := s.db.Collection(table)
	cursor, err := collection.Find(context.TODO(), filter, opts)
	if err != nil {
		log.Println("Find Many Fail", table, filter, opts, err)
		return nil
	}

	return cursor
}

func (s *service) update(table string, filter, data interface{}) bool {
	collection := s.db.Collection(table)
	_, err := collection.UpdateOne(context.TODO(), filter, data)
	if err != nil {
		log.Println("Update One Fail", table, filter, data, err)
		return false
	}

	return true
}

func (s *service) delete(table string, data interface{}) bool {
	collection := s.db.Collection(table)
	_, err := collection.DeleteOne(context.TODO(), bson.D{})
	if err != nil {
		log.Println("Delete Fail", table, data, err)
		return false
	}

	return true
}

func (s *service) aggregate(table string, matchStage, groupStage, projectStage bson.D) *mongo.Cursor {
	collection := s.db.Collection(table)
	cursor, err := collection.Aggregate(context.TODO(), mongo.Pipeline{matchStage, groupStage, projectStage})
	if err != nil {
		log.Println("Aggregate Fail", table, matchStage, groupStage, projectStage, err)
		return nil
	}

	return cursor
}
