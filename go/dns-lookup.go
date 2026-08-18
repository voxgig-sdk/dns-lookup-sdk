package voxgigdnslookupsdk

import (
	"github.com/voxgig-sdk/dns-lookup-sdk/go/core"
	"github.com/voxgig-sdk/dns-lookup-sdk/go/entity"
	"github.com/voxgig-sdk/dns-lookup-sdk/go/feature"
	_ "github.com/voxgig-sdk/dns-lookup-sdk/go/utility"
)

// Type aliases preserve external API.
type DnsLookupSDK = core.DnsLookupSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type DnsLookupEntity = core.DnsLookupEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type DnsLookupError = core.DnsLookupError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewDnsResultEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewDnsResultEntity(client, entopts)
	}
	core.NewDomainEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewDomainEntity(client, entopts)
	}
	core.NewEmailValidateEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewEmailValidateEntity(client, entopts)
	}
	core.NewGenerateEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewGenerateEntity(client, entopts)
	}
	core.NewGrammarEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewGrammarEntity(client, entopts)
	}
	core.NewIpnEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewIpnEntity(client, entopts)
	}
	core.NewRedactEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewRedactEntity(client, entopts)
	}
	core.NewSslEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewSslEntity(client, entopts)
	}
	core.NewUtilityEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewUtilityEntity(client, entopts)
	}
	core.NewWhoiEntityFunc = func(client *core.DnsLookupSDK, entopts map[string]any) core.DnsLookupEntity {
		return entity.NewWhoiEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewDnsLookupSDK = core.NewDnsLookupSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig
var SharedConfig = core.SharedConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewDnsLookupSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *DnsLookupSDK  { return NewDnsLookupSDK(nil) }
func Test() *DnsLookupSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
