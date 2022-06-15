export default class RequestsUtils {
    static jsonToQueryString(json: object): string {
        let arr: string[] = []
        for (let item in json) {
            arr.push(`${item}=${json[item]}`)
        }
        return arr.join("&")
    }
}