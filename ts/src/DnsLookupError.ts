
import { Context } from './Context'


class DnsLookupError extends Error {

  isDnsLookupError = true

  sdk = 'DnsLookup'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  DnsLookupError
}

