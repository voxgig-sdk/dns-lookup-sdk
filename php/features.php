<?php
declare(strict_types=1);

// DnsLookup SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class DnsLookupFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new DnsLookupBaseFeature();
            case "test":
                return new DnsLookupTestFeature();
            default:
                return new DnsLookupBaseFeature();
        }
    }
}
