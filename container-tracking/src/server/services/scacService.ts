import {ILogger} from "../../logging";
import {GetAllScacResponse, Request, Scac} from "../../../../protobuf/tracking/tracking_pb";
import {IScacServiceServer} from "../../../../protobuf/tracking/tracking_grpc_pb";
import {sendUnaryData, ServerUnaryCall} from "@grpc/grpc-js";

interface BaseScac {
    scac: string,
    fullName: string
}

export interface IScacRepository {
    GetAll(): Promise<BaseScac[]>

    getContainerScac(): Promise<BaseScac[]>

    getBillScac(): Promise<BaseScac[]>
}


export class ScacForTrackingRepository implements IScacRepository {
    public async getContainerScac(): Promise<BaseScac[]> {
        return [
            {
                scac: "COSU",
                fullName: "Cosco"
            },
            {
                scac: "MAEU",
                fullName: "Maersk/Sealand Maersk"
            },
            {
                scac: "MSCU",
                fullName: "MSC"
            },
            {
                scac: "ONEY",
                fullName: "One Line"
            },
            {
                scac: "FESO",
                fullName: "Fesco"
            },
            {
                scac: "SKLU",
                fullName: "Sinokor"
            },
            {
                scac: "HALU",
                fullName: "Heung-a"
            },
            {
                scac: "SITC",
                fullName: "SITC"
            }
        ]
    }

    public async getBillScac(): Promise<BaseScac[]> {
        return [
            {
                scac: "FESO",
                fullName: "Fesco"
            },
            {
                scac: "SKLU",
                fullName: "Sinokor"
            },
            {
                scac: "HALU",
                fullName: "Heung-a"
            },
            {
                scac: "SITC",
                fullName: "SITC"
            }
        ]
    }

    public async GetAll(): Promise<BaseScac[]> {
        return [
            {
                scac: "FESO",
                fullName: "Fesco"
            },
            {
                scac: "SKLU",
                fullName: "Sinokor"
            },
            {
                scac: "MAEU",
                fullName: "Maersk/Sealand"
            },
            {
                scac: "COSU",
                fullName: "COSCO"
            },
            {
                scac: "ONEY",
                fullName: "ONE"
            },
            {
                scac: "SITC",
                fullName: "SITC"
            },
            {
                scac: "MSCU",
                fullName: "MSC"
            },
            {
                scac: "HALU",
                fullName: "Heung-a"
            },
            {
                scac: "ZHGU",
                fullName: "Zhong gu"
            },
        ]
    }
}

export class ScacService {
    protected repository: IScacRepository;
    protected logger: ILogger;

    public constructor(repository: IScacRepository, logger: ILogger) {
        this.repository = repository;
        this.logger = logger;
    }

    public async getContainerScac(): Promise<BaseScac[]> {
        try {
            return this.repository.getContainerScac();
        } catch (e) {
            this.logger.ExceptionLog(`get container scac from repository exception: ${String(e)}`)
            throw e;
        }
    }

    public async getBillScac(): Promise<BaseScac[]> {
        try {
            return this.repository.getBillScac();
        } catch (e) {
            this.logger.ExceptionLog(`get bill scac from repository exception: ${String(e)}`)
            throw e;
        }
    }
}

class Converter {
    public oneScacToGrpc(scac: BaseScac): Scac {
        const s = new Scac();
        s.setScac(scac.scac);
        s.setFullname(scac.fullName)
        return s;
    }

    public toGrpc(scac: BaseScac[]): GetAllScacResponse {
        const grpcScac = new GetAllScacResponse();
        for (let index = 0; index < scac.length; index++) {
            grpcScac.addAllScac(this.oneScacToGrpc(scac[index]), index)
        }
        return grpcScac;
    }
}


export class ScacGrpcService implements IScacServiceServer {
    protected service: ScacService;
    protected converter: Converter;

    public constructor(service: ScacService) {
        this.service = service;
        this.converter = new Converter();
    }

    public getContainerScac(call: ServerUnaryCall<Request, null>, callback: sendUnaryData<GetAllScacResponse>) {
        this.service.getContainerScac().then((response: BaseScac[]) => {
                return callback(null, this.converter.toGrpc(response))
            }
        ).catch((e) => {
                return callback(e, null);
            }
        )
    }

    public getBillScac(call: ServerUnaryCall<Request, null>, callback: sendUnaryData<GetAllScacResponse>) {
        this.service.getBillScac().then((response: BaseScac[]) => {
                return callback(null, this.converter.toGrpc(response))
            }
        ).catch((e) => {
                return callback(e, null);
            }
        )
    }
}