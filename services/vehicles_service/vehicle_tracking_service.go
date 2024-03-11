package vehicles_service

import (
	"encoding/json"
	"log"
	ws "scm-api/ws"
)

type TrackingService struct {
	broker *ws.Broker
}

func NewTrackingService(broker *ws.Broker) *TrackingService {
	return &TrackingService{broker: broker}
}

func (ts *TrackingService) BroadcastLocationUpdate(roomID string, locationData interface{}) {
	log.Printf("Broadcasting location update to room: %s", roomID)
	message, err := json.Marshal(locationData)
	if err != nil {
		log.Printf("Error marshaling location data: %v", err)
		return
	}
	ts.broker.Broadcast <- ws.Message{
		RoomID:  roomID,
		Message: message,
	}
}
