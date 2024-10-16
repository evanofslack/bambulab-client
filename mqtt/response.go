package mqtt

type Message struct {
	Print *Print `json:"print,omitempty"`
}

type Print struct {
	Ams                     *Ams            `json:"ams,omitempty"`
	Ipcam                   *Ipcam          `json:"ipcam,omitempty"`
	LightsReport            *[]LightsReport `json:"lights_report,omitempty"`
	Online                  *Online         `json:"online,omitempty"`
	UpgradeState            *UpgradeState   `json:"upgrade_state,omitempty"`
	Upload                  *Upload         `json:"upload,omitempty"`
	VtTray                  *VtTray         `json:"vt_tray,omitempty"`
	AmsRfidStatus           *int            `json:"ams_rfid_status,omitempty"`
	AmsStatus               *int            `json:"ams_status,omitempty"`
	BedTargetTemper         *int            `json:"bed_target_temper,omitempty"`
	BedTemper               *float64        `json:"bed_temper,omitempty"`
	BigFan1Speed            *string         `json:"big_fan1_speed,omitempty"`
	BigFan2Speed            *string         `json:"big_fan2_speed,omitempty"`
	CaliVersion             *int            `json:"cali_version,omitempty"`
	ChamberTemper           *int            `json:"chamber_temper,omitempty"`
	Command                 *string         `json:"command,omitempty"`
	CoolingFanSpeed         *string         `json:"cooling_fan_speed,omitempty"`
	FanGear                 *int            `json:"fan_gear,omitempty"`
	FilamBak                *[]any          `json:"filam_bak,omitempty"`
	ForceUpgrade            *bool           `json:"force_upgrade,omitempty"`
	GcodeFile               *string         `json:"gcode_file,omitempty"`
	GcodeFilePreparePercent *string         `json:"gcode_file_prepare_percent,omitempty"`
	GcodeState              *string         `json:"gcode_state,omitempty"`
	HeatbreakFanSpeed       *string         `json:"heatbreak_fan_speed,omitempty"`
	Hms                     *[]any          `json:"hms,omitempty"`
	HomeFlag                *int            `json:"home_flag,omitempty"`
	HwSwitchState           *int            `json:"hw_switch_state,omitempty"`
	LayerNum                *int            `json:"layer_num,omitempty"`
	Lifecycle               *string         `json:"lifecycle,omitempty"`
	McPercent               *int            `json:"mc_percent,omitempty"`
	McPrintLineNumber       *string         `json:"mc_print_line_number,omitempty"`
	McPrintStage            *string         `json:"mc_print_stage,omitempty"`
	McPrintSubStage         *int            `json:"mc_print_sub_stage,omitempty"`
	McRemainingTime         *int            `json:"mc_remaining_time,omitempty"`
	MessProductionState     *string         `json:"mess_production_state,omitempty"`
	Msg                     *int            `json:"msg,omitempty"`
	NozzleDiameter          *string         `json:"nozzle_diameter,omitempty"`
	NozzleTargetTemper      *int            `json:"nozzle_target_temper,omitempty"`
	NozzleTemper            *float64        `json:"nozzle_temper,omitempty"`
	NozzleType              *string         `json:"nozzle_type,omitempty"`
	PrintError              *int            `json:"print_error,omitempty"`
	PrintType               *string         `json:"print_type,omitempty"`
	ProfileID               *string         `json:"profile_id,omitempty"`
	ProjectID               *string         `json:"project_id,omitempty"`
	QueueEst                *int            `json:"queue_est,omitempty"`
	QueueNumber             *int            `json:"queue_number,omitempty"`
	QueueSts                *int            `json:"queue_sts,omitempty"`
	QueueTotal              *int            `json:"queue_total,omitempty"`
	SObj                    *[]any          `json:"s_obj,omitempty"`
	Sdcard                  *bool           `json:"sdcard,omitempty"`
	SequenceID              *string         `json:"sequence_id,omitempty"`
	SpdLvl                  *int            `json:"spd_lvl,omitempty"`
	SpdMag                  *int            `json:"spd_mag,omitempty"`
	Stg                     *[]int          `json:"stg,omitempty"`
	StgCur                  *int            `json:"stg_cur,omitempty"`
	SubtaskID               *string         `json:"subtask_id,omitempty"`
	SubtaskName             *string         `json:"subtask_name,omitempty"`
	TaskID                  *string         `json:"task_id,omitempty"`
	TotalLayerNum           *int            `json:"total_layer_num,omitempty"`
	WifiSignal              *string         `json:"wifi_signal,omitempty"`
}

type Ipcam struct {
	IpcamDev    *string `json:"ipcam_dev,omitempty"`
	IpcamRecord *string `json:"ipcam_record,omitempty"`
	Timelapse   *string `json:"timelapse,omitempty"`
	Resolution  *string `json:"resolution,omitempty"`
	TutkServer  *string `json:"tutk_server,omitempty"`
	ModeBits    *int    `json:"mode_bits,omitempty"`
}

type Upload struct {
	Status   *string `json:"status,omitempty"`
	Progress *int    `json:"progress,omitempty"`
	Message  *string `json:"message,omitempty"`
}

type UpgradeState struct {
	SequenceID         *int    `json:"sequence_id,omitempty"`
	Progress           *string `json:"progress,omitempty"`
	Status             *string `json:"status,omitempty"`
	ConsistencyRequest *bool   `json:"consistency_request,omitempty"`
	DisState           *int    `json:"dis_state,omitempty"`
	ErrCode            *int    `json:"err_code,omitempty"`
	ForceUpgrade       *bool   `json:"force_upgrade,omitempty"`
	Message            *string `json:"message,omitempty"`
	Module             *string `json:"module,omitempty"`
	NewVersionState    *int    `json:"new_version_state,omitempty"`
	CurStateCode       *int    `json:"cur_state_code,omitempty"`
	NewVerList         *[]any  `json:"new_ver_list,omitempty"`
}

type Online struct {
	Ahb     *bool `json:"ahb,omitempty"`
	Rfid    *bool `json:"rfid,omitempty"`
	Version *int  `json:"version,omitempty"`
}

type Tray struct {
	ID            *string   `json:"id,omitempty"`
	Remain        *int      `json:"remain,omitempty"`
	K             *float64  `json:"k,omitempty"`
	N             *int      `json:"n,omitempty"`
	CaliIdx       *int      `json:"cali_idx,omitempty"`
	TagUID        *string   `json:"tag_uid,omitempty"`
	TrayIDName    *string   `json:"tray_id_name,omitempty"`
	TrayInfoIdx   *string   `json:"tray_info_idx,omitempty"`
	TrayType      *string   `json:"tray_type,omitempty"`
	TraySubBrands *string   `json:"tray_sub_brands,omitempty"`
	TrayColor     *string   `json:"tray_color,omitempty"`
	TrayWeight    *string   `json:"tray_weight,omitempty"`
	TrayDiameter  *string   `json:"tray_diameter,omitempty"`
	TrayTemp      *string   `json:"tray_temp,omitempty"`
	TrayTime      *string   `json:"tray_time,omitempty"`
	BedTempType   *string   `json:"bed_temp_type,omitempty"`
	BedTemp       *string   `json:"bed_temp,omitempty"`
	NozzleTempMax *string   `json:"nozzle_temp_max,omitempty"`
	NozzleTempMin *string   `json:"nozzle_temp_min,omitempty"`
	XcamInfo      *string   `json:"xcam_info,omitempty"`
	TrayUUID      *string   `json:"tray_uuid,omitempty"`
	Ctype         *int      `json:"ctype,omitempty"`
	Cols          *[]string `json:"cols,omitempty"`
}

type AmsInner struct {
	ID       *string `json:"id,omitempty"`
	Humidity *string `json:"humidity,omitempty"`
	Temp     *string `json:"temp,omitempty"`
	Tray     *[]Tray `json:"tray,omitempty"`
}

type Ams struct {
	Ams              *[]AmsInner `json:"ams,omitempty"`
	AmsExistBits     *string     `json:"ams_exist_bits,omitempty"`
	TrayExistBits    *string     `json:"tray_exist_bits,omitempty"`
	TrayIsBblBits    *string     `json:"tray_is_bbl_bits,omitempty"`
	TrayTar          *string     `json:"tray_tar,omitempty"`
	TrayNow          *string     `json:"tray_now,omitempty"`
	TrayPre          *string     `json:"tray_pre,omitempty"`
	TrayReadDoneBits *string     `json:"tray_read_done_bits,omitempty"`
	TrayReadingBits  *string     `json:"tray_reading_bits,omitempty"`
	Version          *int        `json:"version,omitempty"`
	InsertFlag       *bool       `json:"insert_flag,omitempty"`
	PowerOnFlag      *bool       `json:"power_on_flag,omitempty"`
}

type VtTray struct {
	ID            *string  `json:"id,omitempty"`
	TagUID        *string  `json:"tag_uid,omitempty"`
	TrayIDName    *string  `json:"tray_id_name,omitempty"`
	TrayInfoIdx   *string  `json:"tray_info_idx,omitempty"`
	TrayType      *string  `json:"tray_type,omitempty"`
	TraySubBrands *string  `json:"tray_sub_brands,omitempty"`
	TrayColor     *string  `json:"tray_color,omitempty"`
	TrayWeight    *string  `json:"tray_weight,omitempty"`
	TrayDiameter  *string  `json:"tray_diameter,omitempty"`
	TrayTemp      *string  `json:"tray_temp,omitempty"`
	TrayTime      *string  `json:"tray_time,omitempty"`
	BedTempType   *string  `json:"bed_temp_type,omitempty"`
	BedTemp       *string  `json:"bed_temp,omitempty"`
	NozzleTempMax *string  `json:"nozzle_temp_max,omitempty"`
	NozzleTempMin *string  `json:"nozzle_temp_min,omitempty"`
	XcamInfo      *string  `json:"xcam_info,omitempty"`
	TrayUUID      *string  `json:"tray_uuid,omitempty"`
	Remain        *int     `json:"remain,omitempty"`
	K             *float64 `json:"k,omitempty"`
	N             *int     `json:"n,omitempty"`
	CaliIdx       *int     `json:"cali_idx,omitempty"`
}

type LightsReport struct {
	Node *string `json:"node,omitempty"`
	Mode *string `json:"mode,omitempty"`
}
