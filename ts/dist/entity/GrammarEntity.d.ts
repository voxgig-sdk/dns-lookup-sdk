import { DnsLookupEntityBase } from '../DnsLookupEntityBase';
import type { DnsLookupSDK } from '../DnsLookupSDK';
import type { Control } from '../types';
import type { Grammar, GrammarCreateData } from '../DnsLookupTypes';
declare class GrammarEntity extends DnsLookupEntityBase<Grammar> {
    constructor(client: DnsLookupSDK, entopts: any);
    make(this: GrammarEntity): GrammarEntity;
    create(this: any, reqdata?: GrammarCreateData, ctrl?: Control): Promise<GrammarEntity>;
}
export { GrammarEntity };
