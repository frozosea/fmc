interface _GetBillNumberListSchema {
    maxRows: number
    models: []
    copNo: string
    bkgNo: string
    cntrNo: string
    enblFlag: string
    hashColoums: any[]
    hashFields: any[]
}

interface _BaseResp {
    TRANS_RESULT_KEY: string
    Exception: string
    count: string
}

export interface OneyGetBillNumberResponse extends _BaseResp {
    list: _GetBillNumberListSchema[]
}

interface _BaseInfoAboutMovingUnitStruct {
    maxRows: number,
    models: any[],
    vslCd: string,
    no: string,
    copNo: string,
    eventDt: string,
    vslEngNm: string,
    placeNm: string,
    skdVoyNo: string,
    yardNm: string,
    copDtlSeq: string,
    skdDirCd: string,
    actTpCd: string,
    statusNm: string,
    statusCd: string,
    nodCd: string,
    vvd: string,
    lloydNo: string,
    hashColumns: any[]
    hashFields: any[]
}

export interface OneyInfoAboutMovingSchema extends _BaseResp {
    list: _BaseInfoAboutMovingUnitStruct[]
}

interface _ContainerSizeUnitSchema {
    maxRows: number
    models: object
    weight: string
    copNo: string
    blNo: string
    eventDt: string
    cntrTpszCd: string
    copStsCd: string
    piece: string
    sealNo: string
    placeNm: string
    yardNm: string
    statusNm: string
    bkgNo: string
    poNo: string
    yardCd: string
    statusCd: string
    cntrNo: string
    cntrTpszNm: string
    dspbkgNo: string
    socFlg: string
    mvmtStsCd: string
    hashFiels: any[]
    hashFields: object
}

export interface OneyGetContainerSizeSchema extends _BaseResp {
    list: _ContainerSizeUnitSchema[]
}