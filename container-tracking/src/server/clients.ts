import {TrackingByBillNumberClient, TrackingByContainerNumberClient} from "./proto/server_grpc_pb";
import {GRPC_PORT} from "../../config.json";
import {credentials} from "@grpc/grpc-js";
import {Request, TrackingByBillNumberResponse, TrackingByContainerNumberResponse} from "./proto/server_pb";
import {COUNTRY_TYPE, SCAC_TYPE} from "../types";
import {TrackingServiceConverter} from "./services/trackingByContainerNumberService";

export const trackingByContainerNumberClient = new TrackingByContainerNumberClient(
    `localhost:${GRPC_PORT}`,
    credentials.createInsecure(),
);
export const trackingByBillNumberClient = new TrackingByBillNumberClient(`localhost:${GRPC_PORT}`,
    credentials.createInsecure(),)

export function trackContainerByServer(container: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<TrackingByContainerNumberResponse> {
    return new Promise<TrackingByContainerNumberResponse>((resolve, reject) => {
        const request = new Request();
        request.setNumber(container);
        request.setScac(TrackingServiceConverter.convertScacIntoEnum(scac))
        request.setCountry(TrackingServiceConverter.convertCountryTypeIntoEnum(country))
        trackingByContainerNumberClient.trackByContainerNumber(request, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        });
    });
}

export function trackBillNoByServer(number: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<TrackingByBillNumberResponse> {
    return new Promise<TrackingByBillNumberResponse>(((resolve, reject) => {
        const request = new Request()
        request.setNumber(number)
        request.setScac(TrackingServiceConverter.convertScacIntoEnum(scac))
        request.setCountry(TrackingServiceConverter.convertCountryTypeIntoEnum(country))
        trackingByBillNumberClient.trackByBillNumber(request, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        })
    }))
}