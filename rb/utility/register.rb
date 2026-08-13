# DnsLookup SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'graphql'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

DnsLookupUtility.registrar = ->(u) {
  u.clean = DnsLookupUtilities::Clean
  u.done = DnsLookupUtilities::Done
  u.make_error = DnsLookupUtilities::MakeError
  u.feature_add = DnsLookupUtilities::FeatureAdd
  u.feature_hook = DnsLookupUtilities::FeatureHook
  u.feature_init = DnsLookupUtilities::FeatureInit
  u.fetcher = DnsLookupUtilities::Fetcher
  u.make_fetch_def = DnsLookupUtilities::MakeFetchDef
  u.make_context = DnsLookupUtilities::MakeContext
  u.make_options = DnsLookupUtilities::MakeOptions
  u.make_request = DnsLookupUtilities::MakeRequest
  u.make_response = DnsLookupUtilities::MakeResponse
  u.make_result = DnsLookupUtilities::MakeResult
  u.make_point = DnsLookupUtilities::MakePoint
  u.make_spec = DnsLookupUtilities::MakeSpec
  u.make_url = DnsLookupUtilities::MakeUrl
  u.param = DnsLookupUtilities::Param
  u.prepare_auth = DnsLookupUtilities::PrepareAuth
  u.prepare_body = DnsLookupUtilities::PrepareBody
  u.prepare_headers = DnsLookupUtilities::PrepareHeaders
  u.prepare_method = DnsLookupUtilities::PrepareMethod
  u.prepare_params = DnsLookupUtilities::PrepareParams
  u.prepare_path = DnsLookupUtilities::PreparePath
  u.prepare_query = DnsLookupUtilities::PrepareQuery
  u.graphql_body = DnsLookupUtilities::GraphqlBody
  u.graphql_errors = DnsLookupUtilities::GraphqlErrors
  u.result_basic = DnsLookupUtilities::ResultBasic
  u.result_body = DnsLookupUtilities::ResultBody
  u.result_headers = DnsLookupUtilities::ResultHeaders
  u.transform_request = DnsLookupUtilities::TransformRequest
  u.transform_response = DnsLookupUtilities::TransformResponse
}
