package monitor

import (
	"reflect"
	"testing"

	mqtt "github.com/evanofslack/bambulab-client/mqtt"
	"github.com/stretchr/testify/assert"
)

func TestMergeStandalone(t *testing.T) {
    var og, in *mqtt.Ipcam
    in = &mqtt.Ipcam{IpcamDev: strPtr("dev2"), Resolution: strPtr("1080p")}
	og, _ = mergeIpcam(og, in)
    if og == nil || in == nil {
        t.Fatal("nil")
    }
    if *og.IpcamDev != *in.IpcamDev {
        t.Fatal("unequal")
    }
}

func TestMergeIpcam(t *testing.T) {
	tests := []struct {
		name     string
		og       *mqtt.Ipcam
		in       *mqtt.Ipcam
		expected *mqtt.Ipcam
		changed  bool
	}{
		{
			name:     "both nil",
			og:       nil,
			in:       nil,
			expected: nil,
			changed:  false,
		},
		{
			name: "input nil",
			og: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev1"),
				IpcamRecord: strPtr("enable"),
				Resolution:  strPtr("720p"),
				Timelapse:   strPtr("enable"),
			},
			in:       nil,
			expected: &mqtt.Ipcam{IpcamDev: strPtr("dev1"), IpcamRecord: strPtr("enable"), Resolution: strPtr("720p"), Timelapse: strPtr("enable")},
			changed:  false,
		},
		{
			name:     "original nil, input with fields",
			og:       nil,
			in:       &mqtt.Ipcam{IpcamDev: strPtr("dev2"), Resolution: strPtr("1080p")},
			expected: &mqtt.Ipcam{IpcamDev: strPtr("dev2"), Resolution: strPtr("1080p")},
			changed:  true,
		},
		{
			name: "input with updated fields",
			og: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev1"),
				IpcamRecord: strPtr("disable"),
				Resolution:  strPtr("720p"),
				Timelapse:   strPtr("disable"),
			},
			in: &mqtt.Ipcam{
				IpcamDev:   strPtr("dev2"),
				Resolution: strPtr("1080p"),
				Timelapse:  strPtr("enable"),
			},
			expected: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev2"),
				IpcamRecord: strPtr("disable"),
				Resolution:  strPtr("1080p"),
				Timelapse:   strPtr("enable"),
			},
			changed: true,
		},
		{
			name: "input with same fields (no change)",
			og: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev1"),
				IpcamRecord: strPtr("enable"),
				Resolution:  strPtr("720p"),
				Timelapse:   strPtr("enable"),
			},
			in: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev1"),
				IpcamRecord: strPtr("enable"),
				Resolution:  strPtr("720p"),
				Timelapse:   strPtr("enable"),
			},
			expected: &mqtt.Ipcam{
				IpcamDev:    strPtr("dev1"),
				IpcamRecord: strPtr("enable"),
				Resolution:  strPtr("720p"),
				Timelapse:   strPtr("enable"),
			},
			changed: false,
		},
	}

	for i, tt := range tests {
	    if i != 2 {
	        continue
	    }
		t.Run(tt.name, func(t *testing.T) {
			var changed bool
			tt.og, changed = mergeIpcam(tt.og, tt.in)
			if changed != tt.changed {
				t.Errorf("mergeIpcam() changed = %v, expected %v", changed, tt.changed)
			}
			if !reflect.DeepEqual(tt.og, tt.expected) {
				t.Errorf("mergeIpcam() = %v, expected %v", tt.og, tt.expected)
			}
		})
	}
}

func TestMergeUpload(t *testing.T) {
	// Helper function to create a pointer to a string
	strPtr := func(s string) *string {
		return &s
	}

	// Helper function to create a pointer to an int
	intPtr := func(i int) *int {
		return &i
	}

	tests := []struct {
		name         string
		og           *mqtt.Upload
		in           *mqtt.Upload
		expectedOg   *mqtt.Upload
		expectChange bool
	}{
		{
			name: "nil input does not change",
			og:   &mqtt.Upload{Status: strPtr("idle")},
			in:   nil,
			expectedOg: &mqtt.Upload{
				Status: strPtr("idle"),
			},
			expectChange: false,
		},
		{
			name: "status field updated",
			og:   &mqtt.Upload{Status: strPtr("idle")},
			in:   &mqtt.Upload{Status: strPtr("running")},
			expectedOg: &mqtt.Upload{
				Status: strPtr("running"),
			},
			expectChange: true,
		},
		{
			name: "progress field updated",
			og:   &mqtt.Upload{Progress: intPtr(50)},
			in:   &mqtt.Upload{Progress: intPtr(75)},
			expectedOg: &mqtt.Upload{
				Progress: intPtr(75),
			},
			expectChange: true,
		},
		{
			name: "message field updated",
			og:   &mqtt.Upload{Message: strPtr("uploading")},
			in:   &mqtt.Upload{Message: strPtr("complete")},
			expectedOg: &mqtt.Upload{
				Message: strPtr("complete"),
			},
			expectChange: true,
		},
		{
			name: "no fields updated (identical values)",
			og:   &mqtt.Upload{Status: strPtr("idle"), Progress: intPtr(50), Message: strPtr("uploading")},
			in:   &mqtt.Upload{Status: strPtr("idle"), Progress: intPtr(50), Message: strPtr("uploading")},
			expectedOg: &mqtt.Upload{
				Status:   strPtr("idle"),
				Progress: intPtr(50),
				Message:  strPtr("uploading"),
			},
			expectChange: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var changed bool
			tt.og, changed = mergeUpload(tt.og, tt.in)
			assert.Equal(t, tt.expectChange, changed)
			assert.Equal(t, tt.expectedOg, tt.og)
		})
	}
}

func TestMergeLightsReport(t *testing.T) {
	tests := []struct {
		name         string
		og           *[]mqtt.LightsReport
		in           *[]mqtt.LightsReport
		expectedOg   *[]mqtt.LightsReport
		expectChange bool
	}{
		{
			name:         "nil input does not change",
			og:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			in:           nil,
			expectedOg:   &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectChange: false,
		},
		{
			name:         "new input updates lights report",
			og:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			in:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("off")}, {Node: strPtr("light2"), Mode: strPtr("on")}},
			expectedOg:   &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("off")}, {Node: strPtr("light2"), Mode: strPtr("on")}},
			expectChange: true,
		},
		{
			name:         "identical lights report does not change",
			og:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			in:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectedOg:   &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectChange: false,
		},
		{
			name:         "empty input updates empty og",
			og:           &[]mqtt.LightsReport{},
			in:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectedOg:   &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectChange: true,
		},
		{
			name:         "og is nil and in is not nil",
			og:           nil,
			in:           &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectedOg:   &[]mqtt.LightsReport{{Node: strPtr("light1"), Mode: strPtr("on")}},
			expectChange: true,
		},
		{
			name:         "both og and in are nil",
			og:           nil,
			in:           nil,
			expectedOg:   nil,
			expectChange: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var changed bool
			tt.og, changed = mergeLightsReport(tt.og, tt.in)

			// Verify if the change status matches the expected result
			assert.Equal(t, tt.expectChange, changed)

			// Verify if the contents of og are now equal to the expectedOg
			if tt.expectedOg == nil {
				assert.Nil(t, tt.og)
			} else {
				assert.True(t, reflect.DeepEqual(tt.og, tt.expectedOg))
			}
		})
	}
}

func TestMergePrimatives(t *testing.T) {
	tests := []struct {
		name          string
		og            *mqtt.Print
		in            *mqtt.Print
		expectedMerge bool
		expectedPrint *mqtt.Print
	}{
		{
			name: "No changes",
			og: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
			in: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
			expectedMerge: false,
			expectedPrint: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
		},
		{
			name: "Changes in AmsRfidStatus and BedTemper",
			og: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
			in: &mqtt.Print{
				AmsRfidStatus: intPtr(2),
				BedTemper:     floatPtr(65.0),
			},
			expectedMerge: true,
			expectedPrint: &mqtt.Print{
				AmsRfidStatus: intPtr(2),
				BedTemper:     floatPtr(65.0),
			},
		},
		{
			name: "Nil to value changes",
			og: &mqtt.Print{
				AmsRfidStatus: nil,
				BedTemper:     nil,
			},
			in: &mqtt.Print{
				AmsRfidStatus: intPtr(2),
				BedTemper:     floatPtr(65.0),
			},
			expectedMerge: true,
			expectedPrint: &mqtt.Print{
				AmsRfidStatus: intPtr(2),
				BedTemper:     floatPtr(65.0),
			},
		},
		{
			name: "No changes when in is nil",
			og: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
			in:            nil, // `in` being nil should not change anything
			expectedMerge: false,
			expectedPrint: &mqtt.Print{
				AmsRfidStatus: intPtr(1),
				BedTemper:     floatPtr(60.5),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			changed := mergePrimatives(tt.og, tt.in)
			assert.Equal(t, tt.expectedMerge, changed)
			assert.Equal(t, tt.expectedPrint, tt.og)
		})
	}
}

func strPtr(s string) *string     { return &s }
func intPtr(i int) *int           { return &i }
func floatPtr(f float64) *float64 { return &f }
