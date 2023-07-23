// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: station/station.proto

package station

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GetManyStationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Query string `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	Limit int64  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Page  int64  `protobuf:"varint,3,opt,name=page,proto3" json:"page,omitempty"`
	Token string `protobuf:"bytes,4,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetManyStationsRequest) Reset() {
	*x = GetManyStationsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyStationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyStationsRequest) ProtoMessage() {}

func (x *GetManyStationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyStationsRequest.ProtoReflect.Descriptor instead.
func (*GetManyStationsRequest) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{0}
}

func (x *GetManyStationsRequest) GetQuery() string {
	if x != nil {
		return x.Query
	}
	return ""
}

func (x *GetManyStationsRequest) GetLimit() int64 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetManyStationsRequest) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *GetManyStationsRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetManyStationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Stations *PaginatedStation `protobuf:"bytes,1,opt,name=stations,proto3" json:"stations,omitempty"`
}

func (x *GetManyStationsResponse) Reset() {
	*x = GetManyStationsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetManyStationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetManyStationsResponse) ProtoMessage() {}

func (x *GetManyStationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetManyStationsResponse.ProtoReflect.Descriptor instead.
func (*GetManyStationsResponse) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{1}
}

func (x *GetManyStationsResponse) GetStations() *PaginatedStation {
	if x != nil {
		return x.Stations
	}
	return nil
}

type GetOneStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GetOneStationRequest) Reset() {
	*x = GetOneStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneStationRequest) ProtoMessage() {}

func (x *GetOneStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneStationRequest.ProtoReflect.Descriptor instead.
func (*GetOneStationRequest) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{2}
}

func (x *GetOneStationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *GetOneStationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type GetOneStationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Station *Station `protobuf:"bytes,1,opt,name=station,proto3" json:"station,omitempty"`
}

func (x *GetOneStationResponse) Reset() {
	*x = GetOneStationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOneStationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOneStationResponse) ProtoMessage() {}

func (x *GetOneStationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOneStationResponse.ProtoReflect.Descriptor instead.
func (*GetOneStationResponse) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{3}
}

func (x *GetOneStationResponse) GetStation() *Station {
	if x != nil {
		return x.Station
	}
	return nil
}

type SetFavoriteOneStationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Token string `protobuf:"bytes,2,opt,name=token,proto3" json:"token,omitempty"`
	Value bool   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (x *SetFavoriteOneStationRequest) Reset() {
	*x = SetFavoriteOneStationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFavoriteOneStationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFavoriteOneStationRequest) ProtoMessage() {}

func (x *SetFavoriteOneStationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFavoriteOneStationRequest.ProtoReflect.Descriptor instead.
func (*SetFavoriteOneStationRequest) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{4}
}

func (x *SetFavoriteOneStationRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *SetFavoriteOneStationRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *SetFavoriteOneStationRequest) GetValue() bool {
	if x != nil {
		return x.Value
	}
	return false
}

type SetFavoriteOneStationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetFavoriteOneStationResponse) Reset() {
	*x = SetFavoriteOneStationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetFavoriteOneStationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetFavoriteOneStationResponse) ProtoMessage() {}

func (x *SetFavoriteOneStationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetFavoriteOneStationResponse.ProtoReflect.Descriptor instead.
func (*SetFavoriteOneStationResponse) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{5}
}

type PaginatedStation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      []*Station `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty"`
	Count     int64      `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
	Total     int64      `protobuf:"varint,3,opt,name=total,proto3" json:"total,omitempty"`
	Page      int64      `protobuf:"varint,4,opt,name=page,proto3" json:"page,omitempty"`
	PageCount int64      `protobuf:"varint,5,opt,name=page_count,json=pageCount,proto3" json:"page_count,omitempty"`
}

func (x *PaginatedStation) Reset() {
	*x = PaginatedStation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PaginatedStation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PaginatedStation) ProtoMessage() {}

func (x *PaginatedStation) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PaginatedStation.ProtoReflect.Descriptor instead.
func (*PaginatedStation) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{6}
}

func (x *PaginatedStation) GetData() []*Station {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *PaginatedStation) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *PaginatedStation) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

func (x *PaginatedStation) GetPage() int64 {
	if x != nil {
		return x.Page
	}
	return 0
}

func (x *PaginatedStation) GetPageCount() int64 {
	if x != nil {
		return x.PageCount
	}
	return 0
}

type Station struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Commune     string    `protobuf:"bytes,2,opt,name=commune,proto3" json:"commune,omitempty"`
	YWgs84      float64   `protobuf:"fixed64,3,opt,name=y_wgs84,json=yWgs84,proto3" json:"y_wgs84,omitempty"`
	XWgs84      float64   `protobuf:"fixed64,4,opt,name=x_wgs84,json=xWgs84,proto3" json:"x_wgs84,omitempty"`
	Libelle     string    `protobuf:"bytes,5,opt,name=libelle,proto3" json:"libelle,omitempty"`
	Idgaia      string    `protobuf:"bytes,6,opt,name=idgaia,proto3" json:"idgaia,omitempty"`
	Voyageurs   string    `protobuf:"bytes,7,opt,name=voyageurs,proto3" json:"voyageurs,omitempty"`
	GeoPoint_2D []float64 `protobuf:"fixed64,8,rep,packed,name=geo_point_2d,json=geoPoint2d,proto3" json:"geo_point_2d,omitempty"`
	CodeLigne   string    `protobuf:"bytes,9,opt,name=code_ligne,json=codeLigne,proto3" json:"code_ligne,omitempty"`
	XL93        float64   `protobuf:"fixed64,10,opt,name=x_l93,json=xL93,proto3" json:"x_l93,omitempty"`
	CGeo        []float64 `protobuf:"fixed64,11,rep,packed,name=c_geo,json=cGeo,proto3" json:"c_geo,omitempty"`
	RgTroncon   int64     `protobuf:"varint,12,opt,name=rg_troncon,json=rgTroncon,proto3" json:"rg_troncon,omitempty"`
	GeoShape    *Geometry `protobuf:"bytes,13,opt,name=geo_shape,json=geoShape,proto3" json:"geo_shape,omitempty"`
	Pk          string    `protobuf:"bytes,14,opt,name=pk,proto3" json:"pk,omitempty"`
	Idreseau    int64     `protobuf:"varint,15,opt,name=idreseau,proto3" json:"idreseau,omitempty"`
	Departemen  string    `protobuf:"bytes,16,opt,name=departemen,proto3" json:"departemen,omitempty"`
	YL93        float64   `protobuf:"fixed64,17,opt,name=y_l93,json=yL93,proto3" json:"y_l93,omitempty"`
	Fret        string    `protobuf:"bytes,18,opt,name=fret,proto3" json:"fret,omitempty"`
	IsFavorite  bool      `protobuf:"varint,19,opt,name=is_favorite,json=isFavorite,proto3" json:"is_favorite,omitempty"`
}

func (x *Station) Reset() {
	*x = Station{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Station) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Station) ProtoMessage() {}

func (x *Station) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Station.ProtoReflect.Descriptor instead.
func (*Station) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{7}
}

func (x *Station) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Station) GetCommune() string {
	if x != nil {
		return x.Commune
	}
	return ""
}

func (x *Station) GetYWgs84() float64 {
	if x != nil {
		return x.YWgs84
	}
	return 0
}

func (x *Station) GetXWgs84() float64 {
	if x != nil {
		return x.XWgs84
	}
	return 0
}

func (x *Station) GetLibelle() string {
	if x != nil {
		return x.Libelle
	}
	return ""
}

func (x *Station) GetIdgaia() string {
	if x != nil {
		return x.Idgaia
	}
	return ""
}

func (x *Station) GetVoyageurs() string {
	if x != nil {
		return x.Voyageurs
	}
	return ""
}

func (x *Station) GetGeoPoint_2D() []float64 {
	if x != nil {
		return x.GeoPoint_2D
	}
	return nil
}

func (x *Station) GetCodeLigne() string {
	if x != nil {
		return x.CodeLigne
	}
	return ""
}

func (x *Station) GetXL93() float64 {
	if x != nil {
		return x.XL93
	}
	return 0
}

func (x *Station) GetCGeo() []float64 {
	if x != nil {
		return x.CGeo
	}
	return nil
}

func (x *Station) GetRgTroncon() int64 {
	if x != nil {
		return x.RgTroncon
	}
	return 0
}

func (x *Station) GetGeoShape() *Geometry {
	if x != nil {
		return x.GeoShape
	}
	return nil
}

func (x *Station) GetPk() string {
	if x != nil {
		return x.Pk
	}
	return ""
}

func (x *Station) GetIdreseau() int64 {
	if x != nil {
		return x.Idreseau
	}
	return 0
}

func (x *Station) GetDepartemen() string {
	if x != nil {
		return x.Departemen
	}
	return ""
}

func (x *Station) GetYL93() float64 {
	if x != nil {
		return x.YL93
	}
	return 0
}

func (x *Station) GetFret() string {
	if x != nil {
		return x.Fret
	}
	return ""
}

func (x *Station) GetIsFavorite() bool {
	if x != nil {
		return x.IsFavorite
	}
	return false
}

type Geometry struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        string    `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Coordinates []float64 `protobuf:"fixed64,2,rep,packed,name=coordinates,proto3" json:"coordinates,omitempty"`
}

func (x *Geometry) Reset() {
	*x = Geometry{}
	if protoimpl.UnsafeEnabled {
		mi := &file_station_station_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Geometry) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Geometry) ProtoMessage() {}

func (x *Geometry) ProtoReflect() protoreflect.Message {
	mi := &file_station_station_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Geometry.ProtoReflect.Descriptor instead.
func (*Geometry) Descriptor() ([]byte, []int) {
	return file_station_station_proto_rawDescGZIP(), []int{8}
}

func (x *Geometry) GetType() string {
	if x != nil {
		return x.Type
	}
	return ""
}

func (x *Geometry) GetCoordinates() []float64 {
	if x != nil {
		return x.Coordinates
	}
	return nil
}

var File_station_station_proto protoreflect.FileDescriptor

var file_station_station_proto_rawDesc = []byte{
	0x0a, 0x15, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x76, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x12, 0x18, 0x0a, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x42,
	0x02, 0x30, 0x02, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74, 0x12, 0x16, 0x0a, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02, 0x52, 0x04, 0x70, 0x61,
	0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x50, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x4d,
	0x61, 0x6e, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e,
	0x50, 0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x52, 0x08, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x3c, 0x0a, 0x14, 0x47, 0x65,
	0x74, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x43, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f,
	0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x2a, 0x0a, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x07, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x5a, 0x0a,
	0x1c, 0x53, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x1f, 0x0a, 0x1d, 0x53, 0x65, 0x74,
	0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0xa7, 0x01, 0x0a, 0x10, 0x50,
	0x61, 0x67, 0x69, 0x6e, 0x61, 0x74, 0x65, 0x64, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x24, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x18, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12,
	0x18, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02,
	0x30, 0x02, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x12, 0x16, 0x0a, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02, 0x52, 0x04, 0x70, 0x61, 0x67,
	0x65, 0x12, 0x21, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02, 0x52, 0x09, 0x70, 0x61, 0x67, 0x65, 0x43,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x8d, 0x04, 0x0a, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x6e, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x79, 0x5f,
	0x77, 0x67, 0x73, 0x38, 0x34, 0x18, 0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x79, 0x57, 0x67,
	0x73, 0x38, 0x34, 0x12, 0x17, 0x0a, 0x07, 0x78, 0x5f, 0x77, 0x67, 0x73, 0x38, 0x34, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x06, 0x78, 0x57, 0x67, 0x73, 0x38, 0x34, 0x12, 0x18, 0x0a, 0x07,
	0x6c, 0x69, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6c,
	0x69, 0x62, 0x65, 0x6c, 0x6c, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x64, 0x67, 0x61, 0x69, 0x61,
	0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x64, 0x67, 0x61, 0x69, 0x61, 0x12, 0x1c,
	0x0a, 0x09, 0x76, 0x6f, 0x79, 0x61, 0x67, 0x65, 0x75, 0x72, 0x73, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x76, 0x6f, 0x79, 0x61, 0x67, 0x65, 0x75, 0x72, 0x73, 0x12, 0x20, 0x0a, 0x0c,
	0x67, 0x65, 0x6f, 0x5f, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x5f, 0x32, 0x64, 0x18, 0x08, 0x20, 0x03,
	0x28, 0x01, 0x52, 0x0a, 0x67, 0x65, 0x6f, 0x50, 0x6f, 0x69, 0x6e, 0x74, 0x32, 0x64, 0x12, 0x1d,
	0x0a, 0x0a, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6c, 0x69, 0x67, 0x6e, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x63, 0x6f, 0x64, 0x65, 0x4c, 0x69, 0x67, 0x6e, 0x65, 0x12, 0x13, 0x0a,
	0x05, 0x78, 0x5f, 0x6c, 0x39, 0x33, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x78, 0x4c,
	0x39, 0x33, 0x12, 0x13, 0x0a, 0x05, 0x63, 0x5f, 0x67, 0x65, 0x6f, 0x18, 0x0b, 0x20, 0x03, 0x28,
	0x01, 0x52, 0x04, 0x63, 0x47, 0x65, 0x6f, 0x12, 0x21, 0x0a, 0x0a, 0x72, 0x67, 0x5f, 0x74, 0x72,
	0x6f, 0x6e, 0x63, 0x6f, 0x6e, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02, 0x52,
	0x09, 0x72, 0x67, 0x54, 0x72, 0x6f, 0x6e, 0x63, 0x6f, 0x6e, 0x12, 0x2e, 0x0a, 0x09, 0x67, 0x65,
	0x6f, 0x5f, 0x73, 0x68, 0x61, 0x70, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x52, 0x08, 0x67, 0x65, 0x6f, 0x53, 0x68, 0x61, 0x70, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x70, 0x6b,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x70, 0x6b, 0x12, 0x1e, 0x0a, 0x08, 0x69, 0x64,
	0x72, 0x65, 0x73, 0x65, 0x61, 0x75, 0x18, 0x0f, 0x20, 0x01, 0x28, 0x03, 0x42, 0x02, 0x30, 0x02,
	0x52, 0x08, 0x69, 0x64, 0x72, 0x65, 0x73, 0x65, 0x61, 0x75, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65,
	0x70, 0x61, 0x72, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a,
	0x64, 0x65, 0x70, 0x61, 0x72, 0x74, 0x65, 0x6d, 0x65, 0x6e, 0x12, 0x13, 0x0a, 0x05, 0x79, 0x5f,
	0x6c, 0x39, 0x33, 0x18, 0x11, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04, 0x79, 0x4c, 0x39, 0x33, 0x12,
	0x12, 0x0a, 0x04, 0x66, 0x72, 0x65, 0x74, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66,
	0x72, 0x65, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x69, 0x73, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x18, 0x13, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x69, 0x73, 0x46, 0x61, 0x76, 0x6f,
	0x72, 0x69, 0x74, 0x65, 0x22, 0x40, 0x0a, 0x08, 0x47, 0x65, 0x6f, 0x6d, 0x65, 0x74, 0x72, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x74, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x01, 0x52, 0x0b, 0x63, 0x6f, 0x6f, 0x72, 0x64,
	0x69, 0x6e, 0x61, 0x74, 0x65, 0x73, 0x32, 0xa0, 0x02, 0x0a, 0x0a, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x41, 0x50, 0x49, 0x12, 0x56, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1f, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x73, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4d, 0x61, 0x6e, 0x79, 0x53, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x50, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1d,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e,
	0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x6e, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x68, 0x0a, 0x15, 0x53, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x25, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e,
	0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x26, 0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x53, 0x65, 0x74, 0x46, 0x61, 0x76,
	0x6f, 0x72, 0x69, 0x74, 0x65, 0x4f, 0x6e, 0x65, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x72, 0x0a, 0x0b, 0x63, 0x6f, 0x6d,
	0x2e, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x0c, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x48, 0x03, 0x50, 0x01, 0x5a, 0x14, 0x7a, 0x7a, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0xf8, 0x01, 0x00, 0xa2, 0x02, 0x03, 0x53, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x53, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0xca, 0x02, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0xe2, 0x02, 0x13,
	0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x07, 0x53, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_station_station_proto_rawDescOnce sync.Once
	file_station_station_proto_rawDescData = file_station_station_proto_rawDesc
)

func file_station_station_proto_rawDescGZIP() []byte {
	file_station_station_proto_rawDescOnce.Do(func() {
		file_station_station_proto_rawDescData = protoimpl.X.CompressGZIP(file_station_station_proto_rawDescData)
	})
	return file_station_station_proto_rawDescData
}

var file_station_station_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_station_station_proto_goTypes = []interface{}{
	(*GetManyStationsRequest)(nil),        // 0: station.GetManyStationsRequest
	(*GetManyStationsResponse)(nil),       // 1: station.GetManyStationsResponse
	(*GetOneStationRequest)(nil),          // 2: station.GetOneStationRequest
	(*GetOneStationResponse)(nil),         // 3: station.GetOneStationResponse
	(*SetFavoriteOneStationRequest)(nil),  // 4: station.SetFavoriteOneStationRequest
	(*SetFavoriteOneStationResponse)(nil), // 5: station.SetFavoriteOneStationResponse
	(*PaginatedStation)(nil),              // 6: station.PaginatedStation
	(*Station)(nil),                       // 7: station.Station
	(*Geometry)(nil),                      // 8: station.Geometry
}
var file_station_station_proto_depIdxs = []int32{
	6, // 0: station.GetManyStationsResponse.stations:type_name -> station.PaginatedStation
	7, // 1: station.GetOneStationResponse.station:type_name -> station.Station
	7, // 2: station.PaginatedStation.data:type_name -> station.Station
	8, // 3: station.Station.geo_shape:type_name -> station.Geometry
	0, // 4: station.StationAPI.GetManyStations:input_type -> station.GetManyStationsRequest
	2, // 5: station.StationAPI.GetOneStation:input_type -> station.GetOneStationRequest
	4, // 6: station.StationAPI.SetFavoriteOneStation:input_type -> station.SetFavoriteOneStationRequest
	1, // 7: station.StationAPI.GetManyStations:output_type -> station.GetManyStationsResponse
	3, // 8: station.StationAPI.GetOneStation:output_type -> station.GetOneStationResponse
	5, // 9: station.StationAPI.SetFavoriteOneStation:output_type -> station.SetFavoriteOneStationResponse
	7, // [7:10] is the sub-list for method output_type
	4, // [4:7] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_station_station_proto_init() }
func file_station_station_proto_init() {
	if File_station_station_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_station_station_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyStationsRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetManyStationsResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneStationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetOneStationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFavoriteOneStationRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetFavoriteOneStationResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PaginatedStation); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Station); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_station_station_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Geometry); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_station_station_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_station_station_proto_goTypes,
		DependencyIndexes: file_station_station_proto_depIdxs,
		MessageInfos:      file_station_station_proto_msgTypes,
	}.Build()
	File_station_station_proto = out.File
	file_station_station_proto_rawDesc = nil
	file_station_station_proto_goTypes = nil
	file_station_station_proto_depIdxs = nil
}
