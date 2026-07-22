package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewDnsResultEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewDomainEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewEmailValidateEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewGenerateEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewGrammarEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewIpnEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewRedactEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewSslEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewUtilityEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

var NewWhoiEntityFunc func(client *DnsLookupSDK, entopts map[string]any) DnsLookupEntity

