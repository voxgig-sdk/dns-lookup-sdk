import { DnsLookupEntityBase } from '../DnsLookupEntityBase';
import type { DnsLookupSDK } from '../DnsLookupSDK';
import type { Control } from '../types';
import type { Generate, GenerateLoadMatch } from '../DnsLookupTypes';
declare class GenerateEntity extends DnsLookupEntityBase<Generate> {
    constructor(client: DnsLookupSDK, entopts: any);
    make(this: GenerateEntity): GenerateEntity;
    load(this: any, reqmatch?: GenerateLoadMatch, ctrl?: Control): Promise<GenerateEntity>;
}
export { GenerateEntity };
