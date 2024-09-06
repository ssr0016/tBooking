package repositories

import (
	"context"
	"tBookingV1/models"
	"time"
)

type EventRepository struct {
	db any
}

func NewEventRepository(db any) models.EventRepository {
	return &EventRepository{
		db: db,
	}
}

// GetMany implements models.EventRepository.
func (r *EventRepository) GetMany(ctx context.Context) ([]*models.Event, error) {
	events := []*models.Event{}

	events = append(events, &models.Event{
		ID:        "1",
		Name:      "My First Event",
		Location:  "Location 1",
		Date:      time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	return events, nil
}

// GetOne implements models.EventRepository.
func (r *EventRepository) GetOne(ctx context.Context, eventId string) (*models.Event, error) {
	return nil, nil
}

// CreateOne implements models.EventRepository.
func (r *EventRepository) CreateOne(ctx context.Context, event models.Event) (*models.Event, error) {
	return nil, nil
}
