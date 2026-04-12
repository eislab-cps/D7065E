package store

import (
	"sync"
	"time"

	"github.com/eislab-cps/buildingsim/pkg/model"
)

type MemoryStore struct {
	mu sync.RWMutex

	// Building data (read-only, loaded at startup)
	building        model.Building
	floors          map[string]*model.FloorData // level id -> floor data
	crossFloorEdges []model.CrossFloorEdge
	multiFloorGraph *model.NavGraph

	// Equipment (global, in-memory)
	equipment        map[string]*model.Equipment
	sensors          map[string]*model.Sensor    // sensor id -> sensor
	sensorEquipment  map[string]string           // sensor id -> equipment id
	actuators        map[string]*model.Actuator  // actuator id -> actuator
	actuatorEquipment map[string]string          // actuator id -> equipment id
	equipmentVersion int64

	// Sessions
	sessions map[string]*model.Session
}

func NewMemoryStore() *MemoryStore {
	return &MemoryStore{
		floors:            make(map[string]*model.FloorData),
		equipment:         make(map[string]*model.Equipment),
		sensors:           make(map[string]*model.Sensor),
		sensorEquipment:   make(map[string]string),
		actuators:         make(map[string]*model.Actuator),
		actuatorEquipment: make(map[string]string),
		sessions:          make(map[string]*model.Session),
	}
}

// === Building ===

func (s *MemoryStore) SetBuilding(b model.Building) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.building = b
}

func (s *MemoryStore) GetBuilding() model.Building {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.building
}

func (s *MemoryStore) SetFloorData(levelID string, data *model.FloorData) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.floors[levelID] = data
}

func (s *MemoryStore) GetFloorData(levelID string) (*model.FloorData, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	d, ok := s.floors[levelID]
	return d, ok
}

func (s *MemoryStore) SetCrossFloorEdges(edges []model.CrossFloorEdge) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.crossFloorEdges = edges
}

func (s *MemoryStore) GetCrossFloorEdges() []model.CrossFloorEdge {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.crossFloorEdges
}

func (s *MemoryStore) SetMultiFloorGraph(g *model.NavGraph) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.multiFloorGraph = g
}

func (s *MemoryStore) GetMultiFloorGraph() *model.NavGraph {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.multiFloorGraph
}

func (s *MemoryStore) GetFloors() map[string]*model.FloorData {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.floors
}

// === Equipment ===

func (s *MemoryStore) CreateEquipment(e *model.Equipment) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e.Version = s.equipmentVersion
	if e.Sensors == nil {
		e.Sensors = []model.Sensor{}
	}
	if e.Actuators == nil {
		e.Actuators = []model.Actuator{}
	}
	s.equipment[e.ID] = e
}

func (s *MemoryStore) GetEquipment(id string) (*model.Equipment, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.equipment[id]
	return e, ok
}

func (s *MemoryStore) ListEquipment(level, roomFilter, typeFilter, category string) []*model.Equipment {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []*model.Equipment
	for _, e := range s.equipment {
		if level != "" && e.Level != level {
			continue
		}
		if roomFilter != "" && e.Room != roomFilter {
			continue
		}
		if typeFilter != "" && e.Type != typeFilter {
			continue
		}
		if category != "" && e.Category != category {
			continue
		}
		result = append(result, e)
	}
	return result
}

func (s *MemoryStore) UpdateEquipment(e *model.Equipment) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, ok := s.equipment[e.ID]; !ok {
		return false
	}
	s.equipment[e.ID] = e
	return true
}

func (s *MemoryStore) DeleteEquipment(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.equipment[id]
	if !ok {
		return false
	}
	// Clean up sensor/actuator mappings
	for _, sen := range e.Sensors {
		delete(s.sensors, sen.ID)
		delete(s.sensorEquipment, sen.ID)
	}
	for _, act := range e.Actuators {
		delete(s.actuators, act.ID)
		delete(s.actuatorEquipment, act.ID)
	}
	delete(s.equipment, id)
	return true
}

func (s *MemoryStore) BumpEquipmentVersion() int64 {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.equipmentVersion++
	return s.equipmentVersion
}

func (s *MemoryStore) GetEquipmentVersion() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.equipmentVersion
}

// === Sensors ===

func (s *MemoryStore) AddSensor(equipmentID string, sen *model.Sensor) (bool, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.equipment[equipmentID]
	if !ok {
		return false, "equipment not found"
	}
	if _, exists := s.sensors[sen.ID]; exists {
		return false, "sensor already exists"
	}
	sen.Timestamp = time.Now()
	e.Sensors = append(e.Sensors, *sen)
	s.sensors[sen.ID] = &e.Sensors[len(e.Sensors)-1]
	s.sensorEquipment[sen.ID] = equipmentID
	return true, ""
}

func (s *MemoryStore) GetSensorsForEquipment(equipmentID string) ([]model.Sensor, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.equipment[equipmentID]
	if !ok {
		return nil, false
	}
	return e.Sensors, true
}

func (s *MemoryStore) DeleteSensor(sensorID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	eqID, ok := s.sensorEquipment[sensorID]
	if !ok {
		return false
	}
	e := s.equipment[eqID]
	for i, sen := range e.Sensors {
		if sen.ID == sensorID {
			e.Sensors = append(e.Sensors[:i], e.Sensors[i+1:]...)
			break
		}
	}
	delete(s.sensors, sensorID)
	delete(s.sensorEquipment, sensorID)
	return true
}

func (s *MemoryStore) SetSensorValue(sensorID string, val model.SensorValue) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	eqID, ok := s.sensorEquipment[sensorID]
	if !ok {
		return false
	}
	e := s.equipment[eqID]
	for i := range e.Sensors {
		if e.Sensors[i].ID == sensorID {
			e.Sensors[i].DataType = val.DataType
			if val.DataType == "binary" && val.BinaryValue != nil {
				e.Sensors[i].BinaryValue = *val.BinaryValue
			} else {
				e.Sensors[i].Value = val.Value
			}
			e.Sensors[i].Timestamp = time.Now()
			return true
		}
	}
	return false
}

// === Actuators ===

func (s *MemoryStore) AddActuator(equipmentID string, act *model.Actuator) (bool, string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	e, ok := s.equipment[equipmentID]
	if !ok {
		return false, "equipment not found"
	}
	if _, exists := s.actuators[act.ID]; exists {
		return false, "actuator already exists"
	}
	act.Timestamp = time.Now()
	e.Actuators = append(e.Actuators, *act)
	s.actuators[act.ID] = &e.Actuators[len(e.Actuators)-1]
	s.actuatorEquipment[act.ID] = equipmentID
	return true, ""
}

func (s *MemoryStore) GetActuatorsForEquipment(equipmentID string) ([]model.Actuator, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	e, ok := s.equipment[equipmentID]
	if !ok {
		return nil, false
	}
	return e.Actuators, true
}

func (s *MemoryStore) DeleteActuator(actuatorID string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	eqID, ok := s.actuatorEquipment[actuatorID]
	if !ok {
		return false
	}
	e := s.equipment[eqID]
	for i, act := range e.Actuators {
		if act.ID == actuatorID {
			e.Actuators = append(e.Actuators[:i], e.Actuators[i+1:]...)
			break
		}
	}
	delete(s.actuators, actuatorID)
	delete(s.actuatorEquipment, actuatorID)
	return true
}

func (s *MemoryStore) SetActuatorState(actuatorID string, state model.ActuatorState) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	eqID, ok := s.actuatorEquipment[actuatorID]
	if !ok {
		return false
	}
	e := s.equipment[eqID]
	for i := range e.Actuators {
		if e.Actuators[i].ID == actuatorID {
			e.Actuators[i].State = state.State
			e.Actuators[i].Timestamp = time.Now()
			return true
		}
	}
	return false
}

// === Sessions ===

func (s *MemoryStore) CreateSession(id string) *model.Session {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess := &model.Session{
		ID:         id,
		Viewport:   model.Viewport{Mode: "3d", Floor: "level0", Zoom: 1.0},
		Highlights: []model.RoomHighlight{},
		Occupancy:  make(map[int]model.RoomOccupancy),
		Coverage:   []model.CoverageZone{},
		LastWSActive: time.Now(),
	}
	s.sessions[id] = sess
	return sess
}

func (s *MemoryStore) ListSessions() []*model.Session {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var result []*model.Session
	for _, sess := range s.sessions {
		result = append(result, sess)
	}
	return result
}

func (s *MemoryStore) GetSession(id string) (*model.Session, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	sess, ok := s.sessions[id]
	return sess, ok
}

func (s *MemoryStore) DeleteSession(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.sessions[id]
	if ok {
		delete(s.sessions, id)
	}
	return ok
}

func (s *MemoryStore) UpdateSessionViewport(id string, vp model.Viewport) (int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[id]
	if !ok {
		return 0, false
	}
	sess.Viewport = vp
	sess.Version++
	return sess.Version, true
}

func (s *MemoryStore) UpdateSessionHighlights(id string, highlights []model.RoomHighlight) (int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[id]
	if !ok {
		return 0, false
	}
	sess.Highlights = highlights
	sess.Version++
	return sess.Version, true
}

func (s *MemoryStore) UpdateSessionOccupancy(id string, occupancy map[int]model.RoomOccupancy) (int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[id]
	if !ok {
		return 0, false
	}
	sess.Occupancy = occupancy
	sess.Version++
	return sess.Version, true
}

func (s *MemoryStore) UpdateSessionCoverage(id string, coverage []model.CoverageZone) (int64, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[id]
	if !ok {
		return 0, false
	}
	sess.Coverage = coverage
	sess.Version++
	return sess.Version, true
}

func (s *MemoryStore) UpdateSessionRoute(id string, route *model.RouteResult) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	sess, ok := s.sessions[id]
	if !ok {
		return false
	}
	sess.Route = route
	sess.Version++
	return true
}

func (s *MemoryStore) TouchSession(id string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if sess, ok := s.sessions[id]; ok {
		sess.LastWSActive = time.Now()
	}
}

func (s *MemoryStore) PurgeInactiveSessions(maxAge time.Duration) []string {
	s.mu.Lock()
	defer s.mu.Unlock()
	var purged []string
	now := time.Now()
	for id, sess := range s.sessions {
		if now.Sub(sess.LastWSActive) > maxAge {
			delete(s.sessions, id)
			purged = append(purged, id)
		}
	}
	return purged
}
