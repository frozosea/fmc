import RequestSenderException from "../../exceptions";

const fetch = require("cross-fetch");

interface Proxy {
    server: string,
    bypass?: string,
    username?: string,
    password?: string
}

export interface _BaseRequestSenderArgs {
    url: string,
    requestTimeOut?: number,
    UserAgent?: string,
    proxy?: Proxy
}

export interface fetchArgs extends _BaseRequestSenderArgs {
    method: "GET" | "POST" | "PUT" | "PATCH" | "DELETE",
    headers?: any,
    body?: string,
}


export interface IRequest<T extends _BaseRequestSenderArgs> {
    sendRequestAndGetJson(args: T): Promise<any>;

    sendRequestAndGetHtml(args: T): Promise<string>;
}

export class RequestSender implements IRequest<fetchArgs> {
    private async sendRequest(args: fetchArgs): Promise<any> {
        return await fetch(args.url, {"method": args.method, "body": args.body, "headers": args.headers})
    }

    public async sendRequestAndGetHtml(args: fetchArgs): Promise<string> {
        let response = await this.sendRequest(args);
        if (response.status === 200) {
            return await response.text()
        } else {
            throw new RequestSenderException()
        }
    }

    public async sendRequestAndGetJson(args: fetchArgs): Promise<any> {
        let response = await this.sendRequest(args)
        if (response.status === 200 || response.status === 201) {
            try {
                return await response.json()
            } catch (e) {
                return JSON.parse(await response.text())
            }
        } else {
            throw new RequestSenderException()
        }
    }
}


