# DnsLookup SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/ratelimit_feature'
require_relative 'feature/retry_feature'
require_relative 'feature/test_feature'
require_relative 'feature/timeout_feature'


module DnsLookupFeatures
  def self.make_feature(name)
    case name
    when "base"
      DnsLookupBaseFeature.new
    when "ratelimit"
      DnsLookupRatelimitFeature.new
    when "retry"
      DnsLookupRetryFeature.new
    when "test"
      DnsLookupTestFeature.new
    when "timeout"
      DnsLookupTimeoutFeature.new
    else
      DnsLookupBaseFeature.new
    end
  end
end
