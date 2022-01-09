package framework

type StationManager struct {
	IsPlatformFree bool
	TrainQueue     []Train
}

func NewStationManger() *StationManager {
	return &StationManager{
		IsPlatformFree: true,
	}
}

func (s *StationManager) CanArrive(t Train) bool {
	if s.IsPlatformFree {
		s.IsPlatformFree = false
		return true
	}
	s.TrainQueue = append(s.TrainQueue, t)
	return false
}

func (s *StationManager) NotifyAboutDeparture() {
	if !s.IsPlatformFree {
		s.IsPlatformFree = true
	}
	if len(s.TrainQueue) > 0 {
		firstTrainInQueue := s.TrainQueue[0]
		s.TrainQueue = s.TrainQueue[1:]
		firstTrainInQueue.PermitArrival()
	}
}