# DnsLookup SDK feature factory

from dnslookup_sdk.feature.base_feature import DnsLookupBaseFeature
from dnslookup_sdk.feature.ratelimit_feature import DnsLookupRatelimitFeature
from dnslookup_sdk.feature.retry_feature import DnsLookupRetryFeature
from dnslookup_sdk.feature.test_feature import DnsLookupTestFeature
from dnslookup_sdk.feature.timeout_feature import DnsLookupTimeoutFeature


_FEATURES = {
    "base": lambda: DnsLookupBaseFeature(),
    "ratelimit": lambda: DnsLookupRatelimitFeature(),
    "retry": lambda: DnsLookupRetryFeature(),
    "test": lambda: DnsLookupTestFeature(),
    "timeout": lambda: DnsLookupTimeoutFeature(),
}


def _make_feature(name):
    factory = _FEATURES.get(name)
    if factory is not None:
        return factory()
    return _FEATURES["base"]()


# True when this SDK was generated with the named feature class - the
# constructor's tolerance for extend-carried features reads this (an
# active name with no generated class must not become a BaseFeature
# stray when an extend instance carries it).
def _has_feature(name):
    return name in _FEATURES
