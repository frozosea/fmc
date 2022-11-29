import {Server, ServerCredentials} from "@grpc/grpc-js";
import {
    ScacServiceService,
    TrackingByBillNumberService,
    TrackingByContainerNumberService
} from "../../../protobuf/tracking/tracking_grpc_pb";
import {scacGrpcService, trackingByBillNumberService, trackingByContainerNumberService} from "../init";
import {config} from "dotenv";
import {AppDataSource} from "../db/data-source";

export const server = new Server();
// @ts-ignore
server.addService(TrackingByContainerNumberService, trackingByContainerNumberService)
// @ts-ignore
server.addService(TrackingByBillNumberService, trackingByBillNumberService)
// @ts-ignore
server.addService(ScacServiceService, scacGrpcService);

export default function startServer() {
    config()
    AppDataSource.initialize()
        .then(async (source) => {
        })
        .catch((error) => console.log("Error of init db: ", error))
    server.bindAsync(`0.0.0.0:${process.env.GRPC_PORT}`, ServerCredentials.createInsecure(), (error, port) => {
        server.start();
        if (error) console.log(error)
        console.log(`SERVER WAS STARTED ON ${port}`)
    })
}

