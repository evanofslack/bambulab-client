package monitor

import (
	"strconv"
	"strings"

	"github.com/evanofslack/bambulab-client/mqtt"
	opt "github.com/moznion/go-optional"
)

// State is an interpreted view of state of printer.
type State struct {
	Ams          Ams
	Bed          Bed
	Camera       Camera
	Chamber      Chamber
	CurrentPrint CurrentPrint
	Fans         Fans
	Gcode        Gcode
	Lights       Lights
	Nozzle       Nozzle
	Speed        Speed
	UpgradeState UpgradeState
	Upload       Upload
	ProfileID    opt.Option[string]
	ProjectID    opt.Option[string]
	SDCard       opt.Option[bool]
	Wifi         opt.Option[float64]
}

// Ams is AMS metadata
type Ams struct {
	Enabled    opt.Option[bool]
	Inserting  opt.Option[bool]
	Powered    opt.Option[bool]
	RfidStatus opt.Option[int]
	Version    opt.Option[int]
	Units      []AmsUnit
}

// AmsUnit is an AMS unit
type AmsUnit struct {
	Humidity    opt.Option[float64]
	ID          opt.Option[string]
	Temperature opt.Option[float64]
	Trays       []AmsTray
}

// AmsTray is one slot in an AMS
type AmsTray struct {
	Brand     opt.Option[string]
	Color     opt.Option[string]
	ID        opt.Option[string]
	K         opt.Option[float64]
	Name      opt.Option[string]
	Remaining opt.Option[int]
	TempMax   opt.Option[float64]
	TempMin   opt.Option[float64]
	Weight    opt.Option[float64]
}

type Bed struct {
	Temperature       opt.Option[float64]
	TemperatureTarget opt.Option[int]
}

type Camera struct {
	Recording  opt.Option[bool]
	Resolution opt.Option[string]
	Timelapse  opt.Option[bool]
}

type Chamber struct {
	Temperature opt.Option[int]
}

type CurrentPrint struct {
	LayerNumber       opt.Option[int]
	LayerNumberTarget opt.Option[int]
	Percent           opt.Option[int]
	Subtask           opt.Option[string]
	PrintError        opt.Option[int]
}

type Fans struct {
	Auxilliary opt.Option[float64]
	Chamber    opt.Option[float64]
	Part       opt.Option[float64]
	Hotend     opt.Option[float64]
}

type Gcode struct {
	File  opt.Option[string]
	State opt.Option[string]
}

type Lights struct {
	Chamber opt.Option[bool]
}

type Nozzle struct {
	Diameter          opt.Option[float64]
	TemperatureTarget opt.Option[int]
	Temperature       opt.Option[float64]
	Type              opt.Option[string]
}

type Speed struct {
	Level     opt.Option[int]
	LevelName opt.Option[string]
	Magnitude opt.Option[int]
}

type UpgradeState struct {
	ErrorCode    opt.Option[int]
	ForceUpgrade opt.Option[bool]
	Message      opt.Option[string]
	Module       opt.Option[string]
	Progress     opt.Option[string]
	SequenceID   opt.Option[int]
	Status       opt.Option[string]
}

type Upload struct {
	Message  opt.Option[string]
	Progress opt.Option[int]
	Status   opt.Option[string]
}

// fromMessage creates state from mqtt message
func stateFromMessage(m *mqtt.Message) State {
	s := State{}
	p := m.Print
	if p == nil {
		return s
	}
	s.Ams = interpretAms(p.Ams)
	s.Bed = interpretBed(p)
	s.Camera = interpretCamera(p.Ipcam)
	s.Chamber = interpretChamber(p)
	s.CurrentPrint = interpretCurrentPrint(p)
	s.Gcode = interpretGcode(p)
	s.Fans = interpretFans(p)
	s.Lights = interpretLights(p)
	s.Nozzle = interpretNozzle(p)
	s.Speed = interpretSpeed(p)
	s.UpgradeState = interpretUpgradeState(p.UpgradeState)
	s.Upload = interpretUpload(p.Upload)

	s.ProfileID = opt.FromNillable(p.ProfileID)
	s.ProjectID = opt.FromNillable(p.ProjectID)
	s.SDCard = opt.FromNillable(p.Sdcard)
	s.SDCard = opt.FromNillable(p.Sdcard)
	s.Wifi = parseWifi(p.WifiSignal)
	return s
}

func interpretAms(m *mqtt.Ams) Ams {
	ams := Ams{}
	if m == nil {
		return ams
	}
	ams.Powered = opt.FromNillable(m.PowerOnFlag)
	ams.Inserting = opt.FromNillable(m.InsertFlag)
	ams.Version = opt.FromNillable(m.Version)
	inner := m.Ams
	if inner == nil {
		return ams
	}
	ams.Units = interpretAmsInner(m.Ams)

	return ams
}

func interpretAmsInner(in *[]mqtt.AmsInner) []AmsUnit {
	units := []AmsUnit{}
	if in == nil {
		return units
	}
	for _, inner := range *in {
		unit := AmsUnit{}
		unit.Humidity = strToFloat(inner.Humidity)
		unit.Temperature = strToFloat(inner.Temp)
		unit.ID = opt.FromNillable(inner.ID)
		unit.Trays = interpretAmsTray(inner.Tray)
		units = append(units, unit)
	}
	return units
}

func interpretAmsTray(ins *[]mqtt.Tray) []AmsTray {
	trays := []AmsTray{}
	if ins == nil {
		return trays
	}
	for _, in := range *ins {
		tray := AmsTray{}
		tray.Brand = opt.FromNillable(in.TraySubBrands)
		tray.Color = opt.FromNillable(in.TrayColor)
		tray.ID = opt.FromNillable(in.ID)
		tray.K = opt.FromNillable(in.K)
		tray.Name = opt.FromNillable(in.TrayIDName)
		tray.Remaining = opt.FromNillable(in.Remain)
		tray.TempMax = strToFloat(in.NozzleTempMax)
		tray.TempMin = strToFloat(in.NozzleTempMin)
		tray.Weight = strToFloat(in.TrayWeight)
		trays = append(trays, tray)
	}
	return trays
}

func interpretBed(p *mqtt.Print) Bed {
	bed := Bed{}
	if p == nil {
		return bed
	}
	bed.Temperature = opt.FromNillable(p.BedTemper)
	bed.TemperatureTarget = opt.FromNillable(p.BedTargetTemper)
	return bed
}

func interpretCamera(c *mqtt.Ipcam) Camera {
	camera := Camera{}
	if c == nil {
		return camera
	}
	camera.Recording = enabledToBool(c.IpcamRecord)
	camera.Resolution = opt.FromNillable(c.Resolution)
	camera.Timelapse = enabledToBool(c.Timelapse)
	return camera
}

func interpretChamber(p *mqtt.Print) Chamber {
	chamber := Chamber{}
	if p == nil {
		return chamber
	}
	chamber.Temperature = opt.FromNillable(p.ChamberTemper)
	return chamber
}

func interpretCurrentPrint(p *mqtt.Print) CurrentPrint {
	c := CurrentPrint{}
	if p == nil {
		return c
	}
	c.LayerNumber = opt.FromNillable(p.LayerNum)
	c.LayerNumberTarget = opt.FromNillable(p.TotalLayerNum)
	c.Percent = opt.FromNillable(p.McPercent)
	c.Subtask = opt.FromNillable(p.SubtaskName)
	c.PrintError = opt.FromNillable(p.PrintError)
	return c
}

func interpretFans(p *mqtt.Print) Fans {
	f := Fans{}
	if p == nil {
		return f
	}
	f.Auxilliary = fanspeed(p.BigFan1Speed)
	f.Chamber = fanspeed(p.BigFan2Speed)
	f.Part = fanspeed(p.CoolingFanSpeed)
	f.Hotend = fanspeed(p.HeatbreakFanSpeed)
	return f
}

func interpretGcode(p *mqtt.Print) Gcode {
	g := Gcode{}
	if p == nil {
		return g
	}
	g.File = opt.FromNillable(p.GcodeFile)
	g.State = opt.FromNillable(p.GcodeState)
	return g
}

func interpretLights(p *mqtt.Print) Lights {
	l := Lights{}
	if p == nil {
		return l
	}
	lr := p.LightsReport
	if lr == nil {
		return l
	}
	if len(*lr) == 0 {
		return l
	}

	for _, light := range *lr {
		if mode, node := light.Mode, light.Node; mode != nil && node != nil {
			if strings.ToLower(*node) == "chamber_light" {
				switch strings.ToLower(*mode) {
				case "on":
					l.Chamber = opt.Some(true)
				case "off":
					l.Chamber = opt.Some(false)
				default:
					l.Chamber = opt.None[bool]()
				}
			}
		}
	}
	return l
}

func interpretNozzle(p *mqtt.Print) Nozzle {
	n := Nozzle{}
	if p == nil {
		return n
	}
	n.Diameter = strToFloat(p.NozzleDiameter)
	n.Temperature = opt.FromNillable(p.NozzleTemper)
	n.TemperatureTarget = opt.FromNillable(p.NozzleTargetTemper)
	n.Type = opt.FromNillable(p.NozzleType)
	return n
}

func interpretSpeed(p *mqtt.Print) Speed {
	s := Speed{}
	if p == nil {
		return s
	}
	s.Level = opt.FromNillable(p.SpdLvl)
	s.Magnitude = opt.FromNillable(p.SpdMag)
	if s.Level.IsNone() {
		s.LevelName = opt.None[string]()
	}
	if s.Level.IsNone() {
		s.LevelName = opt.None[string]()
	} else {
		switch s.Level.Unwrap() {
		case 1:
			s.LevelName = opt.Some("silent")
		case 2:
			s.LevelName = opt.Some("standard")
		case 3:
			s.LevelName = opt.Some("sport")
		case 4:
			s.LevelName = opt.Some("ludicrous")
		default:
			s.LevelName = opt.None[string]()
		}
	}
	return s
}

func interpretUpgradeState(us *mqtt.UpgradeState) UpgradeState {
	u := UpgradeState{}
	if us == nil {
		return u
	}
	u.ErrorCode = opt.FromNillable(us.ErrCode)
	u.ForceUpgrade = opt.FromNillable(us.ForceUpgrade)
	u.Message = opt.FromNillable(us.Message)
	u.Module = opt.FromNillable(us.Module)
	u.Progress = opt.FromNillable(us.Progress)
	u.SequenceID = opt.FromNillable(us.SequenceID)
	u.Status = opt.FromNillable(us.Status)
	return u
}

func interpretUpload(up *mqtt.Upload) Upload {
	u := Upload{}
	if up == nil {
		return u
	}
	u.Message = opt.FromNillable(up.Message)
	u.Progress = opt.FromNillable(up.Progress)
	u.Status = opt.FromNillable(up.Status)
	return u
}

func strToFloat(in *string) opt.Option[float64] {
	if in == nil {
		return opt.None[float64]()
	}
	if parsed, err := strconv.ParseFloat(*in, 64); err != nil {
		return opt.None[float64]()
	} else {
		return opt.Some(parsed)
	}
}

func fanspeed(in *string) opt.Option[float64] {
	if in == nil {
		return opt.None[float64]()
	}
	raw, err := strconv.ParseFloat(*in, 64)
	if err != nil {
		return opt.None[float64]()
	}
	percent := (raw / 15) * 100
	return opt.Some(percent)
}

func enabledToBool(in *string) opt.Option[bool] {
	if in == nil {
		return opt.None[bool]()
	}
	switch strings.ToLower(*in) {
	case "enabled":
		return opt.Some(true)
	case "disabled":
		return opt.Some(false)
	default:
		return opt.None[bool]()
	}
}

func parseWifi(in *string) opt.Option[float64] {
	if in == nil {
		return opt.None[float64]()
	}
	w := strings.TrimSuffix(*in, "dBm")
	return strToFloat(&w)
}
