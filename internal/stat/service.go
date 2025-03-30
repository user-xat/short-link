package stat

import (
	"log"

	"github.com/user-xat/short-link/internal/models"
	"github.com/user-xat/short-link/pkg/di"
	"github.com/user-xat/short-link/pkg/event"
)

type StatServiceDeps struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

type StatService struct {
	EventBus       *event.EventBus
	StatRepository di.IStatRepository
}

func NewStatService(deps StatServiceDeps) *StatService {
	return &StatService{
		EventBus:       deps.EventBus,
		StatRepository: deps.StatRepository,
	}
}

func (s *StatService) AddClick(msg *event.Event) {
	switch msg.Type {
	case event.EventLinkVisited:
		link, ok := msg.Data.(*models.Link)
		if !ok {
			log.Fatalln("Bad EventLinkVisited Data: ", msg.Data)
		}
		s.StatRepository.AddClick(link.ID)
	}
}
