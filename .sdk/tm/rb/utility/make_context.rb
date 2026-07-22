# DnsLookup SDK utility: make_context
require_relative '../core/context'
module DnsLookupUtilities
  MakeContext = ->(ctxmap, basectx) {
    DnsLookupContext.new(ctxmap, basectx)
  }
end
