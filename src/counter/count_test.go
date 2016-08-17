package counter

import (
    "testing"
    "github.com/stretchr/testify/mock"
    "os"
    "../reader"
)

type MockWriter struct {
    mock.Mock
}

func (m *MockWriter) WriteCount(count map[string]int) {
    m.Called(count)
}

func TestCount(t *testing.T) {

    // setup

    // input
    fp, err := os.Open("../../testdata/Posts_count_test.xml")
    if err != nil {
        t.Error(err)
    }
    defer fp.Close()
    input := reader.NewLineReader(fp)

    // expected
    expected := map[string]int{"tried":100, "control":100, "is":100, "code":200, "the":200, "but":100, "it":100, "in":200, "work":100, "should":100, "my":100, "get":100, "or":100, "then":100, "making":100, "form's":200, "a":500, "build":100, "double":200, "track-bar":100, "for":100, "try":100, "me":100, "setting":100, "error":100, "vb":100, "fine":100, "cannot":100, "doesn't":100, "convert":100, "net":100, "to":400, "i":500, "this":300, "opacity":200, "change":100, "decimal":200, "past":100, "when":200, "worked":100, "use":200, "type":100, "want":100, "implicitly":100, "has":100}

    // mock
    mock := new(MockWriter)
    mock.On("WriteCount", expected)

    // exercise
    Count(input, mock)

    // verify
    mock.AssertExpectations(t)
}

func TestCountCalledTimes(t *testing.T) {

    // setup

    // settings
    channelSize = 100
    logInterval = 100
    countPoolLowerLimit = 30

    // input
    fp, err := os.Open("../../testdata/Posts_count_test.xml")
    if err != nil {
        t.Error(err)
    }
    defer fp.Close()
    input := reader.NewLineReader(fp)

    // expected
    expected := mock.AnythingOfType("map[string]int")

    // mock
    mock := new(MockWriter)
    mock.On("WriteCount", expected).Times(100)

    // exercise
    Count(input, mock)

    // verify
    mock.AssertExpectations(t)
}
