import {SCAC_TYPE} from "../../types";
import {AppDataSource} from "../../db/data-source";
import {Repository} from "typeorm";
import {ContainerScac} from "../../db/entity/containerScac";

export interface IScacContainers {
    addContainer(container: string, scac: SCAC_TYPE): Promise<void>;

    getScac(container: string): Promise<SCAC_TYPE | null>

}

export class ScacRepository implements IScacContainers {
    public repo: Repository<ContainerScac>

    public constructor() {
        this.repo = AppDataSource.getRepository(ContainerScac)
    }

    public async addContainer(container: string, scac: SCAC_TYPE): Promise<void> {
        let object: ContainerScac = new ContainerScac()
        object.container = container
        object.scac = scac
        try{
            await this.repo.save(object)
        }catch (e) {
            console.log(e)
            return
        }
    }

    public async getScac(container: string): Promise<SCAC_TYPE | null> {
        let result: ContainerScac | null = await this.repo.findOneBy({
            container: container
        })
        if (result !== null) return result.scac
        else return null
    }
}