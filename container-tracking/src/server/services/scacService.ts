import {ILogger} from "../../logging";
import {GetAllScacResponse, Request, Scac} from "../proto/tracking_pb";
import {IScacServiceServer} from "../proto/tracking_grpc_pb";
import {sendUnaryData, ServerUnaryCall} from "@grpc/grpc-js";

interface BaseScac {
    scac: string,
    fullName: string
}

export interface IScacRepository {
    GetAll(): Promise<BaseScac[]>
}


export class ScacForTrackingRepository implements IScacRepository {
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

    public async GetAllScac(): Promise<BaseScac[]> {
        try {
            return this.repository.GetAll();
        } catch (e) {
            this.logger.ExceptionLog(`get all scac from repository exception: ${String(e)}`)
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

    public getAll(call: ServerUnaryCall<Request, null>, callback: sendUnaryData<GetAllScacResponse>) {
        this.service.GetAllScac().then((response: BaseScac[]) => {
                return callback(null, this.converter.toGrpc(response))
            }
        ).catch((e) => {
                return callback(e, null);
            }
        )
    }
}