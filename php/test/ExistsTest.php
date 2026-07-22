<?php
declare(strict_types=1);

// DnsLookup SDK exists test

require_once __DIR__ . '/../dnslookup_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = DnsLookupSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
