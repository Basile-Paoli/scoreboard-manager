package events

type Event string

const (
	UpdateScoreboard Event = "update-scoreboard"
	Error            Event = "error"
)

var Events = []struct {
	Value  Event
	TSName string
}{
	{Value: UpdateScoreboard, TSName: "updateScoreboard"},
	{Value: Error, TSName: "error"},
}
