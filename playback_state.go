package avebi

// Video playback state can be [Stopped], [Playing] or [Paused].
type PlaybackState uint8

// Returns a string representation of the playback state
// ("Stopped", "Playing", "Paused", "<invalid>").
func (s PlaybackState) String() string {
	switch s {
	case Stopped:
		return "Stopped"
	case Playing:
		return "Playing"
	case Paused:
		return "Paused"
	default:
		return "<invalid>"
	}
}

const (
	Stopped PlaybackState = iota
	Playing
	Paused
	invalidPlaybackState
)
