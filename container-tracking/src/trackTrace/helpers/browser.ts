import {Browser, Page, webkit} from "playwright";
import RequestSenderException from "../../exceptions";
import {_BaseRequestSenderArgs, IRequest} from "./requestSender";

interface waitingSelector {
    selector: string
    timeout: number
}

interface TypeToIntoField {
    selector: string
    text: any
    clickAfterType?: {
        locatorSelector: string
        key: string
        waitingSelector?: waitingSelector
    }
}

export interface IBrowserArgs extends _BaseRequestSenderArgs {
    waitingSelector?: string
    clickSelectors?: string[]
    TypeToIntoField?: TypeToIntoField
    visibleSelector?: string
}

export class BrowserRequestSender implements IRequest<IBrowserArgs> {
    private browser: Browser;

    public constructor(develop: boolean) {
        if (!develop) {
            this.launch().then((browse: Browser) => {
                this.browser = browse;
            }).catch(() => {
            })
        }
    }

    protected async launch(): Promise<Browser> {
        return await webkit.launch({headless: false})
    }

    // private async blockedResources(page: Page): Promise<void> {
    //     const blockedResource = BLOCKING_RESOURCE_TYPES
    //     await page.route("**/*", (route) => {
    //         return blockedResource.includes(route.request().resourceType()) ? route.abort() : route.continue()
    //     })
    // }

    protected async sendRequestDecorator(fn: Function, args: _BaseRequestSenderArgs): Promise<Function> {
        return async () => {
            if (this.browser === undefined) {
                this.browser = await webkit.launch({headless: false})
            }
            const context = await this.browser.newContext(args)
            const page = await context.newPage()
            try {
                await fn(page);
            } catch (e) {
            } finally {
                await context.close();
            }
        }
    }

    public async sendRequestAndGetHtml(args: IBrowserArgs): Promise<string> {
        let content: string = ""
        let func = await this.sendRequestDecorator(async function (page: Page) {
            let resp = await page.goto(args.url, {timeout: args.requestTimeOut, waitUntil: "domcontentloaded"})
            if (args.TypeToIntoField) {
                await page.type(args.TypeToIntoField.selector, args.TypeToIntoField.text)
                if (args.TypeToIntoField.clickAfterType) {
                    await page.locator(args.TypeToIntoField.clickAfterType.locatorSelector).press(args.TypeToIntoField.clickAfterType.key)
                    await page.waitForSelector(args.TypeToIntoField.clickAfterType.waitingSelector.selector, {timeout: args.TypeToIntoField.clickAfterType.waitingSelector.timeout})
                    content = await page.content()
                } else {
                    content = await page.content()
                }
            } else {
                content = await resp.text();
            }
            await page.close()
        }, args);
        await func()
        return content

    }

    public async sendRequestAndGetJson(args: IBrowserArgs): Promise<any> {
        let content: object = {}
        let func = await this.sendRequestDecorator(
            async (page: Page) => {
                let resp = await page.goto(args.url, {timeout: args.requestTimeOut, waitUntil: "domcontentloaded"})
                content = await resp.json()
                await page.close()
            }, args
        )
        await func()
        if (content !== {}) {
            return content
        } else {
            throw new RequestSenderException()
        }
    }

    public async sendRequestAndGetImage(args: IBrowserArgs): Promise<any> {
        return
    }
}