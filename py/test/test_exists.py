# DnsLookup SDK exists test

import pytest
from dnslookup_sdk import DnsLookupSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = DnsLookupSDK.test(None, None)
        assert testsdk is not None
