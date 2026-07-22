-- DnsLookup SDK error

local DnsLookupError = {}
DnsLookupError.__index = DnsLookupError


function DnsLookupError.new(code, msg, ctx)
  local self = setmetatable({}, DnsLookupError)
  self.is_sdk_error = true
  self.sdk = "DnsLookup"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function DnsLookupError:error()
  return self.msg
end


function DnsLookupError:__tostring()
  return self.msg
end


return DnsLookupError
