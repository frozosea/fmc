import {SkluBillNumber} from "../sklu/sklu";
import {BaseContainerConstructor, ITrackingArgs} from "../../base";
import {fetchArgs} from "../../helpers/requestSender";
import {IUnlocodesRepo} from "../../TrackingByContainerNumber/sklu/unlocodesRepo";
import {ITrackingByBillNumberResponse} from "../../../types";
import {HaluRequest} from "../../TrackingByContainerNumber/halu/halu";





export class HaluBillNumber extends SkluBillNumber {
    public constructor(args: BaseContainerConstructor<fetchArgs>, unlocodesRepo: IUnlocodesRepo) {
        super(args, unlocodesRepo);
        this.skluRequest = new HaluRequest(args.requestSender, args.UserAgentGenerator)
    }

    public async trackByBillNumber(args: ITrackingArgs): Promise<ITrackingByBillNumberResponse> {
        let result = await super.trackByBillNumber(args);
        result.scac = "HALU"
        return result
    }
}