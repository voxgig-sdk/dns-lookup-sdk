import { DnsLookupEntityBase } from '../DnsLookupEntityBase';
import type { DnsLookupSDK } from '../DnsLookupSDK';
import type { Control } from '../types';
import type { Whoi, WhoiListMatch } from '../DnsLookupTypes';
declare class WhoiEntity extends DnsLookupEntityBase<Whoi> {
    constructor(client: DnsLookupSDK, entopts: any);
    make(this: WhoiEntity): WhoiEntity;
    list(this: any, reqmatch?: WhoiListMatch, ctrl?: Control): Promise<WhoiEntity[]>;
}
export { WhoiEntity };
