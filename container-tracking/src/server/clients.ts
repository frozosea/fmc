import {TrackingByBillNumberClient, TrackingByContainerNumberClient} from "./proto/tracking_grpc_pb";
import {credentials} from "@grpc/grpc-js";
import {Request, TrackingByBillNumberResponse, TrackingByContainerNumberResponse} from "./proto/tracking_pb";
import {COUNTRY_TYPE, SCAC_TYPE} from "../types";
import {TrackingServiceConverter} from "./services/trackingByContainerNumberService";

export const trackingByContainerNumberClient = new TrackingByContainerNumberClient(
    `localhost:${process.env.GRPC_PORT}`,
    credentials.createInsecure(),
);
export const trackingByBillNumberClient = new TrackingByBillNumberClient(`localhost:${process.env.GRPC_PORT}`,
    credentials.createInsecure(),)

export function trackContainerByServer(container: string, scac: SCAC_TYPE, country: COUNTRY_TYPE): Promise<TrackingByContainerNumberResponse> {
    return new Promise<TrackingByContainerNumberResponse>((resolve, reject) => {
        const request = new Request();
        request.setNumber(container);
        request.setScac(scac)
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
        request.setScac(scac)
        request.setCountry(TrackingServiceConverter.convertCountryTypeIntoEnum(country))
        trackingByBillNumberClient.trackByBillNumber(request, (err, result) => {
            if (err) reject(err);
            else resolve(result);
        })
    }))
}