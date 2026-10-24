# DnsLookup SDK utility: prepare_body
require_relative 'media'
module DnsLookupUtilities
  PrepareBody = ->(ctx) {
    return nil unless ctx.op.input == "data"
    return DnsLookupUtilities.raw_body(ctx.reqdata) if DnsLookupUtilities.raw_request?(ctx.point)
    ctx.utility.transform_request.call(ctx)
  }
end
