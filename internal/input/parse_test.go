package input

import (
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	type args struct {
		rawEvent RawEvent
	}
	tests := []struct {
		name string
		args args
		want Event
	}{
		{
			name: "t1", args: args{
			struct {
				EventType EventType
				Payoff    []byte
			}{
				EventType: KeyPress, Payoff: []byte(`
    deviceId: 17 (17)
    detail: 37
    flags:
    root: 4977.00/686.00
    event: 4977.00/686.00
    buttons:
    modifiers: locked 0x10 latched 0 base 0 effective: 0x10
    group: locked 0 latched 0 base 0 effective: 0
    valuators:
    windows: root 0x1a0 event 0x1a0 child 0x8e00002
EVENT type 13 (RawKeyPress)
    deviceId: 3 (17)
    detail: 54
    valuators:`),
			},
		}, want: Event{
			EventType: KeyPress,
			DeviceId:  17,
			Detail:    37,
			Modifiers: 0x10,
		},
		},

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Parse(tt.args.rawEvent); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
