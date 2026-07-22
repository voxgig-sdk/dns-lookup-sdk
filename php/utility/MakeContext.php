<?php
declare(strict_types=1);

// DnsLookup SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class DnsLookupMakeContext
{
    public static function call(array $ctxmap, ?DnsLookupContext $basectx): DnsLookupContext
    {
        return new DnsLookupContext($ctxmap, $basectx);
    }
}
