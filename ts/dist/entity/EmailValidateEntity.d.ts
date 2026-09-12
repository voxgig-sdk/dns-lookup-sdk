import { DnsLookupEntityBase } from '../DnsLookupEntityBase';
import type { DnsLookupSDK } from '../DnsLookupSDK';
import type { Control } from '../types';
import type { EmailValidate, EmailValidateLoadMatch } from '../DnsLookupTypes';
declare class EmailValidateEntity extends DnsLookupEntityBase<EmailValidate> {
    constructor(client: DnsLookupSDK, entopts: any);
    make(this: EmailValidateEntity): EmailValidateEntity;
    load(this: any, reqmatch?: EmailValidateLoadMatch, ctrl?: Control): Promise<EmailValidateEntity>;
}
export { EmailValidateEntity };
