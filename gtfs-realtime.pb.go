// Code generated by protoc-gen-go.
// source: gtfs-realtime.proto
// DO NOT EDIT!

/*
Package transit_realtime is a generated protocol buffer package.

It is generated from these files:
	gtfs-realtime.proto

It has these top-level messages:
	FeedMessage
	FeedHeader
	FeedEntity
	TripUpdate
	VehiclePosition
	Alert
	TimeRange
	Position
	TripDescriptor
	VehicleDescriptor
	EntitySelector
	TranslatedString
*/
package main

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type FeedHeader_Incrementality int32

const (
	FeedHeader_FULL_DATASET FeedHeader_Incrementality = 0
	FeedHeader_DIFFERENTIAL FeedHeader_Incrementality = 1
)

var FeedHeader_Incrementality_name = map[int32]string{
	0: "FULL_DATASET",
	1: "DIFFERENTIAL",
}
var FeedHeader_Incrementality_value = map[string]int32{
	"FULL_DATASET": 0,
	"DIFFERENTIAL": 1,
}

func (x FeedHeader_Incrementality) Enum() *FeedHeader_Incrementality {
	p := new(FeedHeader_Incrementality)
	*p = x
	return p
}
func (x FeedHeader_Incrementality) String() string {
	return proto.EnumName(FeedHeader_Incrementality_name, int32(x))
}
func (x *FeedHeader_Incrementality) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FeedHeader_Incrementality_value, data, "FeedHeader_Incrementality")
	if err != nil {
		return err
	}
	*x = FeedHeader_Incrementality(value)
	return nil
}

type TripUpdate_StopTimeUpdate_ScheduleRelationship int32

const (
	TripUpdate_StopTimeUpdate_SCHEDULED TripUpdate_StopTimeUpdate_ScheduleRelationship = 0
	TripUpdate_StopTimeUpdate_SKIPPED   TripUpdate_StopTimeUpdate_ScheduleRelationship = 1
	TripUpdate_StopTimeUpdate_NO_DATA   TripUpdate_StopTimeUpdate_ScheduleRelationship = 2
)

var TripUpdate_StopTimeUpdate_ScheduleRelationship_name = map[int32]string{
	0: "SCHEDULED",
	1: "SKIPPED",
	2: "NO_DATA",
}
var TripUpdate_StopTimeUpdate_ScheduleRelationship_value = map[string]int32{
	"SCHEDULED": 0,
	"SKIPPED":   1,
	"NO_DATA":   2,
}

func (x TripUpdate_StopTimeUpdate_ScheduleRelationship) Enum() *TripUpdate_StopTimeUpdate_ScheduleRelationship {
	p := new(TripUpdate_StopTimeUpdate_ScheduleRelationship)
	*p = x
	return p
}
func (x TripUpdate_StopTimeUpdate_ScheduleRelationship) String() string {
	return proto.EnumName(TripUpdate_StopTimeUpdate_ScheduleRelationship_name, int32(x))
}
func (x *TripUpdate_StopTimeUpdate_ScheduleRelationship) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TripUpdate_StopTimeUpdate_ScheduleRelationship_value, data, "TripUpdate_StopTimeUpdate_ScheduleRelationship")
	if err != nil {
		return err
	}
	*x = TripUpdate_StopTimeUpdate_ScheduleRelationship(value)
	return nil
}

type VehiclePosition_VehicleStopStatus int32

const (
	VehiclePosition_INCOMING_AT   VehiclePosition_VehicleStopStatus = 0
	VehiclePosition_STOPPED_AT    VehiclePosition_VehicleStopStatus = 1
	VehiclePosition_IN_TRANSIT_TO VehiclePosition_VehicleStopStatus = 2
)

var VehiclePosition_VehicleStopStatus_name = map[int32]string{
	0: "INCOMING_AT",
	1: "STOPPED_AT",
	2: "IN_TRANSIT_TO",
}
var VehiclePosition_VehicleStopStatus_value = map[string]int32{
	"INCOMING_AT":   0,
	"STOPPED_AT":    1,
	"IN_TRANSIT_TO": 2,
}

func (x VehiclePosition_VehicleStopStatus) Enum() *VehiclePosition_VehicleStopStatus {
	p := new(VehiclePosition_VehicleStopStatus)
	*p = x
	return p
}
func (x VehiclePosition_VehicleStopStatus) String() string {
	return proto.EnumName(VehiclePosition_VehicleStopStatus_name, int32(x))
}
func (x *VehiclePosition_VehicleStopStatus) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VehiclePosition_VehicleStopStatus_value, data, "VehiclePosition_VehicleStopStatus")
	if err != nil {
		return err
	}
	*x = VehiclePosition_VehicleStopStatus(value)
	return nil
}

type VehiclePosition_CongestionLevel int32

const (
	VehiclePosition_UNKNOWN_CONGESTION_LEVEL VehiclePosition_CongestionLevel = 0
	VehiclePosition_RUNNING_SMOOTHLY         VehiclePosition_CongestionLevel = 1
	VehiclePosition_STOP_AND_GO              VehiclePosition_CongestionLevel = 2
	VehiclePosition_CONGESTION               VehiclePosition_CongestionLevel = 3
	VehiclePosition_SEVERE_CONGESTION        VehiclePosition_CongestionLevel = 4
)

var VehiclePosition_CongestionLevel_name = map[int32]string{
	0: "UNKNOWN_CONGESTION_LEVEL",
	1: "RUNNING_SMOOTHLY",
	2: "STOP_AND_GO",
	3: "CONGESTION",
	4: "SEVERE_CONGESTION",
}
var VehiclePosition_CongestionLevel_value = map[string]int32{
	"UNKNOWN_CONGESTION_LEVEL": 0,
	"RUNNING_SMOOTHLY":         1,
	"STOP_AND_GO":              2,
	"CONGESTION":               3,
	"SEVERE_CONGESTION":        4,
}

func (x VehiclePosition_CongestionLevel) Enum() *VehiclePosition_CongestionLevel {
	p := new(VehiclePosition_CongestionLevel)
	*p = x
	return p
}
func (x VehiclePosition_CongestionLevel) String() string {
	return proto.EnumName(VehiclePosition_CongestionLevel_name, int32(x))
}
func (x *VehiclePosition_CongestionLevel) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(VehiclePosition_CongestionLevel_value, data, "VehiclePosition_CongestionLevel")
	if err != nil {
		return err
	}
	*x = VehiclePosition_CongestionLevel(value)
	return nil
}

type Alert_Cause int32

const (
	Alert_UNKNOWN_CAUSE     Alert_Cause = 1
	Alert_OTHER_CAUSE       Alert_Cause = 2
	Alert_TECHNICAL_PROBLEM Alert_Cause = 3
	Alert_STRIKE            Alert_Cause = 4
	Alert_DEMONSTRATION     Alert_Cause = 5
	Alert_ACCIDENT          Alert_Cause = 6
	Alert_HOLIDAY           Alert_Cause = 7
	Alert_WEATHER           Alert_Cause = 8
	Alert_MAINTENANCE       Alert_Cause = 9
	Alert_CONSTRUCTION      Alert_Cause = 10
	Alert_POLICE_ACTIVITY   Alert_Cause = 11
	Alert_MEDICAL_EMERGENCY Alert_Cause = 12
)

var Alert_Cause_name = map[int32]string{
	1:  "UNKNOWN_CAUSE",
	2:  "OTHER_CAUSE",
	3:  "TECHNICAL_PROBLEM",
	4:  "STRIKE",
	5:  "DEMONSTRATION",
	6:  "ACCIDENT",
	7:  "HOLIDAY",
	8:  "WEATHER",
	9:  "MAINTENANCE",
	10: "CONSTRUCTION",
	11: "POLICE_ACTIVITY",
	12: "MEDICAL_EMERGENCY",
}
var Alert_Cause_value = map[string]int32{
	"UNKNOWN_CAUSE":     1,
	"OTHER_CAUSE":       2,
	"TECHNICAL_PROBLEM": 3,
	"STRIKE":            4,
	"DEMONSTRATION":     5,
	"ACCIDENT":          6,
	"HOLIDAY":           7,
	"WEATHER":           8,
	"MAINTENANCE":       9,
	"CONSTRUCTION":      10,
	"POLICE_ACTIVITY":   11,
	"MEDICAL_EMERGENCY": 12,
}

func (x Alert_Cause) Enum() *Alert_Cause {
	p := new(Alert_Cause)
	*p = x
	return p
}
func (x Alert_Cause) String() string {
	return proto.EnumName(Alert_Cause_name, int32(x))
}
func (x *Alert_Cause) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Alert_Cause_value, data, "Alert_Cause")
	if err != nil {
		return err
	}
	*x = Alert_Cause(value)
	return nil
}

type Alert_Effect int32

const (
	Alert_NO_SERVICE         Alert_Effect = 1
	Alert_REDUCED_SERVICE    Alert_Effect = 2
	Alert_SIGNIFICANT_DELAYS Alert_Effect = 3
	Alert_DETOUR             Alert_Effect = 4
	Alert_ADDITIONAL_SERVICE Alert_Effect = 5
	Alert_MODIFIED_SERVICE   Alert_Effect = 6
	Alert_OTHER_EFFECT       Alert_Effect = 7
	Alert_UNKNOWN_EFFECT     Alert_Effect = 8
	Alert_STOP_MOVED         Alert_Effect = 9
)

var Alert_Effect_name = map[int32]string{
	1: "NO_SERVICE",
	2: "REDUCED_SERVICE",
	3: "SIGNIFICANT_DELAYS",
	4: "DETOUR",
	5: "ADDITIONAL_SERVICE",
	6: "MODIFIED_SERVICE",
	7: "OTHER_EFFECT",
	8: "UNKNOWN_EFFECT",
	9: "STOP_MOVED",
}
var Alert_Effect_value = map[string]int32{
	"NO_SERVICE":         1,
	"REDUCED_SERVICE":    2,
	"SIGNIFICANT_DELAYS": 3,
	"DETOUR":             4,
	"ADDITIONAL_SERVICE": 5,
	"MODIFIED_SERVICE":   6,
	"OTHER_EFFECT":       7,
	"UNKNOWN_EFFECT":     8,
	"STOP_MOVED":         9,
}

func (x Alert_Effect) Enum() *Alert_Effect {
	p := new(Alert_Effect)
	*p = x
	return p
}
func (x Alert_Effect) String() string {
	return proto.EnumName(Alert_Effect_name, int32(x))
}
func (x *Alert_Effect) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Alert_Effect_value, data, "Alert_Effect")
	if err != nil {
		return err
	}
	*x = Alert_Effect(value)
	return nil
}

type TripDescriptor_ScheduleRelationship int32

const (
	TripDescriptor_SCHEDULED   TripDescriptor_ScheduleRelationship = 0
	TripDescriptor_ADDED       TripDescriptor_ScheduleRelationship = 1
	TripDescriptor_UNSCHEDULED TripDescriptor_ScheduleRelationship = 2
	TripDescriptor_CANCELED    TripDescriptor_ScheduleRelationship = 3
)

var TripDescriptor_ScheduleRelationship_name = map[int32]string{
	0: "SCHEDULED",
	1: "ADDED",
	2: "UNSCHEDULED",
	3: "CANCELED",
}
var TripDescriptor_ScheduleRelationship_value = map[string]int32{
	"SCHEDULED":   0,
	"ADDED":       1,
	"UNSCHEDULED": 2,
	"CANCELED":    3,
}

func (x TripDescriptor_ScheduleRelationship) Enum() *TripDescriptor_ScheduleRelationship {
	p := new(TripDescriptor_ScheduleRelationship)
	*p = x
	return p
}
func (x TripDescriptor_ScheduleRelationship) String() string {
	return proto.EnumName(TripDescriptor_ScheduleRelationship_name, int32(x))
}
func (x *TripDescriptor_ScheduleRelationship) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TripDescriptor_ScheduleRelationship_value, data, "TripDescriptor_ScheduleRelationship")
	if err != nil {
		return err
	}
	*x = TripDescriptor_ScheduleRelationship(value)
	return nil
}

type FeedMessage struct {
	Header           *FeedHeader   `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Entity           []*FeedEntity `protobuf:"bytes,2,rep,name=entity" json:"entity,omitempty"`
	XXX_unrecognized []byte        `json:"-"`
}

func (m *FeedMessage) Reset()         { *m = FeedMessage{} }
func (m *FeedMessage) String() string { return proto.CompactTextString(m) }
func (*FeedMessage) ProtoMessage()    {}

func (m *FeedMessage) GetHeader() *FeedHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *FeedMessage) GetEntity() []*FeedEntity {
	if m != nil {
		return m.Entity
	}
	return nil
}

type FeedHeader struct {
	GtfsRealtimeVersion *string                    `protobuf:"bytes,1,req,name=gtfs_realtime_version" json:"gtfs_realtime_version,omitempty"`
	Incrementality      *FeedHeader_Incrementality `protobuf:"varint,2,opt,name=incrementality,enum=transit_realtime.FeedHeader_Incrementality,def=0" json:"incrementality,omitempty"`
	Timestamp           *uint64                    `protobuf:"varint,3,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_extensions      map[int32]proto.Extension  `json:"-"`
	XXX_unrecognized    []byte                     `json:"-"`
}

func (m *FeedHeader) Reset()         { *m = FeedHeader{} }
func (m *FeedHeader) String() string { return proto.CompactTextString(m) }
func (*FeedHeader) ProtoMessage()    {}

var extRange_FeedHeader = []proto.ExtensionRange{
	{1000, 1999},
}

func (*FeedHeader) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_FeedHeader
}
func (m *FeedHeader) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_FeedHeader_Incrementality FeedHeader_Incrementality = FeedHeader_FULL_DATASET

func (m *FeedHeader) GetGtfsRealtimeVersion() string {
	if m != nil && m.GtfsRealtimeVersion != nil {
		return *m.GtfsRealtimeVersion
	}
	return ""
}

func (m *FeedHeader) GetIncrementality() FeedHeader_Incrementality {
	if m != nil && m.Incrementality != nil {
		return *m.Incrementality
	}
	return Default_FeedHeader_Incrementality
}

func (m *FeedHeader) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type FeedEntity struct {
	Id               *string          `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	IsDeleted        *bool            `protobuf:"varint,2,opt,name=is_deleted,def=0" json:"is_deleted,omitempty"`
	TripUpdate       *TripUpdate      `protobuf:"bytes,3,opt,name=trip_update" json:"trip_update,omitempty"`
	Vehicle          *VehiclePosition `protobuf:"bytes,4,opt,name=vehicle" json:"vehicle,omitempty"`
	Alert            *Alert           `protobuf:"bytes,5,opt,name=alert" json:"alert,omitempty"`
	XXX_unrecognized []byte           `json:"-"`
}

func (m *FeedEntity) Reset()         { *m = FeedEntity{} }
func (m *FeedEntity) String() string { return proto.CompactTextString(m) }
func (*FeedEntity) ProtoMessage()    {}

const Default_FeedEntity_IsDeleted bool = false

func (m *FeedEntity) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *FeedEntity) GetIsDeleted() bool {
	if m != nil && m.IsDeleted != nil {
		return *m.IsDeleted
	}
	return Default_FeedEntity_IsDeleted
}

func (m *FeedEntity) GetTripUpdate() *TripUpdate {
	if m != nil {
		return m.TripUpdate
	}
	return nil
}

func (m *FeedEntity) GetVehicle() *VehiclePosition {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

func (m *FeedEntity) GetAlert() *Alert {
	if m != nil {
		return m.Alert
	}
	return nil
}

type TripUpdate struct {
	Trip             *TripDescriptor              `protobuf:"bytes,1,req,name=trip" json:"trip,omitempty"`
	Vehicle          *VehicleDescriptor           `protobuf:"bytes,3,opt,name=vehicle" json:"vehicle,omitempty"`
	StopTimeUpdate   []*TripUpdate_StopTimeUpdate `protobuf:"bytes,2,rep,name=stop_time_update" json:"stop_time_update,omitempty"`
	Timestamp        *uint64                      `protobuf:"varint,4,opt,name=timestamp" json:"timestamp,omitempty"`
	XXX_extensions   map[int32]proto.Extension    `json:"-"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *TripUpdate) Reset()         { *m = TripUpdate{} }
func (m *TripUpdate) String() string { return proto.CompactTextString(m) }
func (*TripUpdate) ProtoMessage()    {}

var extRange_TripUpdate = []proto.ExtensionRange{
	{1000, 1999},
}

func (*TripUpdate) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TripUpdate
}
func (m *TripUpdate) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *TripUpdate) GetTrip() *TripDescriptor {
	if m != nil {
		return m.Trip
	}
	return nil
}

func (m *TripUpdate) GetVehicle() *VehicleDescriptor {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

func (m *TripUpdate) GetStopTimeUpdate() []*TripUpdate_StopTimeUpdate {
	if m != nil {
		return m.StopTimeUpdate
	}
	return nil
}

func (m *TripUpdate) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

type TripUpdate_StopTimeEvent struct {
	Delay            *int32                    `protobuf:"varint,1,opt,name=delay" json:"delay,omitempty"`
	Time             *int64                    `protobuf:"varint,2,opt,name=time" json:"time,omitempty"`
	Uncertainty      *int32                    `protobuf:"varint,3,opt,name=uncertainty" json:"uncertainty,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *TripUpdate_StopTimeEvent) Reset()         { *m = TripUpdate_StopTimeEvent{} }
func (m *TripUpdate_StopTimeEvent) String() string { return proto.CompactTextString(m) }
func (*TripUpdate_StopTimeEvent) ProtoMessage()    {}

var extRange_TripUpdate_StopTimeEvent = []proto.ExtensionRange{
	{1000, 1999},
}

func (*TripUpdate_StopTimeEvent) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TripUpdate_StopTimeEvent
}
func (m *TripUpdate_StopTimeEvent) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *TripUpdate_StopTimeEvent) GetDelay() int32 {
	if m != nil && m.Delay != nil {
		return *m.Delay
	}
	return 0
}

func (m *TripUpdate_StopTimeEvent) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *TripUpdate_StopTimeEvent) GetUncertainty() int32 {
	if m != nil && m.Uncertainty != nil {
		return *m.Uncertainty
	}
	return 0
}

type TripUpdate_StopTimeUpdate struct {
	StopSequence         *uint32                                         `protobuf:"varint,1,opt,name=stop_sequence" json:"stop_sequence,omitempty"`
	StopId               *string                                         `protobuf:"bytes,4,opt,name=stop_id" json:"stop_id,omitempty"`
	Arrival              *TripUpdate_StopTimeEvent                       `protobuf:"bytes,2,opt,name=arrival" json:"arrival,omitempty"`
	Departure            *TripUpdate_StopTimeEvent                       `protobuf:"bytes,3,opt,name=departure" json:"departure,omitempty"`
	ScheduleRelationship *TripUpdate_StopTimeUpdate_ScheduleRelationship `protobuf:"varint,5,opt,name=schedule_relationship,enum=transit_realtime.TripUpdate_StopTimeUpdate_ScheduleRelationship,def=0" json:"schedule_relationship,omitempty"`
	XXX_extensions       map[int32]proto.Extension                       `json:"-"`
	XXX_unrecognized     []byte                                          `json:"-"`
}

func (m *TripUpdate_StopTimeUpdate) Reset()         { *m = TripUpdate_StopTimeUpdate{} }
func (m *TripUpdate_StopTimeUpdate) String() string { return proto.CompactTextString(m) }
func (*TripUpdate_StopTimeUpdate) ProtoMessage()    {}

var extRange_TripUpdate_StopTimeUpdate = []proto.ExtensionRange{
	{1000, 1999},
}

func (*TripUpdate_StopTimeUpdate) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TripUpdate_StopTimeUpdate
}
func (m *TripUpdate_StopTimeUpdate) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_TripUpdate_StopTimeUpdate_ScheduleRelationship TripUpdate_StopTimeUpdate_ScheduleRelationship = TripUpdate_StopTimeUpdate_SCHEDULED

func (m *TripUpdate_StopTimeUpdate) GetStopSequence() uint32 {
	if m != nil && m.StopSequence != nil {
		return *m.StopSequence
	}
	return 0
}

func (m *TripUpdate_StopTimeUpdate) GetStopId() string {
	if m != nil && m.StopId != nil {
		return *m.StopId
	}
	return ""
}

func (m *TripUpdate_StopTimeUpdate) GetArrival() *TripUpdate_StopTimeEvent {
	if m != nil {
		return m.Arrival
	}
	return nil
}

func (m *TripUpdate_StopTimeUpdate) GetDeparture() *TripUpdate_StopTimeEvent {
	if m != nil {
		return m.Departure
	}
	return nil
}

func (m *TripUpdate_StopTimeUpdate) GetScheduleRelationship() TripUpdate_StopTimeUpdate_ScheduleRelationship {
	if m != nil && m.ScheduleRelationship != nil {
		return *m.ScheduleRelationship
	}
	return Default_TripUpdate_StopTimeUpdate_ScheduleRelationship
}

type VehiclePosition struct {
	Trip                *TripDescriptor                    `protobuf:"bytes,1,opt,name=trip" json:"trip,omitempty"`
	Vehicle             *VehicleDescriptor                 `protobuf:"bytes,8,opt,name=vehicle" json:"vehicle,omitempty"`
	Position            *Position                          `protobuf:"bytes,2,opt,name=position" json:"position,omitempty"`
	CurrentStopSequence *uint32                            `protobuf:"varint,3,opt,name=current_stop_sequence" json:"current_stop_sequence,omitempty"`
	StopId              *string                            `protobuf:"bytes,7,opt,name=stop_id" json:"stop_id,omitempty"`
	CurrentStatus       *VehiclePosition_VehicleStopStatus `protobuf:"varint,4,opt,name=current_status,enum=transit_realtime.VehiclePosition_VehicleStopStatus,def=2" json:"current_status,omitempty"`
	Timestamp           *uint64                            `protobuf:"varint,5,opt,name=timestamp" json:"timestamp,omitempty"`
	CongestionLevel     *VehiclePosition_CongestionLevel   `protobuf:"varint,6,opt,name=congestion_level,enum=transit_realtime.VehiclePosition_CongestionLevel" json:"congestion_level,omitempty"`
	XXX_extensions      map[int32]proto.Extension          `json:"-"`
	XXX_unrecognized    []byte                             `json:"-"`
}

func (m *VehiclePosition) Reset()         { *m = VehiclePosition{} }
func (m *VehiclePosition) String() string { return proto.CompactTextString(m) }
func (*VehiclePosition) ProtoMessage()    {}

var extRange_VehiclePosition = []proto.ExtensionRange{
	{1000, 1999},
}

func (*VehiclePosition) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_VehiclePosition
}
func (m *VehiclePosition) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_VehiclePosition_CurrentStatus VehiclePosition_VehicleStopStatus = VehiclePosition_IN_TRANSIT_TO

func (m *VehiclePosition) GetTrip() *TripDescriptor {
	if m != nil {
		return m.Trip
	}
	return nil
}

func (m *VehiclePosition) GetVehicle() *VehicleDescriptor {
	if m != nil {
		return m.Vehicle
	}
	return nil
}

func (m *VehiclePosition) GetPosition() *Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *VehiclePosition) GetCurrentStopSequence() uint32 {
	if m != nil && m.CurrentStopSequence != nil {
		return *m.CurrentStopSequence
	}
	return 0
}

func (m *VehiclePosition) GetStopId() string {
	if m != nil && m.StopId != nil {
		return *m.StopId
	}
	return ""
}

func (m *VehiclePosition) GetCurrentStatus() VehiclePosition_VehicleStopStatus {
	if m != nil && m.CurrentStatus != nil {
		return *m.CurrentStatus
	}
	return Default_VehiclePosition_CurrentStatus
}

func (m *VehiclePosition) GetTimestamp() uint64 {
	if m != nil && m.Timestamp != nil {
		return *m.Timestamp
	}
	return 0
}

func (m *VehiclePosition) GetCongestionLevel() VehiclePosition_CongestionLevel {
	if m != nil && m.CongestionLevel != nil {
		return *m.CongestionLevel
	}
	return VehiclePosition_UNKNOWN_CONGESTION_LEVEL
}

type Alert struct {
	ActivePeriod     []*TimeRange              `protobuf:"bytes,1,rep,name=active_period" json:"active_period,omitempty"`
	InformedEntity   []*EntitySelector         `protobuf:"bytes,5,rep,name=informed_entity" json:"informed_entity,omitempty"`
	Cause            *Alert_Cause              `protobuf:"varint,6,opt,name=cause,enum=transit_realtime.Alert_Cause,def=1" json:"cause,omitempty"`
	Effect           *Alert_Effect             `protobuf:"varint,7,opt,name=effect,enum=transit_realtime.Alert_Effect,def=8" json:"effect,omitempty"`
	Url              *TranslatedString         `protobuf:"bytes,8,opt,name=url" json:"url,omitempty"`
	HeaderText       *TranslatedString         `protobuf:"bytes,10,opt,name=header_text" json:"header_text,omitempty"`
	DescriptionText  *TranslatedString         `protobuf:"bytes,11,opt,name=description_text" json:"description_text,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Alert) Reset()         { *m = Alert{} }
func (m *Alert) String() string { return proto.CompactTextString(m) }
func (*Alert) ProtoMessage()    {}

var extRange_Alert = []proto.ExtensionRange{
	{1000, 1999},
}

func (*Alert) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Alert
}
func (m *Alert) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

const Default_Alert_Cause Alert_Cause = Alert_UNKNOWN_CAUSE
const Default_Alert_Effect Alert_Effect = Alert_UNKNOWN_EFFECT

func (m *Alert) GetActivePeriod() []*TimeRange {
	if m != nil {
		return m.ActivePeriod
	}
	return nil
}

func (m *Alert) GetInformedEntity() []*EntitySelector {
	if m != nil {
		return m.InformedEntity
	}
	return nil
}

func (m *Alert) GetCause() Alert_Cause {
	if m != nil && m.Cause != nil {
		return *m.Cause
	}
	return Default_Alert_Cause
}

func (m *Alert) GetEffect() Alert_Effect {
	if m != nil && m.Effect != nil {
		return *m.Effect
	}
	return Default_Alert_Effect
}

func (m *Alert) GetUrl() *TranslatedString {
	if m != nil {
		return m.Url
	}
	return nil
}

func (m *Alert) GetHeaderText() *TranslatedString {
	if m != nil {
		return m.HeaderText
	}
	return nil
}

func (m *Alert) GetDescriptionText() *TranslatedString {
	if m != nil {
		return m.DescriptionText
	}
	return nil
}

type TimeRange struct {
	Start            *uint64 `protobuf:"varint,1,opt,name=start" json:"start,omitempty"`
	End              *uint64 `protobuf:"varint,2,opt,name=end" json:"end,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TimeRange) Reset()         { *m = TimeRange{} }
func (m *TimeRange) String() string { return proto.CompactTextString(m) }
func (*TimeRange) ProtoMessage()    {}

func (m *TimeRange) GetStart() uint64 {
	if m != nil && m.Start != nil {
		return *m.Start
	}
	return 0
}

func (m *TimeRange) GetEnd() uint64 {
	if m != nil && m.End != nil {
		return *m.End
	}
	return 0
}

type Position struct {
	Latitude         *float32                  `protobuf:"fixed32,1,req,name=latitude" json:"latitude,omitempty"`
	Longitude        *float32                  `protobuf:"fixed32,2,req,name=longitude" json:"longitude,omitempty"`
	Bearing          *float32                  `protobuf:"fixed32,3,opt,name=bearing" json:"bearing,omitempty"`
	Odometer         *float64                  `protobuf:"fixed64,4,opt,name=odometer" json:"odometer,omitempty"`
	Speed            *float32                  `protobuf:"fixed32,5,opt,name=speed" json:"speed,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *Position) Reset()         { *m = Position{} }
func (m *Position) String() string { return proto.CompactTextString(m) }
func (*Position) ProtoMessage()    {}

var extRange_Position = []proto.ExtensionRange{
	{1000, 1999},
}

func (*Position) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_Position
}
func (m *Position) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *Position) GetLatitude() float32 {
	if m != nil && m.Latitude != nil {
		return *m.Latitude
	}
	return 0
}

func (m *Position) GetLongitude() float32 {
	if m != nil && m.Longitude != nil {
		return *m.Longitude
	}
	return 0
}

func (m *Position) GetBearing() float32 {
	if m != nil && m.Bearing != nil {
		return *m.Bearing
	}
	return 0
}

func (m *Position) GetOdometer() float64 {
	if m != nil && m.Odometer != nil {
		return *m.Odometer
	}
	return 0
}

func (m *Position) GetSpeed() float32 {
	if m != nil && m.Speed != nil {
		return *m.Speed
	}
	return 0
}

type TripDescriptor struct {
	TripId               *string                              `protobuf:"bytes,1,opt,name=trip_id" json:"trip_id,omitempty"`
	RouteId              *string                              `protobuf:"bytes,5,opt,name=route_id" json:"route_id,omitempty"`
	StartTime            *string                              `protobuf:"bytes,2,opt,name=start_time" json:"start_time,omitempty"`
	StartDate            *string                              `protobuf:"bytes,3,opt,name=start_date" json:"start_date,omitempty"`
	ScheduleRelationship *TripDescriptor_ScheduleRelationship `protobuf:"varint,4,opt,name=schedule_relationship,enum=transit_realtime.TripDescriptor_ScheduleRelationship" json:"schedule_relationship,omitempty"`
	XXX_extensions       map[int32]proto.Extension            `json:"-"`
	XXX_unrecognized     []byte                               `json:"-"`
}

func (m *TripDescriptor) Reset()         { *m = TripDescriptor{} }
func (m *TripDescriptor) String() string { return proto.CompactTextString(m) }
func (*TripDescriptor) ProtoMessage()    {}

var extRange_TripDescriptor = []proto.ExtensionRange{
	{1000, 1999},
}

func (*TripDescriptor) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_TripDescriptor
}
func (m *TripDescriptor) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *TripDescriptor) GetTripId() string {
	if m != nil && m.TripId != nil {
		return *m.TripId
	}
	return ""
}

func (m *TripDescriptor) GetRouteId() string {
	if m != nil && m.RouteId != nil {
		return *m.RouteId
	}
	return ""
}

func (m *TripDescriptor) GetStartTime() string {
	if m != nil && m.StartTime != nil {
		return *m.StartTime
	}
	return ""
}

func (m *TripDescriptor) GetStartDate() string {
	if m != nil && m.StartDate != nil {
		return *m.StartDate
	}
	return ""
}

func (m *TripDescriptor) GetScheduleRelationship() TripDescriptor_ScheduleRelationship {
	if m != nil && m.ScheduleRelationship != nil {
		return *m.ScheduleRelationship
	}
	return TripDescriptor_SCHEDULED
}

type VehicleDescriptor struct {
	Id               *string                   `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Label            *string                   `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	LicensePlate     *string                   `protobuf:"bytes,3,opt,name=license_plate" json:"license_plate,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *VehicleDescriptor) Reset()         { *m = VehicleDescriptor{} }
func (m *VehicleDescriptor) String() string { return proto.CompactTextString(m) }
func (*VehicleDescriptor) ProtoMessage()    {}

var extRange_VehicleDescriptor = []proto.ExtensionRange{
	{1000, 1999},
}

func (*VehicleDescriptor) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_VehicleDescriptor
}
func (m *VehicleDescriptor) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *VehicleDescriptor) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *VehicleDescriptor) GetLabel() string {
	if m != nil && m.Label != nil {
		return *m.Label
	}
	return ""
}

func (m *VehicleDescriptor) GetLicensePlate() string {
	if m != nil && m.LicensePlate != nil {
		return *m.LicensePlate
	}
	return ""
}

type EntitySelector struct {
	AgencyId         *string                   `protobuf:"bytes,1,opt,name=agency_id" json:"agency_id,omitempty"`
	RouteId          *string                   `protobuf:"bytes,2,opt,name=route_id" json:"route_id,omitempty"`
	RouteType        *int32                    `protobuf:"varint,3,opt,name=route_type" json:"route_type,omitempty"`
	Trip             *TripDescriptor           `protobuf:"bytes,4,opt,name=trip" json:"trip,omitempty"`
	StopId           *string                   `protobuf:"bytes,5,opt,name=stop_id" json:"stop_id,omitempty"`
	XXX_extensions   map[int32]proto.Extension `json:"-"`
	XXX_unrecognized []byte                    `json:"-"`
}

func (m *EntitySelector) Reset()         { *m = EntitySelector{} }
func (m *EntitySelector) String() string { return proto.CompactTextString(m) }
func (*EntitySelector) ProtoMessage()    {}

var extRange_EntitySelector = []proto.ExtensionRange{
	{1000, 1999},
}

func (*EntitySelector) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_EntitySelector
}
func (m *EntitySelector) ExtensionMap() map[int32]proto.Extension {
	if m.XXX_extensions == nil {
		m.XXX_extensions = make(map[int32]proto.Extension)
	}
	return m.XXX_extensions
}

func (m *EntitySelector) GetAgencyId() string {
	if m != nil && m.AgencyId != nil {
		return *m.AgencyId
	}
	return ""
}

func (m *EntitySelector) GetRouteId() string {
	if m != nil && m.RouteId != nil {
		return *m.RouteId
	}
	return ""
}

func (m *EntitySelector) GetRouteType() int32 {
	if m != nil && m.RouteType != nil {
		return *m.RouteType
	}
	return 0
}

func (m *EntitySelector) GetTrip() *TripDescriptor {
	if m != nil {
		return m.Trip
	}
	return nil
}

func (m *EntitySelector) GetStopId() string {
	if m != nil && m.StopId != nil {
		return *m.StopId
	}
	return ""
}

type TranslatedString struct {
	Translation      []*TranslatedString_Translation `protobuf:"bytes,1,rep,name=translation" json:"translation,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *TranslatedString) Reset()         { *m = TranslatedString{} }
func (m *TranslatedString) String() string { return proto.CompactTextString(m) }
func (*TranslatedString) ProtoMessage()    {}

func (m *TranslatedString) GetTranslation() []*TranslatedString_Translation {
	if m != nil {
		return m.Translation
	}
	return nil
}

type TranslatedString_Translation struct {
	Text             *string `protobuf:"bytes,1,req,name=text" json:"text,omitempty"`
	Language         *string `protobuf:"bytes,2,opt,name=language" json:"language,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *TranslatedString_Translation) Reset()         { *m = TranslatedString_Translation{} }
func (m *TranslatedString_Translation) String() string { return proto.CompactTextString(m) }
func (*TranslatedString_Translation) ProtoMessage()    {}

func (m *TranslatedString_Translation) GetText() string {
	if m != nil && m.Text != nil {
		return *m.Text
	}
	return ""
}

func (m *TranslatedString_Translation) GetLanguage() string {
	if m != nil && m.Language != nil {
		return *m.Language
	}
	return ""
}

func init() {
	proto.RegisterEnum("transit_realtime.FeedHeader_Incrementality", FeedHeader_Incrementality_name, FeedHeader_Incrementality_value)
	proto.RegisterEnum("transit_realtime.TripUpdate_StopTimeUpdate_ScheduleRelationship", TripUpdate_StopTimeUpdate_ScheduleRelationship_name, TripUpdate_StopTimeUpdate_ScheduleRelationship_value)
	proto.RegisterEnum("transit_realtime.VehiclePosition_VehicleStopStatus", VehiclePosition_VehicleStopStatus_name, VehiclePosition_VehicleStopStatus_value)
	proto.RegisterEnum("transit_realtime.VehiclePosition_CongestionLevel", VehiclePosition_CongestionLevel_name, VehiclePosition_CongestionLevel_value)
	proto.RegisterEnum("transit_realtime.Alert_Cause", Alert_Cause_name, Alert_Cause_value)
	proto.RegisterEnum("transit_realtime.Alert_Effect", Alert_Effect_name, Alert_Effect_value)
	proto.RegisterEnum("transit_realtime.TripDescriptor_ScheduleRelationship", TripDescriptor_ScheduleRelationship_name, TripDescriptor_ScheduleRelationship_value)
}
