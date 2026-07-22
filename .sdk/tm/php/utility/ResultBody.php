<?php
declare(strict_types=1);

// DnsLookup SDK utility: result_body

class DnsLookupResultBody
{
    public static function call(DnsLookupContext $ctx): ?DnsLookupResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
