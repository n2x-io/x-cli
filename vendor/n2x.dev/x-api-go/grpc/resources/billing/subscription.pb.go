// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0--rc1
// source: n2x/protobuf/resources/v1/billing/subscription.proto

package billing

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

type SubscriptionStatus int32

const (
	SubscriptionStatus_TRIALING           SubscriptionStatus = 0
	SubscriptionStatus_ACTIVE             SubscriptionStatus = 1
	SubscriptionStatus_INCOMPLETE         SubscriptionStatus = 2
	SubscriptionStatus_INCOMPLETE_EXPIRED SubscriptionStatus = 3
	SubscriptionStatus_PAST_DUE           SubscriptionStatus = 4
	SubscriptionStatus_CANCELED           SubscriptionStatus = 5
	SubscriptionStatus_UNPAID             SubscriptionStatus = 6
	SubscriptionStatus_UNKNOWN            SubscriptionStatus = 100
	SubscriptionStatus_NONE               SubscriptionStatus = 101
)

// Enum value maps for SubscriptionStatus.
var (
	SubscriptionStatus_name = map[int32]string{
		0:   "TRIALING",
		1:   "ACTIVE",
		2:   "INCOMPLETE",
		3:   "INCOMPLETE_EXPIRED",
		4:   "PAST_DUE",
		5:   "CANCELED",
		6:   "UNPAID",
		100: "UNKNOWN",
		101: "NONE",
	}
	SubscriptionStatus_value = map[string]int32{
		"TRIALING":           0,
		"ACTIVE":             1,
		"INCOMPLETE":         2,
		"INCOMPLETE_EXPIRED": 3,
		"PAST_DUE":           4,
		"CANCELED":           5,
		"UNPAID":             6,
		"UNKNOWN":            100,
		"NONE":               101,
	}
)

func (x SubscriptionStatus) Enum() *SubscriptionStatus {
	p := new(SubscriptionStatus)
	*p = x
	return p
}

func (x SubscriptionStatus) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SubscriptionStatus) Descriptor() protoreflect.EnumDescriptor {
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_enumTypes[0].Descriptor()
}

func (SubscriptionStatus) Type() protoreflect.EnumType {
	return &file_n2x_protobuf_resources_v1_billing_subscription_proto_enumTypes[0]
}

func (x SubscriptionStatus) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SubscriptionStatus.Descriptor instead.
func (SubscriptionStatus) EnumDescriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescGZIP(), []int{0}
}

type Subscription struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccountID string `protobuf:"bytes,1,opt,name=accountID,proto3" json:"accountID,omitempty"` // string subscriptionID = 2; // serviceID-pricingPlanID
	ServiceID string `protobuf:"bytes,3,opt,name=serviceID,proto3" json:"serviceID,omitempty"`
	// string pricingPlanID = 4;
	ServiceProviderID   string `protobuf:"bytes,5,opt,name=serviceProviderID,proto3" json:"serviceProviderID,omitempty"`
	StripeCustomerID    string `protobuf:"bytes,11,opt,name=stripeCustomerID,proto3" json:"stripeCustomerID,omitempty"`
	StripeCustomerEmail string `protobuf:"bytes,12,opt,name=stripeCustomerEmail,proto3" json:"stripeCustomerEmail,omitempty"`
	// string stripeProductID = 12;
	// string stripePriceID = 13;
	StripeSubscriptionID                   string              `protobuf:"bytes,14,opt,name=stripeSubscriptionID,proto3" json:"stripeSubscriptionID,omitempty"`
	DefaultStripePaymentMethodID           string              `protobuf:"bytes,21,opt,name=defaultStripePaymentMethodID,proto3" json:"defaultStripePaymentMethodID,omitempty"`
	AutomaticTax                           bool                `protobuf:"varint,22,opt,name=automaticTax,proto3" json:"automaticTax,omitempty"`
	Discount                               *Discount           `protobuf:"bytes,25,opt,name=discount,proto3" json:"discount,omitempty"`
	CreationDate                           int64               `protobuf:"varint,31,opt,name=creationDate,proto3" json:"creationDate,omitempty"`
	LastModified                           int64               `protobuf:"varint,32,opt,name=lastModified,proto3" json:"lastModified,omitempty"`
	StartDate                              int64               `protobuf:"varint,41,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate                                int64               `protobuf:"varint,42,opt,name=endDate,proto3" json:"endDate,omitempty"`
	TrialStartDate                         int64               `protobuf:"varint,51,opt,name=trialStartDate,proto3" json:"trialStartDate,omitempty"`
	TrialEndDate                           int64               `protobuf:"varint,52,opt,name=trialEndDate,proto3" json:"trialEndDate,omitempty"`
	CancelAtPeriodEnd                      bool                `protobuf:"varint,61,opt,name=cancelAtPeriodEnd,proto3" json:"cancelAtPeriodEnd,omitempty"`
	CancelationDate                        int64               `protobuf:"varint,62,opt,name=cancelationDate,proto3" json:"cancelationDate,omitempty"`
	CurrentPeriodStart                     int64               `protobuf:"varint,71,opt,name=currentPeriodStart,proto3" json:"currentPeriodStart,omitempty"`
	CurrentPeriodEnd                       int64               `protobuf:"varint,72,opt,name=currentPeriodEnd,proto3" json:"currentPeriodEnd,omitempty"`
	LatestStripeInvoiceID                  string              `protobuf:"bytes,91,opt,name=latestStripeInvoiceID,proto3" json:"latestStripeInvoiceID,omitempty"`
	LatestStripeHostedInvoiceURL           string              `protobuf:"bytes,92,opt,name=latestStripeHostedInvoiceURL,proto3" json:"latestStripeHostedInvoiceURL,omitempty"`
	LatestStripeInvoicePaymentIntentStatus string              `protobuf:"bytes,95,opt,name=latestStripeInvoicePaymentIntentStatus,proto3" json:"latestStripeInvoicePaymentIntentStatus,omitempty"`
	Status                                 SubscriptionStatus  `protobuf:"varint,101,opt,name=status,proto3,enum=billing.SubscriptionStatus" json:"status,omitempty"`
	Items                                  []*SubscriptionItem `protobuf:"bytes,201,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Subscription) Reset() {
	*x = Subscription{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Subscription) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Subscription) ProtoMessage() {}

func (x *Subscription) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Subscription.ProtoReflect.Descriptor instead.
func (*Subscription) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescGZIP(), []int{0}
}

func (x *Subscription) GetAccountID() string {
	if x != nil {
		return x.AccountID
	}
	return ""
}

func (x *Subscription) GetServiceID() string {
	if x != nil {
		return x.ServiceID
	}
	return ""
}

func (x *Subscription) GetServiceProviderID() string {
	if x != nil {
		return x.ServiceProviderID
	}
	return ""
}

func (x *Subscription) GetStripeCustomerID() string {
	if x != nil {
		return x.StripeCustomerID
	}
	return ""
}

func (x *Subscription) GetStripeCustomerEmail() string {
	if x != nil {
		return x.StripeCustomerEmail
	}
	return ""
}

func (x *Subscription) GetStripeSubscriptionID() string {
	if x != nil {
		return x.StripeSubscriptionID
	}
	return ""
}

func (x *Subscription) GetDefaultStripePaymentMethodID() string {
	if x != nil {
		return x.DefaultStripePaymentMethodID
	}
	return ""
}

func (x *Subscription) GetAutomaticTax() bool {
	if x != nil {
		return x.AutomaticTax
	}
	return false
}

func (x *Subscription) GetDiscount() *Discount {
	if x != nil {
		return x.Discount
	}
	return nil
}

func (x *Subscription) GetCreationDate() int64 {
	if x != nil {
		return x.CreationDate
	}
	return 0
}

func (x *Subscription) GetLastModified() int64 {
	if x != nil {
		return x.LastModified
	}
	return 0
}

func (x *Subscription) GetStartDate() int64 {
	if x != nil {
		return x.StartDate
	}
	return 0
}

func (x *Subscription) GetEndDate() int64 {
	if x != nil {
		return x.EndDate
	}
	return 0
}

func (x *Subscription) GetTrialStartDate() int64 {
	if x != nil {
		return x.TrialStartDate
	}
	return 0
}

func (x *Subscription) GetTrialEndDate() int64 {
	if x != nil {
		return x.TrialEndDate
	}
	return 0
}

func (x *Subscription) GetCancelAtPeriodEnd() bool {
	if x != nil {
		return x.CancelAtPeriodEnd
	}
	return false
}

func (x *Subscription) GetCancelationDate() int64 {
	if x != nil {
		return x.CancelationDate
	}
	return 0
}

func (x *Subscription) GetCurrentPeriodStart() int64 {
	if x != nil {
		return x.CurrentPeriodStart
	}
	return 0
}

func (x *Subscription) GetCurrentPeriodEnd() int64 {
	if x != nil {
		return x.CurrentPeriodEnd
	}
	return 0
}

func (x *Subscription) GetLatestStripeInvoiceID() string {
	if x != nil {
		return x.LatestStripeInvoiceID
	}
	return ""
}

func (x *Subscription) GetLatestStripeHostedInvoiceURL() string {
	if x != nil {
		return x.LatestStripeHostedInvoiceURL
	}
	return ""
}

func (x *Subscription) GetLatestStripeInvoicePaymentIntentStatus() string {
	if x != nil {
		return x.LatestStripeInvoicePaymentIntentStatus
	}
	return ""
}

func (x *Subscription) GetStatus() SubscriptionStatus {
	if x != nil {
		return x.Status
	}
	return SubscriptionStatus_TRIALING
}

func (x *Subscription) GetItems() []*SubscriptionItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type Discount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StripeCouponID string `protobuf:"bytes,1,opt,name=stripeCouponID,proto3" json:"stripeCouponID,omitempty"`
	PercentOff     string `protobuf:"bytes,11,opt,name=percentOff,proto3" json:"percentOff,omitempty"`
	Duration       string `protobuf:"bytes,21,opt,name=duration,proto3" json:"duration,omitempty"`
	Start          int64  `protobuf:"varint,22,opt,name=start,proto3" json:"start,omitempty"`
	End            int64  `protobuf:"varint,23,opt,name=end,proto3" json:"end,omitempty"`
}

func (x *Discount) Reset() {
	*x = Discount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Discount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Discount) ProtoMessage() {}

func (x *Discount) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Discount.ProtoReflect.Descriptor instead.
func (*Discount) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescGZIP(), []int{1}
}

func (x *Discount) GetStripeCouponID() string {
	if x != nil {
		return x.StripeCouponID
	}
	return ""
}

func (x *Discount) GetPercentOff() string {
	if x != nil {
		return x.PercentOff
	}
	return ""
}

func (x *Discount) GetDuration() string {
	if x != nil {
		return x.Duration
	}
	return ""
}

func (x *Discount) GetStart() int64 {
	if x != nil {
		return x.Start
	}
	return 0
}

func (x *Discount) GetEnd() int64 {
	if x != nil {
		return x.End
	}
	return 0
}

type SubscriptionItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ItemID   string `protobuf:"bytes,1,opt,name=itemID,proto3" json:"itemID,omitempty"`
	PriceID  string `protobuf:"bytes,11,opt,name=priceID,proto3" json:"priceID,omitempty"`
	Quantity int64  `protobuf:"varint,21,opt,name=quantity,proto3" json:"quantity,omitempty"`
}

func (x *SubscriptionItem) Reset() {
	*x = SubscriptionItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubscriptionItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubscriptionItem) ProtoMessage() {}

func (x *SubscriptionItem) ProtoReflect() protoreflect.Message {
	mi := &file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubscriptionItem.ProtoReflect.Descriptor instead.
func (*SubscriptionItem) Descriptor() ([]byte, []int) {
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescGZIP(), []int{2}
}

func (x *SubscriptionItem) GetItemID() string {
	if x != nil {
		return x.ItemID
	}
	return ""
}

func (x *SubscriptionItem) GetPriceID() string {
	if x != nil {
		return x.PriceID
	}
	return ""
}

func (x *SubscriptionItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

var File_n2x_protobuf_resources_v1_billing_subscription_proto protoreflect.FileDescriptor

var file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDesc = []byte{
	0x0a, 0x34, 0x6e, 0x32, 0x78, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x72,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x62, 0x69, 0x6c, 0x6c,
	0x69, 0x6e, 0x67, 0x2f, 0x73, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x22,
	0xda, 0x08, 0x0a, 0x0c, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x1c, 0x0a, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x44, 0x12, 0x1c,
	0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x2c, 0x0a, 0x11,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49,
	0x44, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x44, 0x12, 0x2a, 0x0a, 0x10, 0x73, 0x74,
	0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x18, 0x0b,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x10, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x65, 0x72, 0x49, 0x44, 0x12, 0x30, 0x0a, 0x13, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x0c, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x13, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x32, 0x0a, 0x14, 0x73, 0x74, 0x72, 0x69,
	0x70, 0x65, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44,
	0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x14, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x53, 0x75,
	0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x1c,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x50, 0x61, 0x79,
	0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x44, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x1c, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70,
	0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x49, 0x44,
	0x12, 0x22, 0x0a, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x54, 0x61, 0x78,
	0x18, 0x16, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69,
	0x63, 0x54, 0x61, 0x78, 0x12, 0x2d, 0x0a, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x19, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67,
	0x2e, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x08, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44,
	0x61, 0x74, 0x65, 0x18, 0x1f, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x61, 0x73, 0x74, 0x4d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x20, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x6c,
	0x61, 0x73, 0x74, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x73,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x29, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64,
	0x44, 0x61, 0x74, 0x65, 0x18, 0x2a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44,
	0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x53, 0x74, 0x61, 0x72,
	0x74, 0x44, 0x61, 0x74, 0x65, 0x18, 0x33, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0e, 0x74, 0x72, 0x69,
	0x61, 0x6c, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x74,
	0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x34, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0c, 0x74, 0x72, 0x69, 0x61, 0x6c, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12,
	0x2c, 0x0a, 0x11, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x41, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f,
	0x64, 0x45, 0x6e, 0x64, 0x18, 0x3d, 0x20, 0x01, 0x28, 0x08, 0x52, 0x11, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x41, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x12, 0x28, 0x0a,
	0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x3e, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2e, 0x0a, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x47, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x12, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69,
	0x6f, 0x64, 0x53, 0x74, 0x61, 0x72, 0x74, 0x12, 0x2a, 0x0a, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65,
	0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64, 0x45, 0x6e, 0x64, 0x18, 0x48, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x10, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x50, 0x65, 0x72, 0x69, 0x6f, 0x64,
	0x45, 0x6e, 0x64, 0x12, 0x34, 0x0a, 0x15, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x5b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x15, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65,
	0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x42, 0x0a, 0x1c, 0x6c, 0x61, 0x74,
	0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x65, 0x64, 0x49,
	0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x18, 0x5c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x1c, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x48, 0x6f, 0x73,
	0x74, 0x65, 0x64, 0x49, 0x6e, 0x76, 0x6f, 0x69, 0x63, 0x65, 0x55, 0x52, 0x4c, 0x12, 0x56, 0x0a,
	0x26, 0x6c, 0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x6e, 0x76,
	0x6f, 0x69, 0x63, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x5f, 0x20, 0x01, 0x28, 0x09, 0x52, 0x26, 0x6c,
	0x61, 0x74, 0x65, 0x73, 0x74, 0x53, 0x74, 0x72, 0x69, 0x70, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x69,
	0x63, 0x65, 0x50, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x65, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x2e,
	0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x30, 0x0a, 0x05, 0x69, 0x74,
	0x65, 0x6d, 0x73, 0x18, 0xc9, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x62, 0x69, 0x6c,
	0x6c, 0x69, 0x6e, 0x67, 0x2e, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0x96, 0x01, 0x0a,
	0x08, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x26, 0x0a, 0x0e, 0x73, 0x74, 0x72,
	0x69, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0e, 0x73, 0x74, 0x72, 0x69, 0x70, 0x65, 0x43, 0x6f, 0x75, 0x70, 0x6f, 0x6e, 0x49,
	0x44, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x4f, 0x66, 0x66, 0x18,
	0x0b, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x65, 0x72, 0x63, 0x65, 0x6e, 0x74, 0x4f, 0x66,
	0x66, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x15, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x73, 0x74, 0x61, 0x72, 0x74, 0x18, 0x16, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x65, 0x6e, 0x64, 0x18, 0x17, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x03, 0x65, 0x6e, 0x64, 0x22, 0x60, 0x0a, 0x10, 0x53, 0x75, 0x62, 0x73, 0x63, 0x72, 0x69,
	0x70, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x74, 0x65,
	0x6d, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x69, 0x74, 0x65, 0x6d, 0x49,
	0x44, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x18, 0x0b, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x70, 0x72, 0x69, 0x63, 0x65, 0x49, 0x44, 0x12, 0x1a, 0x0a, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x15, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71,
	0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2a, 0x95, 0x01, 0x0a, 0x12, 0x53, 0x75, 0x62, 0x73,
	0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x0c,
	0x0a, 0x08, 0x54, 0x52, 0x49, 0x41, 0x4c, 0x49, 0x4e, 0x47, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06,
	0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x0e, 0x0a, 0x0a, 0x49, 0x4e, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x10, 0x02, 0x12, 0x16, 0x0a, 0x12, 0x49, 0x4e, 0x43, 0x4f,
	0x4d, 0x50, 0x4c, 0x45, 0x54, 0x45, 0x5f, 0x45, 0x58, 0x50, 0x49, 0x52, 0x45, 0x44, 0x10, 0x03,
	0x12, 0x0c, 0x0a, 0x08, 0x50, 0x41, 0x53, 0x54, 0x5f, 0x44, 0x55, 0x45, 0x10, 0x04, 0x12, 0x0c,
	0x0a, 0x08, 0x43, 0x41, 0x4e, 0x43, 0x45, 0x4c, 0x45, 0x44, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06,
	0x55, 0x4e, 0x50, 0x41, 0x49, 0x44, 0x10, 0x06, 0x12, 0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x4b, 0x4e,
	0x4f, 0x57, 0x4e, 0x10, 0x64, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x4e, 0x45, 0x10, 0x65, 0x42,
	0x29, 0x5a, 0x27, 0x6e, 0x32, 0x78, 0x2e, 0x64, 0x65, 0x76, 0x2f, 0x78, 0x2d, 0x61, 0x70, 0x69,
	0x2d, 0x67, 0x6f, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x73, 0x2f, 0x62, 0x69, 0x6c, 0x6c, 0x69, 0x6e, 0x67, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescOnce sync.Once
	file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescData = file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDesc
)

func file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescGZIP() []byte {
	file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescOnce.Do(func() {
		file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescData = protoimpl.X.CompressGZIP(file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescData)
	})
	return file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDescData
}

var file_n2x_protobuf_resources_v1_billing_subscription_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_n2x_protobuf_resources_v1_billing_subscription_proto_goTypes = []any{
	(SubscriptionStatus)(0),  // 0: billing.SubscriptionStatus
	(*Subscription)(nil),     // 1: billing.Subscription
	(*Discount)(nil),         // 2: billing.Discount
	(*SubscriptionItem)(nil), // 3: billing.SubscriptionItem
}
var file_n2x_protobuf_resources_v1_billing_subscription_proto_depIdxs = []int32{
	2, // 0: billing.Subscription.discount:type_name -> billing.Discount
	0, // 1: billing.Subscription.status:type_name -> billing.SubscriptionStatus
	3, // 2: billing.Subscription.items:type_name -> billing.SubscriptionItem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_n2x_protobuf_resources_v1_billing_subscription_proto_init() }
func file_n2x_protobuf_resources_v1_billing_subscription_proto_init() {
	if File_n2x_protobuf_resources_v1_billing_subscription_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Subscription); i {
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
		file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*Discount); i {
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
		file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*SubscriptionItem); i {
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
			RawDescriptor: file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_n2x_protobuf_resources_v1_billing_subscription_proto_goTypes,
		DependencyIndexes: file_n2x_protobuf_resources_v1_billing_subscription_proto_depIdxs,
		EnumInfos:         file_n2x_protobuf_resources_v1_billing_subscription_proto_enumTypes,
		MessageInfos:      file_n2x_protobuf_resources_v1_billing_subscription_proto_msgTypes,
	}.Build()
	File_n2x_protobuf_resources_v1_billing_subscription_proto = out.File
	file_n2x_protobuf_resources_v1_billing_subscription_proto_rawDesc = nil
	file_n2x_protobuf_resources_v1_billing_subscription_proto_goTypes = nil
	file_n2x_protobuf_resources_v1_billing_subscription_proto_depIdxs = nil
}
