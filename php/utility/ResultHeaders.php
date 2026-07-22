<?php
declare(strict_types=1);

// DnsLookup SDK utility: result_headers

class DnsLookupResultHeaders
{
    public static function call(DnsLookupContext $ctx): ?DnsLookupResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
