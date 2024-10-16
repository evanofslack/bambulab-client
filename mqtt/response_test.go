package mqtt

import (
	"testing"

    "encoding/json"
	"github.com/stretchr/testify/assert"
)

func TestUnmarshalMessage(t *testing.T) {
	raw := `
	{
        "print": {
                "ipcam": {
                        "ipcam_dev": "1",
                        "ipcam_record": "enable",
                        "timelapse": "enable",
                        "resolution": "",
                        "tutk_server": "enable",
                        "mode_bits": 3
                },
                "upload": {
                        "status": "idle",
                        "progress": 0,
                        "message": ""
                },
                "nozzle_temper": 219.1875,
                "nozzle_target_temper": 220,
                "bed_temper": 54.875,
                "bed_target_temper": 55,
                "chamber_temper": 5,
                "mc_print_stage": "2",
                "heatbreak_fan_speed": "15",
                "cooling_fan_speed": "15",
                "big_fan1_speed": "15",
                "big_fan2_speed": "8",
                "mc_percent": 42,
                "mc_remaining_time": 147,
                "ams_status": 768,
                "ams_rfid_status": 2,
                "hw_switch_state": 1,
                "spd_mag": 100,
                "spd_lvl": 2,
                "print_error": 0,
                "lifecycle": "product",
                "wifi_signal": "-62dBm",
                "gcode_state": "RUNNING",
                "gcode_file_prepare_percent": "100",
                "queue_number": 0,
                "queue_total": 0,
                "queue_est": 0,
                "queue_sts": 0,
                "project_id": "144131202",
                "profile_id": "138104297",
                "task_id": "287384806",
                "subtask_id": "287384807",
                "subtask_name": "Keychain",
                "gcode_file": "Keychain.3mf",
                "stg": [
                        2,
                        14,
                        1
                ],
                "stg_cur": 0,
                "print_type": "cloud",
                "home_flag": 24266039,
                "mc_print_line_number": "156230",
                "mc_print_sub_stage": 0,
                "sdcard": true,
                "force_upgrade": false,
                "mess_production_state": "active",
                "layer_num": 66,
                "total_layer_num": 182,
                "s_obj": [],
                "filam_bak": [],
                "fan_gear": 9895935,
                "nozzle_diameter": "0.4",
                "nozzle_type": "stainless_steel",
                "cali_version": 0,
                "upgrade_state": {
                        "sequence_id": 0,
                        "progress": "",
                        "status": "",
                        "consistency_request": false,
                        "dis_state": 0,
                        "err_code": 0,
                        "force_upgrade": false,
                        "message": "0%, 0B/s",
                        "module": "",
                        "new_version_state": 2,
                        "cur_state_code": 1,
                        "new_ver_list": []
                },
                "hms": [],
                "online": {
                        "ahb": false,
                        "rfid": false,
                        "version": 667260902
                },
                "ams": {
                        "ams": [
                                {
                                        "id": "0",
                                        "humidity": "5",
                                        "temp": "0.0",
                                        "tray": [
                                                {
                                                        "id": "0",
                                                        "remain": -1,
                                                        "k": 0.019999999552965164,
                                                        "n": 1,
                                                        "cali_idx": -1,
                                                        "tag_uid": "1AB20B7300000100",
                                                        "tray_id_name": "A01-K1",
                                                        "tray_info_idx": "GFA01",
                                                        "tray_type": "PLA",
                                                        "tray_sub_brands": "PLA Matte",
                                                        "tray_color": "000000FF",
                                                        "tray_weight": "1000",
                                                        "tray_diameter": "1.75",
                                                        "tray_temp": "55",
                                                        "tray_time": "8",
                                                        "bed_temp_type": "0",
                                                        "bed_temp": "0",
                                                        "nozzle_temp_max": "230",
                                                        "nozzle_temp_min": "190",
                                                        "xcam_info": "A4388813E803E803CDCC4C3F",
                                                        "tray_uuid": "3D4338337A2B41F6949E72558F62852D",
                                                        "ctype": 0,
                                                        "cols": [
                                                                "000000FF"
                                                        ]
                                                },
                                                {
                                                        "id": "1",
                                                        "remain": -1,
                                                        "k": 0.019999999552965164,
                                                        "n": 1,
                                                        "cali_idx": -1,
                                                        "tag_uid": "81D2DC0F00080100",
                                                        "tray_id_name": "A00-A0",
                                                        "tray_info_idx": "GFA00",
                                                        "tray_type": "PLA",
                                                        "tray_sub_brands": "PLA Basic",
                                                        "tray_color": "FF6A13FF",
                                                        "tray_weight": "250",
                                                        "tray_diameter": "1.75",
                                                        "tray_temp": "55",
                                                        "tray_time": "8",
                                                        "bed_temp_type": "0",
                                                        "bed_temp": "0",
                                                        "nozzle_temp_max": "230",
                                                        "nozzle_temp_min": "190",
                                                        "xcam_info": "8813100EE803E8039A99193F",
                                                        "tray_uuid": "2516B3FB9C2E495393776814029FB298",
                                                        "ctype": 0,
                                                        "cols": [
                                                                "FF6A13FF"
                                                        ]
                                                },
                                                {
                                                        "id": "2",
                                                        "remain": -1,
                                                        "k": 0.019999999552965164,
                                                        "n": 1,
                                                        "cali_idx": -1,
                                                        "tag_uid": "BAA36E7400000100",
                                                        "tray_id_name": "A01-B6",
                                                        "tray_info_idx": "GFA01",
                                                        "tray_type": "PLA",
                                                        "tray_sub_brands": "PLA Matte",
                                                        "tray_color": "042F56FF",
                                                        "tray_weight": "1000",
                                                        "tray_diameter": "1.75",
                                                        "tray_temp": "55",
                                                        "tray_time": "8",
                                                        "bed_temp_type": "1",
                                                        "bed_temp": "35",
                                                        "nozzle_temp_max": "230",
                                                        "nozzle_temp_min": "190",
                                                        "xcam_info": "AC0DAC0DE803E8030000803F",
                                                        "tray_uuid": "0BAAF7ABAFD947EFA10A17FB9A88D5F5",
                                                        "ctype": 0,
                                                        "cols": [
                                                                "042F56FF"
                                                        ]
                                                },
                                                {
                                                        "id": "3",
                                                        "remain": -1,
                                                        "k": 0.019999999552965164,
                                                        "n": 1,
                                                        "cali_idx": -1,
                                                        "tag_uid": "517B940300080100",
                                                        "tray_id_name": "A00-G1",
                                                        "tray_info_idx": "GFA00",
                                                        "tray_type": "PLA",
                                                        "tray_sub_brands": "PLA Basic",
                                                        "tray_color": "00AE42FF",
                                                        "tray_weight": "250",
                                                        "tray_diameter": "1.75",
                                                        "tray_temp": "55",
                                                        "tray_time": "8",
                                                        "bed_temp_type": "0",
                                                        "bed_temp": "0",
                                                        "nozzle_temp_max": "230",
                                                        "nozzle_temp_min": "190",
                                                        "xcam_info": "88138813E803E8039A99193F",
                                                        "tray_uuid": "916A162886554592B479D29E7964F1C0",
                                                        "ctype": 0,
                                                        "cols": [
                                                                "00AE42FF"
                                                        ]
                                                }
                                        ]
                                }
                        ],
                        "ams_exist_bits": "1",
                        "tray_exist_bits": "f",
                        "tray_is_bbl_bits": "f",
                        "tray_tar": "1",
                        "tray_now": "1",
                        "tray_pre": "1",
                        "tray_read_done_bits": "f",
                        "tray_reading_bits": "0",
                        "version": 144,
                        "insert_flag": true,
                        "power_on_flag": false
                },
                "vt_tray": {
                        "id": "254",
                        "tag_uid": "0000000000000000",
                        "tray_id_name": "",
                        "tray_info_idx": "",
                        "tray_type": "",
                        "tray_sub_brands": "",
                        "tray_color": "00000000",
                        "tray_weight": "0",
                        "tray_diameter": "0.00",
                        "tray_temp": "0",
                        "tray_time": "0",
                        "bed_temp_type": "0",
                        "bed_temp": "0",
                        "nozzle_temp_max": "0",
                        "nozzle_temp_min": "0",
                        "xcam_info": "000000000000000000000000",
                        "tray_uuid": "00000000000000000000000000000000",
                        "remain": 0,
                        "k": 0.019999999552965164,
                        "n": 1,
                        "cali_idx": -1
                },
                "lights_report": [
                        {
                                "node": "chamber_light",
                                "mode": "on"
                        }
                ],
                "command": "push_status",
                "msg": 0,
                "sequence_id": "19796"
        }
    }`
    var m Message
    err := json.Unmarshal([]byte(raw), &m)
    assert.Nil(t, err)

    us := m.Print.Upload.Status
    assert.NotNil(t, us)
    assert.Equal(t, "idle", *us)

    amsp := m.Print.Ams.Ams
    assert.NotNil(t, amsp)
    amss := *amsp
    ams1 := amss[0]
    assert.Equal(t, "5", *ams1.Humidity)

    tray := *ams1.Tray
    tray1 := tray[0]
    subbrand := tray1.TraySubBrands
    assert.Equal(t, "PLA Matte", *subbrand)
}
