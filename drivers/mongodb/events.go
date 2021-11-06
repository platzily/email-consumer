package mongodb

import (
	"context"
	"time"

	"github.com/platzily/consumer/config"
	"github.com/platzily/consumer/domains"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const EVENTS_COLLECTION = "events"

var env *config.MongoDBConfig = config.ReadMongoDBConfig()

type EventRepository struct {
	conn *mongo.Client
}

func New(db *mongo.Client) domains.EventModel {

	return &EventRepository{
		conn: db,
	}
}

func (er *EventRepository) GetById(id int64) (domains.Event, error) {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	var result domains.Event
	err := collection.FindOne(ctxOperation, bson.M{"_id": id}).Decode(&result)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Errorf("No documents found with the given id: %v", err)
			return domains.Event{}, err
		}
		log.Errorf("Error finding event; %v", err)
		return domains.Event{}, err
	}

	return result, nil
}

func (er *EventRepository) UpdateById(id int64, state string) error {

	ctxOperation, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := er.conn.Database(env.Database).Collection(EVENTS_COLLECTION)

	newEventState := domains.EventStateHistory{
		State:     state,
		CreatedAt: time.Now(),
	}

	_, err := collection.UpdateByID(ctxOperation, id, bson.M{"$set": newEventState})

	if err != nil {
		log.Errorf("Error updating event: %s with err %v", id, err)
		return err
	}

	return nil
}
