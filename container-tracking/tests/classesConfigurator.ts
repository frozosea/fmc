import {fetchArgs, IRequest, RequestSender} from "../src/trackTrace/helpers/requestSender";
import {IUserAgentGenerator, UserAgentGenerator} from "../src/trackTrace/helpers/userAgentGenerator";
import {IUnlocodesRepo, UnlocodesRepo} from "../src/trackTrace/TrackingByContainerNumber/sklu/unlocodesRepo";
import {Datetime, IDatetime} from "../src/trackTrace/helpers/datetime";
import {BrowserRequestSender, IBrowserArgs} from "../src/trackTrace/helpers/browser";

interface ConfigConstructorArgs {
    userAgentGenerator: IUserAgentGenerator,
    requestSender: IRequest<fetchArgs>,
    browserRequestSender: IRequest<IBrowserArgs>,
    unlocodesRepo: IUnlocodesRepo,
    datetime: IDatetime

}

class ClassesConfigurator {
    public USER_AGENT_GENERATOR: IUserAgentGenerator;
    public REQUEST_SENDER: IRequest<fetchArgs>;
    public BROWSER: IRequest<IBrowserArgs>;
    public UNLOCODES_REPO: IUnlocodesRepo;
    public DATETIME: IDatetime

    public constructor(args: ConfigConstructorArgs) {
        this.USER_AGENT_GENERATOR = args.userAgentGenerator;
        this.REQUEST_SENDER = args.requestSender;
        this.BROWSER = args.browserRequestSender;
        this.UNLOCODES_REPO = args.unlocodesRepo;
        this.DATETIME = args.datetime;
    }

}

export const config = new ClassesConfigurator({
    userAgentGenerator: new UserAgentGenerator(),
    requestSender: new RequestSender(),
    browserRequestSender: new BrowserRequestSender(true),
    unlocodesRepo: new UnlocodesRepo(),
    datetime: new Datetime()
})