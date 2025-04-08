package support

import "github.com/teamcubation/teamcandidates/projects/teamcandidates-api/internal/event/handler/dto"

type ListEventsResponse struct {
	List dto.EventList `json:"events_list"`
}
