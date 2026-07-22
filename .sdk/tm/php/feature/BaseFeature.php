<?php
declare(strict_types=1);

// DnsLookup SDK base feature

class DnsLookupBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(DnsLookupContext $ctx, array $options): void {}
    public function PostConstruct(DnsLookupContext $ctx): void {}
    public function PostConstructEntity(DnsLookupContext $ctx): void {}
    public function SetData(DnsLookupContext $ctx): void {}
    public function GetData(DnsLookupContext $ctx): void {}
    public function GetMatch(DnsLookupContext $ctx): void {}
    public function SetMatch(DnsLookupContext $ctx): void {}
    public function PrePoint(DnsLookupContext $ctx): void {}
    public function PreSpec(DnsLookupContext $ctx): void {}
    public function PreRequest(DnsLookupContext $ctx): void {}
    public function PreResponse(DnsLookupContext $ctx): void {}
    public function PreResult(DnsLookupContext $ctx): void {}
    public function PreDone(DnsLookupContext $ctx): void {}
    public function PreUnexpected(DnsLookupContext $ctx): void {}
}
