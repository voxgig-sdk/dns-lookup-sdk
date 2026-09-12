import { DnsLookupEntityBase } from '../DnsLookupEntityBase';
import type { DnsLookupSDK } from '../DnsLookupSDK';
import type { Control } from '../types';
import type { Ssl, SslListMatch } from '../DnsLookupTypes';
declare class SslEntity extends DnsLookupEntityBase<Ssl> {
    constructor(client: DnsLookupSDK, entopts: any);
    make(this: SslEntity): SslEntity;
    list(this: any, reqmatch?: SslListMatch, ctrl?: Control): Promise<SslEntity[]>;
}
export { SslEntity };
