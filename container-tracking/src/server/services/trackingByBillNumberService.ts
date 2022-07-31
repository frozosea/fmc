import {ITrackingByBillNumberServer} from "../proto/tracking_grpc_pb";
import {Request, TrackingByBillNumberResponse} from "../proto/tracking_pb";
import {sendUnaryData, ServerUnaryCall} from '@grpc/grpc-js';
import {ILogger} from "../../logging";
import {ServiceSerializer, TrackingServiceConverter} from "./trackingByContainerNumberService";
import BillNumberTrackingController from "../../trackingByBillNumberController";
import {ITrackingByBillNumberResponse, SCAC_TYPE} from "../../types";


export class BillNumberServiceSerializer extends ServiceSerializer {
    public serializeBillNumberResponseIntoGrpc(response: ITrackingByBillNumberResponse): TrackingByBillNumberResponse {
        let grpcEmptyResp = new TrackingByBillNumberResponse()
        grpcEmptyResp.setInfoAboutMovingList(this.serializeInfoAboutMoving(response.infoAboutMoving))
        grpcEmptyResp.setEtaFinalDelivery(response.etaFinalDelivery)
        grpcEmptyResp.setBillno(response.billNo)
        grpcEmptyResp.setScac(TrackingServiceConverter.convertScacIntoEnum(response.scac))
        return grpcEmptyResp
    }
}

export class TrackingBybillNumberService implements ITrackingByBillNumberServer {
    protected trackingController: BillNumberTrackingController;
    protected logger: ILogger;
    protected serializer: BillNumberServiceSerializer;

    public constructor(trackingController: BillNumberTrackingController, logger: ILogger) {
        this.logger = logger;
        this.trackingController = trackingController;
        this.serializer = new BillNumberServiceSerializer()
    }

    public trackByBillNumber(call: ServerUnaryCall<Request, TrackingByBillNumberResponse>, callback: sendUnaryData<TrackingByBillNumberResponse>) {
        let container: string = call.request.getNumber()
        let scac: SCAC_TYPE = TrackingServiceConverter.convertEnumIntoScacType(call.request.getScac())
        let country = TrackingServiceConverter.convertEnumCountryIntoCountryType(call.request.getCountry())
        this.logger.InfoLog(`${container}: ${scac}`)
        this.trackingController.trackByBillNumber({
            number: container,
            scac: scac,
            country: country
        }).then((result: ITrackingByBillNumberResponse) => {
                callback(null, this.serializer.serializeBillNumberResponseIntoGrpc(result))
                return
            }
        ).catch((error: Error) => {
            this.logger.ExceptionLog(`find container error: ${error}`)
            callback(error, null)
            return
        })
        return
    }
}