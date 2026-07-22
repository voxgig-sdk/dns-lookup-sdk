# DnsLookup SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module DnsLookupFeatures
  def self.make_feature(name)
    case name
    when "base"
      DnsLookupBaseFeature.new
    when "test"
      DnsLookupTestFeature.new
    else
      DnsLookupBaseFeature.new
    end
  end
end
