import {IUserAgentGenerator} from "./helpers/userAgentGenerator";
import {_BaseRequestSenderArgs, IRequest} from "./helpers/requestSender";
import {IDatetime} from "./helpers/datetime";
import {SCAC_TYPE} from "../types";


export interface OneTrackingEvent {
    time: number
    location: string
    operationName: string
    vessel: string

}

export interface TrackingContainerResponse {
    container: string
    containerSize: string
    scac: SCAC_TYPE
    infoAboutMoving: OneTrackingEvent[]
}

export interface ITrackingArgs {
    number: string
}

export interface TrackingArgsWithScac extends ITrackingArgs {
    scac: SCAC_TYPE
}

export interface ITrackingByContainerNumber {
    trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse>;
}

export interface BaseContainerConstructor<T extends _BaseRequestSenderArgs> {
    UserAgentGenerator: IUserAgentGenerator
    requestSender: IRequest<T>
    datetime: IDatetime
}

export abstract class BaseTrackerByContainerNumber<T extends _BaseRequestSenderArgs> implements ITrackingByContainerNumber {
    protected UserAgentGenerator: IUserAgentGenerator;
    protected requestSender: IRequest<T>;
    protected datetime: IDatetime;

    protected constructor(args: BaseContainerConstructor<T>) {
        this.requestSender = args.requestSender;
        this.UserAgentGenerator = args.UserAgentGenerator;
        this.datetime = args.datetime;
    }

    public abstract trackContainer(args: ITrackingArgs): Promise<TrackingContainerResponse>
}