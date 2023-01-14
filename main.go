package main

import (
	"encoding/json"
	"log"

	"github.com/fiatjaf/relayer"
	"github.com/nbd-wtf/go-nostr"
)

type DoNothingStore struct{}

// Implement relayer's Storage interface

func (d *DoNothingStore) Init() error                      { return nil }
func (d *DoNothingStore) DeleteEvent(string, string) error { return nil }
func (d *DoNothingStore) SaveEvent(*nostr.Event) error     { return nil }
func (d *DoNothingStore) QueryEvents(*nostr.Filter) ([]nostr.Event, error) {
	return []nostr.Event{}, nil
}

type Relay struct{}

// Implement relay's Relay interface

func (r *Relay) Name() string             { return "foo" }
func (r *Relay) Storage() relayer.Storage { return &DoNothingStore{} }
func (r *Relay) OnInitialized(s *relayer.Server) {
	s.Router().Path("/").HandlerFunc(handleMainPage)
}
func (r *Relay) Init() error             { return nil }
func (r *Relay) BeforeSave(*nostr.Event) {}
func (r *Relay) AfterSave(*nostr.Event)  {}
func (r *Relay) AcceptEvent(evt *nostr.Event) bool {
	// block events that are too large
	jsonb, _ := json.Marshal(evt)
	return len(jsonb) <= 10000
}

func main() {
	if err := relayer.Start(&Relay{}); err != nil {
		log.Fatalf("server terminated: %v", err)
	}
}
