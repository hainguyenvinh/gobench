package worker

import (
	"testing"

	gometrics "github.com/rcrowley/go-metrics"
	"github.com/stretchr/testify/assert"
)

func loadValidPlugin(w *Worker) error {
	so := "./script/valid.so"
	return w.Load(so)
}

type nilLog struct{}

func (l *nilLog) Counter(id, title string, time, c int64) error {
	return nil
}

func (l *nilLog) Histogram(id, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (l *nilLog) Gauge(id, title string, time int64, g int64) error {
	return nil
}

func newNilLog() metricLogger {
	return &nilLog{}
}

// func TestNodeWithConfig(t *testing.T) {
// }

func TestNew(t *testing.T) {
	n1, err := NewWorker(newNilLog())
	assert.Nil(t, err)
	n2, err := NewWorker(newNilLog())
	assert.Nil(t, err)

	assert.Equal(t, n1, n2)

	assert.Equal(t, n1.status, Idle)
	assert.Nil(t, n1.cancel)
	assert.Equal(t, n1.units, make(map[string]unit))

	assert.False(t, n1.Running())
}

func TestLoadPlugin(t *testing.T) {
	n, _ := NewWorker(newNilLog())
	so := "./script/valid.so"
	assert.Nil(t, n.Load(so))
	assert.NotNil(t, n.vus)
	assert.False(t, n.Running())
}

// func TestRunPlugin(t *testing.T) {
// 	n, _ := NewWorker(newNilLog())
// 	assert.Nil(t, loadValidPlugin(n))
// 	assert.Nil(t, n.Run())
// 	assert.False(t, n.Running())
// }
