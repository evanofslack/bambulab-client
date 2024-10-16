package monitor

import (
	"reflect"

	mqtt "github.com/evanofslack/bambulab-client/mqtt"
)

// mergeState compares original state to an incoming (partial) state update from
// mqtt message. If incoming message has a given field, and the field is different than
// the current state, update the current state.
// Returns true if an update was made (aka the state changed).
func mergeState(og *mqtt.Message, in *mqtt.Message) (*mqtt.Message, bool) {
	changed := false
	// No data, nothing to merge
	if in == nil || in.Print == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.Message{}
	}
	if og.Print == nil {
		og.Print = &mqtt.Print{}
	}
	if in.Print.Ipcam != nil {
		var ipcamChanged bool
		og.Print.Ipcam, ipcamChanged = mergeIpcam(og.Print.Ipcam, in.Print.Ipcam)
		changed = changed || ipcamChanged
	}
	if in.Print.Upload != nil {
	    var uploadChanged bool
		og.Print.Upload, uploadChanged = mergeUpload(og.Print.Upload, in.Print.Upload)
		changed = changed || uploadChanged
	}
	if in.Print.UpgradeState != nil {
	    var upgradeStateChanged bool
		og.Print.UpgradeState, upgradeStateChanged = mergeUpgradeState(og.Print.UpgradeState, in.Print.UpgradeState)
		changed = changed || upgradeStateChanged
	}
	if in.Print.LightsReport != nil {
	    var lightsReportChanged bool
		og.Print.LightsReport, lightsReportChanged = mergeLightsReport(og.Print.LightsReport, in.Print.LightsReport)
		changed = changed || lightsReportChanged
	}
	if in.Print.VtTray != nil {
	    var vtTrayChanged bool
		og.Print.VtTray, vtTrayChanged = mergeVtTray(og.Print.VtTray, in.Print.VtTray)
		changed = changed || vtTrayChanged
	}
	if in.Print.Online != nil {
	    var onlineChanged bool
		og.Print.Online, onlineChanged = mergeOnline(og.Print.Online, in.Print.Online)
		changed = changed || onlineChanged
	}
	if in.Print.Ams != nil {
	    var amsChanged bool
		og.Print.Ams, amsChanged = mergeAms(og.Print.Ams, in.Print.Ams)
		changed = changed || amsChanged
	}
	changed = mergePrimatives(og.Print, in.Print) || changed
	return og, changed
}

func mergeIpcam(og *mqtt.Ipcam, in *mqtt.Ipcam) (*mqtt.Ipcam, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.Ipcam{}
	}
	if in.IpcamDev != nil {
		if og.IpcamDev == nil {
			og.IpcamDev = new(string)
			changed = true
		}
		if *og.IpcamDev != *in.IpcamDev {
		    *og.IpcamDev = *in.IpcamDev
			changed = true
		}
	}
	if in.IpcamRecord != nil {
		if og.IpcamRecord == nil {
			og.IpcamRecord = new(string)
			changed = true
		}
		if *og.IpcamRecord != *in.IpcamRecord {
			og.IpcamRecord = in.IpcamRecord
			changed = true
		}
	}
	if in.Resolution != nil {
		if og.Resolution == nil {
			og.Resolution = new(string)
			changed = true
		}
		if *og.Resolution != *in.Resolution {
			og.Resolution = in.Resolution
			changed = true
		}
	}
	if in.Timelapse != nil {
		if og.Timelapse == nil {
			og.Timelapse = new(string)
			changed = true
		}
		if *og.Timelapse != *in.Timelapse {
			og.Timelapse = in.Timelapse
			changed = true
		}
	}
	return og, changed
}

func mergeUpload(og *mqtt.Upload, in *mqtt.Upload) (*mqtt.Upload, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.Upload{}
	}

	if in.Status != nil {
		if og.Status == nil {
			og.Status = new(string)
			changed = true
		}
		if *og.Status != *in.Status {
			og.Status = in.Status
			changed = true
		}
	}
	if in.Progress != nil {
		if og.Progress == nil {
			og.Progress = new(int)
			changed = true
		}
		if *og.Progress != *in.Progress {
			og.Progress = in.Progress
			changed = true
		}
	}
	if in.Message != nil {
		if og.Message == nil {
			og.Message = new(string)
			changed = true
		}
		if *og.Message != *in.Message {
			og.Message = in.Message
			changed = true
		}
	}
	return og, changed
}

func mergeUpgradeState(og *mqtt.UpgradeState, in *mqtt.UpgradeState) (*mqtt.UpgradeState, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.UpgradeState{}
	}

	if in.SequenceID != nil {
		if og.SequenceID == nil {
			og.SequenceID = new(int)
			changed = true
		}
		if *og.SequenceID != *in.SequenceID {
			og.SequenceID = in.SequenceID
			changed = true
		}
	}
	if in.Progress != nil {
		if og.Progress == nil {
			og.Progress = new(string)
			changed = true
		}
		if *og.Progress != *in.Progress {
			og.Progress = in.Progress
			changed = true
		}
	}
	if in.Status != nil {
		if og.Status == nil {
			og.Status = new(string)
			changed = true
		}
		if *og.Status != *in.Status {
			og.Status = in.Status
			changed = true
		}
	}
	if in.ConsistencyRequest != nil {
		if og.ConsistencyRequest == nil {
			og.ConsistencyRequest = new(bool)
			changed = true
		}
		if *og.ConsistencyRequest != *in.ConsistencyRequest {
			og.ConsistencyRequest = in.ConsistencyRequest
			changed = true
		}
	}
	if in.DisState != nil {
		if og.DisState == nil {
			og.DisState = new(int)
			changed = true
		}
		if *og.DisState != *in.DisState {
			og.DisState = in.DisState
			changed = true
		}
	}
	if in.ErrCode != nil {
		if og.ErrCode == nil {
			og.ErrCode = new(int)
			changed = true
		}
		if *og.ErrCode != *in.ErrCode {
			og.ErrCode = in.ErrCode
			changed = true
		}
	}
	if in.ForceUpgrade != nil {
		if og.ForceUpgrade == nil {
			og.ForceUpgrade = new(bool)
			changed = true
		}
		if *og.ForceUpgrade != *in.ForceUpgrade {
			og.ForceUpgrade = in.ForceUpgrade
			changed = true
		}
	}
	if in.Message != nil {
		if og.Message == nil {
			og.Message = new(string)
			changed = true
		}
		if *og.Message != *in.Message {
			og.Message = in.Message
			changed = true
		}
	}
	if in.Module != nil {
		if og.Module == nil {
			og.Module = new(string)
			changed = true
		}
		if *og.Module != *in.Module {
			og.Module = in.Module
			changed = true
		}
	}
	if in.NewVersionState != nil {
		if og.NewVersionState == nil {
			og.NewVersionState = new(int)
			changed = true
		}
		if *og.NewVersionState != *in.NewVersionState {
			og.NewVersionState = in.NewVersionState
			changed = true
		}
	}
	if in.CurStateCode != nil {
		if og.CurStateCode == nil {
			og.CurStateCode = new(int)
			changed = true
		}
		if *og.CurStateCode != *in.CurStateCode {
			og.CurStateCode = in.CurStateCode
			changed = true
		}
	}
	if in.NewVerList != nil {
		if og.NewVerList == nil {
			og.NewVerList = new([]any)
			changed = true
		}
		if !reflect.DeepEqual(*og.NewVerList, *in.NewVerList) {
			og.NewVerList = in.NewVerList
			changed = true
		}
	}
	return og, changed
}

func mergeLightsReport(og *[]mqtt.LightsReport, in *[]mqtt.LightsReport) (*[]mqtt.LightsReport, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		lr := []mqtt.LightsReport{}
		og = &lr
	}
	if !reflect.DeepEqual(*og, *in) {
		*og = *in
		changed = true
	}
	return og, changed
}

func mergeVtTray(og *mqtt.VtTray, in *mqtt.VtTray) (*mqtt.VtTray, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.VtTray{}
	}
	if in.ID != nil {
		if og.ID == nil {
			og.ID = new(string)
			changed = true
		}
		if *og.ID != *in.ID {
			og.ID = in.ID
			changed = true
		}
	}
	if in.TagUID != nil {
		if og.TagUID == nil {
			og.TagUID = new(string)
			changed = true
		}
		if *og.TagUID != *in.TagUID {
			og.TagUID = in.TagUID
			changed = true
		}
	}
	if in.TrayIDName != nil {
		if og.TrayIDName == nil {
			og.TrayIDName = new(string)
			changed = true
		}
		if *og.TrayIDName != *in.TrayIDName {
			og.TrayIDName = in.TrayIDName
			changed = true
		}
	}
	if in.TrayInfoIdx != nil {
		if og.TrayInfoIdx == nil {
			og.TrayInfoIdx = new(string)
			changed = true
		}
		if *og.TrayInfoIdx != *in.TrayInfoIdx {
			og.TrayInfoIdx = in.TrayInfoIdx
			changed = true
		}
	}
	if in.TrayType != nil {
		if og.TrayType == nil {
			og.TrayType = new(string)
			changed = true
		}
		if *og.TrayType != *in.TrayType {
			og.TrayType = in.TrayType
			changed = true
		}
	}
	if in.TraySubBrands != nil {
		if og.TraySubBrands == nil {
			og.TraySubBrands = new(string)
			changed = true
		}
		if *og.TraySubBrands != *in.TraySubBrands {
			og.TraySubBrands = in.TraySubBrands
			changed = true
		}
	}
	if in.TrayColor != nil {
		if og.TrayColor == nil {
			og.TrayColor = new(string)
			changed = true
		}
		if *og.TrayColor != *in.TrayColor {
			og.TrayColor = in.TrayColor
			changed = true
		}
	}
	if in.TrayWeight != nil {
		if og.TrayWeight == nil {
			og.TrayWeight = new(string)
			changed = true
		}
		if *og.TrayWeight != *in.TrayWeight {
			og.TrayWeight = in.TrayWeight
			changed = true
		}
	}
	if in.TrayDiameter != nil {
		if og.TrayDiameter == nil {
			og.TrayDiameter = new(string)
			changed = true
		}
		if *og.TrayDiameter != *in.TrayDiameter {
			og.TrayDiameter = in.TrayDiameter
			changed = true
		}
	}
	if in.TrayTemp != nil {
		if og.TrayTemp == nil {
			og.TrayTemp = new(string)
			changed = true
		}
		if *og.TrayTemp != *in.TrayTemp {
			og.TrayTemp = in.TrayTemp
			changed = true
		}
	}
	if in.TrayTime != nil {
		if og.TrayTime == nil {
			og.TrayTime = new(string)
			changed = true
		}
		if *og.TrayTime != *in.TrayTime {
			og.TrayTime = in.TrayTime
			changed = true
		}
	}
	if in.BedTempType != nil {
		if og.BedTempType == nil {
			og.BedTempType = new(string)
			changed = true
		}
		if *og.BedTempType != *in.BedTempType {
			og.BedTempType = in.BedTempType
			changed = true
		}
	}
	if in.BedTemp != nil {
		if og.BedTemp == nil {
			og.BedTemp = new(string)
			changed = true
		}
		if *og.BedTemp != *in.BedTemp {
			og.BedTemp = in.BedTemp
			changed = true
		}
	}
	if in.NozzleTempMax != nil {
		if og.NozzleTempMax == nil {
			og.NozzleTempMax = new(string)
			changed = true
		}
		if *og.NozzleTempMax != *in.NozzleTempMax {
			og.NozzleTempMax = in.NozzleTempMax
			changed = true
		}
	}
	if in.NozzleTempMin != nil {
		if og.NozzleTempMin == nil {
			og.NozzleTempMin = new(string)
			changed = true
		}
		if *og.NozzleTempMin != *in.NozzleTempMin {
			og.NozzleTempMin = in.NozzleTempMin
			changed = true
		}
	}
	if in.XcamInfo != nil {
		if og.XcamInfo == nil {
			og.XcamInfo = new(string)
			changed = true
		}
		if *og.XcamInfo != *in.XcamInfo {
			og.XcamInfo = in.XcamInfo
			changed = true
		}
	}
	if in.Remain != nil {
		if og.Remain == nil {
			og.Remain = new(int)
			changed = true
		}
		if *og.Remain != *in.Remain {
			og.Remain = in.Remain
			changed = true
		}
	}
	if in.K != nil {
		if og.K == nil {
			og.K = new(float64)
			changed = true
		}
		if *og.K != *in.K {
			og.K = in.K
			changed = true
		}
	}
	if in.N != nil {
		if og.N == nil {
			og.N = new(int)
			changed = true
		}
		if *og.N != *in.N {
			og.N = in.N
			changed = true
		}
	}
	if in.CaliIdx != nil {
		if og.CaliIdx == nil {
			og.CaliIdx = new(int)
			changed = true
		}
		if *og.CaliIdx != *in.CaliIdx {
			og.CaliIdx = in.CaliIdx
			changed = true
		}
	}
	return og, changed
}

func mergeAmsInner(og *[]mqtt.AmsInner, in *[]mqtt.AmsInner) (*[]mqtt.AmsInner, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		ai := []mqtt.AmsInner{}
		og = &ai
	}
	if !reflect.DeepEqual(*og, *in) {
		*og = *in
		changed = true
	}
	return og, changed
}

func mergeAms(og *mqtt.Ams, in *mqtt.Ams) (*mqtt.Ams, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.Ams{}
	}

	var amsInnerChanged bool
	og.Ams, amsInnerChanged = mergeAmsInner(og.Ams, in.Ams)
    changed = changed || amsInnerChanged

	if in.AmsExistBits != nil {
		if og.AmsExistBits == nil {
			og.AmsExistBits = new(string)
			changed = true
		}
		if *og.AmsExistBits != *in.AmsExistBits {
			og.AmsExistBits = in.AmsExistBits
			changed = true
		}
	}
	if in.TrayExistBits != nil {
		if og.TrayExistBits == nil {
			og.TrayExistBits = new(string)
			changed = true
		}
		if *og.TrayExistBits != *in.TrayExistBits {
			og.TrayExistBits = in.TrayExistBits
			changed = true
		}
	}
	if in.TrayIsBblBits != nil {
		if og.TrayIsBblBits == nil {
			og.TrayIsBblBits = new(string)
			changed = true
		}
		if *og.TrayIsBblBits != *in.TrayIsBblBits {
			og.TrayIsBblBits = in.TrayIsBblBits
			changed = true
		}
	}
	if in.TrayTar != nil {
		if og.TrayTar == nil {
			og.TrayTar = new(string)
			changed = true
		}
		if *og.TrayTar != *in.TrayTar {
			og.TrayTar = in.TrayTar
			changed = true
		}
	}
	if in.TrayNow != nil {
		if og.TrayNow == nil {
			og.TrayNow = new(string)
			changed = true
		}
		if *og.TrayNow != *in.TrayNow {
			og.TrayNow = in.TrayNow
			changed = true
		}
	}
	if in.TrayPre != nil {
		if og.TrayPre == nil {
			og.TrayPre = new(string)
			changed = true
		}
		if *og.TrayPre != *in.TrayPre {
			og.TrayPre = in.TrayPre
			changed = true
		}
	}
	if in.TrayReadDoneBits != nil {
		if og.TrayReadDoneBits == nil {
			og.TrayReadDoneBits = new(string)
			changed = true
		}
		if *og.TrayReadDoneBits != *in.TrayReadDoneBits {
			og.TrayReadDoneBits = in.TrayReadDoneBits
			changed = true
		}
	}
	if in.TrayReadingBits != nil {
		if og.TrayReadingBits == nil {
			og.TrayReadingBits = new(string)
			changed = true
		}
		if *og.TrayReadingBits != *in.TrayReadingBits {
			og.TrayReadingBits = in.TrayReadingBits
			changed = true
		}
	}
	if in.Version != nil {
		if og.Version == nil {
			og.Version = new(int)
			changed = true
		}
		if *og.Version != *in.Version {
			og.Version = in.Version
			changed = true
		}
	}
	if in.InsertFlag != nil {
		if og.InsertFlag == nil {
			og.InsertFlag = new(bool)
			changed = true
		}
		if *og.InsertFlag != *in.InsertFlag {
			og.InsertFlag = in.InsertFlag
			changed = true
		}
	}
	if in.PowerOnFlag != nil {
		if og.PowerOnFlag == nil {
			og.PowerOnFlag = new(bool)
			changed = true
		}
		if *og.PowerOnFlag != *in.PowerOnFlag {
			og.PowerOnFlag = in.PowerOnFlag
			changed = true
		}
	}
	return og, changed
}

func mergeOnline(og *mqtt.Online, in *mqtt.Online) (*mqtt.Online, bool) {
	changed := false
	if in == nil {
		return og, changed
	}
	if og == nil {
		og = &mqtt.Online{}
	}

	if in.Ahb != nil {
		if og.Ahb == nil {
			og.Ahb = new(bool)
			changed = true
		}
		if *og.Ahb != *in.Ahb {
			og.Ahb = in.Ahb
			changed = true
		}
	}
	if in.Rfid != nil {
		if og.Rfid == nil {
			og.Rfid = new(bool)
			changed = true
		}
		if *og.Rfid != *in.Rfid {
			og.Rfid = in.Rfid
			changed = true
		}
	}
	if in.Version != nil {
		if og.Version == nil {
			og.Version = new(int)
			changed = true
		}
		if *og.Version != *in.Version {
			og.Version = in.Version
			changed = true
		}
	}
	return og, changed
}

func mergePrimatives(og *mqtt.Print, in *mqtt.Print) bool {
	changed := false
	if in == nil {
		return changed
	}
	if in.AmsRfidStatus != nil {
		if og.AmsRfidStatus == nil {
			og.AmsRfidStatus = new(int)
			changed = true
		}
		if *og.AmsRfidStatus != *in.AmsRfidStatus {
			og.AmsRfidStatus = in.AmsRfidStatus
			changed = true
		}
	}
	if in.AmsStatus != nil {
		if og.AmsStatus == nil {
			og.AmsStatus = new(int)
			changed = true
		}
		if *og.AmsStatus != *in.AmsStatus {
			og.AmsStatus = in.AmsStatus
			changed = true
		}
	}
	if in.BedTargetTemper != nil {
		if og.BedTargetTemper == nil {
			og.BedTargetTemper = new(int)
			changed = true
		}
		if *og.BedTargetTemper != *in.BedTargetTemper {
			og.BedTargetTemper = in.BedTargetTemper
			changed = true
		}
	}
	if in.BedTemper != nil {
		if og.BedTemper == nil {
			og.BedTemper = new(float64)
			changed = true
		}
		if *og.BedTemper != *in.BedTemper {
			og.BedTemper = in.BedTemper
			changed = true
		}
	}
	if in.BigFan1Speed != nil {
		if og.BigFan1Speed == nil {
			og.BigFan1Speed = new(string)
			changed = true
		}
		if *og.BigFan1Speed != *in.BigFan1Speed {
			og.BigFan1Speed = in.BigFan1Speed
			changed = true
		}
	}
	if in.BigFan2Speed != nil {
		if og.BigFan2Speed == nil {
			og.BigFan2Speed = new(string)
			changed = true
		}
		if *og.BigFan2Speed != *in.BigFan2Speed {
			og.BigFan2Speed = in.BigFan2Speed
			changed = true
		}
	}
	if in.CaliVersion != nil {
		if og.CaliVersion == nil {
			og.CaliVersion = new(int)
			changed = true
		}
		if *og.CaliVersion != *in.CaliVersion {
			og.CaliVersion = in.CaliVersion
			changed = true
		}
	}
	if in.ChamberTemper != nil {
		if og.ChamberTemper == nil {
			og.ChamberTemper = new(int)
			changed = true
		}
		if *og.ChamberTemper != *in.ChamberTemper {
			og.ChamberTemper = in.ChamberTemper
			changed = true
		}
	}
	if in.Command != nil {
		if og.Command == nil {
			og.Command = new(string)
			changed = true
		}
		if *og.Command != *in.Command {
			og.Command = in.Command
			changed = true
		}
	}
	if in.CoolingFanSpeed != nil {
		if og.CoolingFanSpeed == nil {
			og.CoolingFanSpeed = new(string)
			changed = true
		}
		if *og.CoolingFanSpeed != *in.CoolingFanSpeed {
			og.CoolingFanSpeed = in.CoolingFanSpeed
			changed = true
		}
	}
	if in.FanGear != nil {
		if og.FanGear == nil {
			og.FanGear = new(int)
			changed = true
		}
		if *og.FanGear != *in.FanGear {
			og.FanGear = in.FanGear
			changed = true
		}
	}
	if in.FilamBak != nil {
		if og.FilamBak == nil {
			og.FilamBak = new([]any)
			changed = true
		}
		if !reflect.DeepEqual(*og.FilamBak, *in.FilamBak) {
			og.FilamBak = in.FilamBak
			changed = true
		}
	}
	if in.ForceUpgrade != nil {
		if og.ForceUpgrade == nil {
			og.ForceUpgrade = new(bool)
			changed = true
		}
		if *og.ForceUpgrade != *in.ForceUpgrade {
			og.ForceUpgrade = in.ForceUpgrade
			changed = true
		}
	}
	if in.GcodeFile != nil {
		if og.GcodeFile == nil {
			og.GcodeFile = new(string)
			changed = true
		}
		if *og.GcodeFile != *in.GcodeFile {
			og.GcodeFile = in.GcodeFile
			changed = true
		}
	}
	if in.GcodeFilePreparePercent != nil {
		if og.GcodeFilePreparePercent == nil {
			og.GcodeFilePreparePercent = new(string)
			changed = true
		}
		if *og.GcodeFilePreparePercent != *in.GcodeFilePreparePercent {
			og.GcodeFilePreparePercent = in.GcodeFilePreparePercent
			changed = true
		}
	}
	if in.GcodeState != nil {
		if og.GcodeState == nil {
			og.GcodeState = new(string)
			changed = true
		}
		if *og.GcodeState != *in.GcodeState {
			og.GcodeState = in.GcodeState
			changed = true
		}
	}
	if in.HeatbreakFanSpeed != nil {
		if og.HeatbreakFanSpeed == nil {
			og.HeatbreakFanSpeed = new(string)
			changed = true
		}
		if *og.HeatbreakFanSpeed != *in.HeatbreakFanSpeed {
			og.HeatbreakFanSpeed = in.HeatbreakFanSpeed
			changed = true
		}
	}
	if in.Hms != nil {
		if og.Hms == nil {
			og.Hms = new([]any)
			changed = true
		}
		if !reflect.DeepEqual(*og.Hms, *in.Hms) {
			og.Hms = in.Hms
			changed = true
		}
	}
	if in.HomeFlag != nil {
		if og.HomeFlag == nil {
			og.HomeFlag = new(int)
			changed = true
		}
		if *og.HomeFlag != *in.HomeFlag {
			og.HomeFlag = in.HomeFlag
			changed = true
		}
	}
	if in.HwSwitchState != nil {
		if og.HwSwitchState == nil {
			og.HwSwitchState = new(int)
			changed = true
		}
		if *og.HwSwitchState != *in.HwSwitchState {
			og.HwSwitchState = in.HwSwitchState
			changed = true
		}
	}
	if in.LayerNum != nil {
		if og.LayerNum == nil {
			og.LayerNum = new(int)
			changed = true
		}
		if *og.LayerNum != *in.LayerNum {
			og.LayerNum = in.LayerNum
			changed = true
		}
	}
	if in.Lifecycle != nil {
		if og.Lifecycle == nil {
			og.Lifecycle = new(string)
			changed = true
		}
		if *og.Lifecycle != *in.Lifecycle {
			og.Lifecycle = in.Lifecycle
			changed = true
		}
	}
	if in.McPercent != nil {
		if og.McPercent == nil {
			og.McPercent = new(int)
			changed = true
		}
		if *og.McPercent != *in.McPercent {
			og.McPercent = in.McPercent
			changed = true
		}
	}
	if in.McPrintLineNumber != nil {
		if og.McPrintLineNumber == nil {
			og.McPrintLineNumber = new(string)
			changed = true
		}
		if *og.McPrintLineNumber != *in.McPrintLineNumber {
			og.McPrintLineNumber = in.McPrintLineNumber
			changed = true
		}
	}
	if in.McPrintStage != nil {
		if og.McPrintStage == nil {
			og.McPrintStage = new(string)
			changed = true
		}
		if *og.McPrintStage != *in.McPrintStage {
			og.McPrintStage = in.McPrintStage
			changed = true
		}
	}
	if in.McPrintSubStage != nil {
		if og.McPrintSubStage == nil {
			og.McPrintSubStage = new(int)
			changed = true
		}
		if *og.McPrintSubStage != *in.McPrintSubStage {
			og.McPrintSubStage = in.McPrintSubStage
			changed = true
		}
	}
	if in.McRemainingTime != nil {
		if og.McRemainingTime == nil {
			og.McRemainingTime = new(int)
			changed = true
		}
		if *og.McRemainingTime != *in.McRemainingTime {
			og.McRemainingTime = in.McRemainingTime
			changed = true
		}
	}
	if in.MessProductionState != nil {
		if og.MessProductionState == nil {
			og.MessProductionState = new(string)
			changed = true
		}
		if *og.MessProductionState != *in.MessProductionState {
			og.MessProductionState = in.MessProductionState
			changed = true
		}
	}
	if in.Msg != nil {
		if og.Msg == nil {
			og.Msg = new(int)
			changed = true
		}
		if *og.Msg != *in.Msg {
			og.Msg = in.Msg
			changed = true
		}
	}
	if in.NozzleDiameter != nil {
		if og.NozzleDiameter == nil {
			og.NozzleDiameter = new(string)
			changed = true
		}
		if *og.NozzleDiameter != *in.NozzleDiameter {
			og.NozzleDiameter = in.NozzleDiameter
			changed = true
		}
	}
	if in.NozzleTargetTemper != nil {
		if og.NozzleTargetTemper == nil {
			og.NozzleTargetTemper = new(int)
			changed = true
		}
		if *og.NozzleTargetTemper != *in.NozzleTargetTemper {
			og.NozzleTargetTemper = in.NozzleTargetTemper
			changed = true
		}
	}
	if in.NozzleTemper != nil {
		if og.NozzleTemper == nil {
			og.NozzleTemper = new(float64)
			changed = true
		}
		if *og.NozzleTemper != *in.NozzleTemper {
			og.NozzleTemper = in.NozzleTemper
			changed = true
		}
	}
	if in.NozzleType != nil {
		if og.NozzleType == nil {
			og.NozzleType = new(string)
			changed = true
		}
		if *og.NozzleType != *in.NozzleType {
			og.NozzleType = in.NozzleType
			changed = true
		}
	}
	if in.PrintError != nil {
		if og.PrintError == nil {
			og.PrintError = new(int)
			changed = true
		}
		if *og.PrintError != *in.PrintError {
			og.PrintError = in.PrintError
			changed = true
		}
	}
	if in.PrintType != nil {
		if og.PrintType == nil {
			og.PrintType = new(string)
			changed = true
		}
		if *og.PrintType != *in.PrintType {
			og.PrintType = in.PrintType
			changed = true
		}
	}
	if in.ProfileID != nil {
		if og.ProfileID == nil {
			og.ProfileID = new(string)
			changed = true
		}
		if *og.ProfileID != *in.ProfileID {
			og.ProfileID = in.ProfileID
			changed = true
		}
	}
	if in.ProjectID != nil {
		if og.ProjectID == nil {
			og.ProjectID = new(string)
			changed = true
		}
		if *og.ProjectID != *in.ProjectID {
			og.ProjectID = in.ProjectID
			changed = true
		}
	}
	if in.QueueEst != nil {
		if og.QueueEst == nil {
			og.QueueEst = new(int)
			changed = true
		}
		if *og.QueueEst != *in.QueueEst {
			og.QueueEst = in.QueueEst
			changed = true
		}
	}
	if in.QueueNumber != nil {
		if og.QueueNumber == nil {
			og.QueueNumber = new(int)
			changed = true
		}
		if *og.QueueNumber != *in.QueueNumber {
			og.QueueNumber = in.QueueNumber
			changed = true
		}
	}
	if in.QueueSts != nil {
		if og.QueueSts == nil {
			og.QueueSts = new(int)
			changed = true
		}
		if *og.QueueSts != *in.QueueSts {
			og.QueueSts = in.QueueSts
			changed = true
		}
	}
	if in.SObj != nil {
		if og.SObj == nil {
			og.SObj = new([]any)
			changed = true
		}
		if !reflect.DeepEqual(*og.SObj, *in.SObj) {
			og.SObj = in.SObj
			changed = true
		}
	}
	if in.Sdcard != nil {
		if og.Sdcard == nil {
			og.Sdcard = new(bool)
			changed = true
		}
		if *og.Sdcard != *in.Sdcard {
			og.Sdcard = in.Sdcard
			changed = true
		}
	}
	if in.SequenceID != nil {
		if og.SequenceID == nil {
			og.SequenceID = new(string)
			changed = true
		}
		if *og.SequenceID != *in.SequenceID {
			og.SequenceID = in.SequenceID
			changed = true
		}
	}
	if in.SpdLvl != nil {
		if og.SpdLvl == nil {
			og.SpdLvl = new(int)
			changed = true
		}
		if *og.SpdLvl != *in.SpdLvl {
			og.SpdLvl = in.SpdLvl
			changed = true
		}
	}
	if in.SpdMag != nil {
		if og.SpdMag == nil {
			og.SpdMag = new(int)
			changed = true
		}
		if *og.SpdMag != *in.SpdMag {
			og.SpdMag = in.SpdMag
			changed = true
		}
	}
	if in.Stg != nil {
		if og.Stg == nil {
			og.Stg = new([]int)
			changed = true
		}
		if !reflect.DeepEqual(*og.Stg, *in.Stg) {
			og.Stg = in.Stg
			changed = true
		}
	}
	if in.StgCur != nil {
		if og.StgCur == nil {
			og.StgCur = new(int)
			changed = true
		}
		if *og.StgCur != *in.StgCur {
			og.StgCur = in.StgCur
			changed = true
		}
	}
	if in.SubtaskID != nil {
		if og.SubtaskID == nil {
			og.SubtaskID = new(string)
			changed = true
		}
		if *og.SubtaskID != *in.SubtaskID {
			og.SubtaskID = in.SubtaskID
			changed = true
		}
	}
	if in.SubtaskName != nil {
		if og.SubtaskName == nil {
			og.SubtaskName = new(string)
			changed = true
		}
		if *og.SubtaskName != *in.SubtaskName {
			og.SubtaskName = in.SubtaskName
			changed = true
		}
	}
	if in.TaskID != nil {
		if og.TaskID == nil {
			og.TaskID = new(string)
			changed = true
		}
		if *og.TaskID != *in.TaskID {
			og.TaskID = in.TaskID
			changed = true
		}
	}
	if in.TotalLayerNum != nil {
		if og.TotalLayerNum == nil {
			og.TotalLayerNum = new(int)
			changed = true
		}
		if *og.TotalLayerNum != *in.TotalLayerNum {
			og.TotalLayerNum = in.TotalLayerNum
			changed = true
		}
	}
	if in.WifiSignal != nil {
		if og.WifiSignal == nil {
			og.WifiSignal = new(string)
			changed = true
		}
		if *og.WifiSignal != *in.WifiSignal {
			og.WifiSignal = in.WifiSignal
			changed = true
		}
	}
	return changed
}
