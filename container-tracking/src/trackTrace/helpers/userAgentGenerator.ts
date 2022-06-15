export interface IUserAgentGenerator {
    generateUserAgent(): string;
}

export class UserAgentGenerator implements IUserAgentGenerator {
    public generateUserAgent(): string {
        return ""
    }
}