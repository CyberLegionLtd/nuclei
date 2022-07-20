package monitor

import (
	"os"
	"strings"
	"sync"
	"sync/atomic"

	jsoniter "github.com/json-iterator/go"
)

// Scan is the global scan monitor instance
var Scan *ScanMonitor

func init() {
	Scan = NewScanMonitor()
}

// ScanMonitor is a monitor structure for scan monitoring
type ScanMonitor struct {
	mu *sync.Mutex

	// counters for processing
	targets   int64
	templates int64
	// inflight target-template pairs
	inFlightTargetTemplates map[string]struct{}
}

// NewScanMonitor returns a scan monitor instance
func NewScanMonitor() *ScanMonitor {
	return &ScanMonitor{inFlightTargetTemplates: make(map[string]struct{}), mu: &sync.Mutex{}}
}

// IncrementTargets increments targets counter by 1
func (s *ScanMonitor) IncrementTargets() {
	atomic.AddInt64(&s.targets, 1)
}

// DecrementTargets decrements targets counter by 1
func (s *ScanMonitor) DecrementTargets() {
	atomic.AddInt64(&s.targets, -1)
}

// IncrementTemplates increments templates counter by 1
func (s *ScanMonitor) IncrementTemplates() {
	atomic.AddInt64(&s.templates, 1)
}

// DecrementTemplates decrements templates counter by 1
func (s *ScanMonitor) DecrementTemplates() {
	atomic.AddInt64(&s.templates, -1)
}

// InsertTargetTemplate inserts a target template to set
func (s *ScanMonitor) InsertTargetTemplate(target, template string) {
	s.mu.Lock()
	input := strings.Join([]string{template, target}, ":")
	if _, ok := s.inFlightTargetTemplates[input]; !ok {
		s.inFlightTargetTemplates[input] = struct{}{}
	}
	s.mu.Unlock()

	s.syncToMonitorFile() // sync with each insert operation
}

// DeleteTargetTemplate deletes a target template from set
func (s *ScanMonitor) DeleteTargetTemplate(target, template string) {
	s.mu.Lock()
	input := strings.Join([]string{template, target}, ":")
	delete(s.inFlightTargetTemplates, input)
	s.mu.Unlock()
}

type scanMonitorFileData struct {
	Targets   int64    `json:"targets"`
	Templates int64    `json:"templates"`
	Inflight  []string `json:"inflight"`
}

func (s *ScanMonitor) syncToMonitorFile() {
	file, err := os.Create("scan.monitor")
	if err != nil {
		return
	}
	defer file.Close()

	data := &scanMonitorFileData{
		Targets:   atomic.LoadInt64(&s.targets),
		Templates: atomic.LoadInt64(&s.templates),
	}
	s.mu.Lock()
	data.Inflight = make([]string, 0, len(s.inFlightTargetTemplates))
	for k := range s.inFlightTargetTemplates {
		data.Inflight = append(data.Inflight, k)
	}
	s.mu.Unlock()

	encoder := jsoniter.NewEncoder(file)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(data)
}
