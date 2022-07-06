import {sendUnaryData, ServerUnaryCall} from '@grpc/grpc-js';
import ContainerTrackingController from "../../containerTrackingController";
import {ITrackingByContainerNumberServer} from "../proto/server_grpc_pb";
import {Country, InfoAboutMoving, Request, Scac, TrackingByContainerNumberResponse} from "../proto/server_pb";
import {COUNTRY_TYPE, OneTrackingEvent, SCAC_TYPE, TrackingContainerResponse} from "../../types";
import {ILogger} from "../../logging";


export class TrackingServiceConverter {
    public static convertEnumIntoScacType(scac: Scac): SCAC_TYPE {
        let obj = {}
        for (let key in Object.keys(Scac)) {
            obj[key] = Object.keys(Scac)[key]
        }
        return obj[scac]
    }

    public static convertScacIntoEnum(scac: SCAC_TYPE): Scac {
        let obj = {}
        for (let key in Object.keys(Scac)) {
            obj[Object.keys(Scac)[key]] = Object.values(Scac)[key]
        }
        return obj[scac]
    }

    public static convertEnumCountryIntoCountryType(cntry: Country): COUNTRY_TYPE {
        let obj = {}
        for (let key in Object.keys(Country)) {
            obj[key] = Object.keys(Country)[key]
        }
        return obj[cntry]
    }

    public static convertCountryTypeIntoEnum(coutry: COUNTRY_TYPE): Country {
        let obj = {}
        for (let key in Object.keys(Country)) {
            obj[Object.keys(Country)[key]] = Object.values(Country)[key]
        }
        return obj[coutry]
    }
}


export class ServiceSerializer {
    protected serializeInfoAboutMoving(response: OneTrackingEvent[]) {
        let grpcInfoAboutMovingArray: InfoAboutMoving[] = []
        for (let item of response) {
            let grpcInfoAboutMoving = new InfoAboutMoving()
            grpcInfoAboutMoving.setTime(item.time)
            grpcInfoAboutMoving.setOperationName(item.operationName)
            grpcInfoAboutMoving.setLocation(item.location)
            grpcInfoAboutMoving.setVessel(item.vessel)
            grpcInfoAboutMovingArray.push(grpcInfoAboutMoving)
        }
        return grpcInfoAboutMovingArray
    }

    public serializeResponseIntoGrpc(response: TrackingContainerResponse): TrackingByContainerNumberResponse {
        let grpcEmptyResp = new TrackingByContainerNumberResponse()
        grpcEmptyResp.setInfoAboutMovingList(this.serializeInfoAboutMoving(response.infoAboutMoving))
        grpcEmptyResp.setContainerSize(response.containerSize)
        grpcEmptyResp.setContainer(response.container)
        grpcEmptyResp.setScac(TrackingServiceConverter.convertScacIntoEnum(response.scac))
        return grpcEmptyResp
    }
}

export class TrackingByContainerNumberService implements ITrackingByContainerNumberServer {
    protected trackingController: ContainerTrackingController;
    protected logger: ILogger;
    protected serializer: ServiceSerializer;

    public constructor(trackingController: ContainerTrackingController, logger: ILogger) {
        this.logger = logger;
        this.trackingController = trackingController;
        this.serializer = new ServiceSerializer()
    }

    public trackByContainerNumber(call: ServerUnaryCall<Request, TrackingByContainerNumberResponse>, callback: sendUnaryData<TrackingByContainerNumberResponse>): void {
        let container: string = call.request.getNumber()
        let scac: SCAC_TYPE = TrackingServiceConverter.convertEnumIntoScacType(call.request.getScac())
        let country = TrackingServiceConverter.convertEnumCountryIntoCountryType(call.request.getCountry())
        this.logger.InfoLog(`${container}: ${scac}`)
        this.trackingController.trackContainer({
            number: container,
            scac: scac,
            country: country
        }).then((result: TrackingContainerResponse) => {
                callback(null, this.serializer.serializeResponseIntoGrpc(result))
            }
        ).catch((error: Error) => {
            this.logger.ExceptionLog(`find container error: ${error}`)
            callback(error, null)
        })
        return
    }
}

