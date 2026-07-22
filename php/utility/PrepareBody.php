<?php
declare(strict_types=1);

// DnsLookup SDK utility: prepare_body

class DnsLookupPrepareBody
{
    public static function call(DnsLookupContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
