
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { DnsLookupSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await DnsLookupSDK.test()
    equal(null !== testsdk, true)
  })

})
