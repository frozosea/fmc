import {Server, ServerCredentials} from "@grpc/grpc-js";
import {TrackingByContainerNumberService} from "./proto/server_grpc_pb";
import {grpcService} from "../containers";

export const server = new Server();
// @ts-ignore
server.addService(TrackingByContainerNumberService, grpcService)

export default function startServer() {
    server.bindAsync(`localhost:${process.env.GRPC_PORT}`, ServerCredentials.createInsecure(), (error, port) => {
        console.log("SERVER WAS STARTED")
        server.start()
    })
}

