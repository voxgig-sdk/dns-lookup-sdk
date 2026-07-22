# DnsLookup SDK feature factory

from feature.base_feature import DnsLookupBaseFeature
from feature.test_feature import DnsLookupTestFeature


def _make_feature(name):
    features = {
        "base": lambda: DnsLookupBaseFeature(),
        "test": lambda: DnsLookupTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
