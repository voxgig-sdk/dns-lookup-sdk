import { Context } from './Context';
declare class DnsLookupError extends Error {
    isDnsLookupError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { DnsLookupError };
