
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DnsLookupSDK } from '..'


describe('exists', async () => {

  test('test-mode', () => {
    const testsdk = DnsLookupSDK.test()
    equal(testsdk instanceof DnsLookupSDK, true,
      'DnsLookupSDK.test() must return a client synchronously')
  })

})
