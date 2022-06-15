import {Unlocode} from "../../../db/entity/Unlocode";
import {AppDataSource} from "../../../db/data-source";
import {Repository} from "typeorm";

export interface UnlocodeObject {
    unlocode: string,
    fullname: string
}

export interface IUnlocodesRepo {
    addUnlocode(obj: UnlocodeObject): Promise<void>;

    getUnlocode(unlocode): Promise<string>;
}

export class UnlocodesRepo implements IUnlocodesRepo {
    private repo: Repository<Unlocode>

    public constructor() {
        this.repo = AppDataSource.getRepository(Unlocode)
    }

    public async addUnlocode(obj: UnlocodeObject): Promise<void> {
        let unlo: Unlocode = new Unlocode()
        unlo.unlocode = obj.unlocode
        unlo.fullname = obj.fullname
        await this.repo.save(unlo)
    }

    public async getUnlocode(unlocode: string): Promise<string> {
        let result: Unlocode | null = await this.repo.findOneBy({
            unlocode: unlocode
        })
        if (result !== null) {
            return result.fullname
        }
        return unlocode
    }
}
