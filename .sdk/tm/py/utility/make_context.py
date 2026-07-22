# DnsLookup SDK utility: make_context

from core.context import DnsLookupContext


def make_context_util(ctxmap, basectx):
    return DnsLookupContext(ctxmap, basectx)
