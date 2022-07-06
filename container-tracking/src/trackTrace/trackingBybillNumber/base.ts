import {ITrackingArgs, ITrackingByBillNumberResponse} from "../../types";


export interface IBillNumberTracker {
    trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse>;
}