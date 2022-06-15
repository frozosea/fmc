const fs = require("fs")

const jsdom = require("jsdom");
const {JSDOM} = jsdom;

function getTableJsonAndDocInstance(stringHtml) {
    let doc = new JSDOM(stringHtml).window.document
    let table = doc.querySelector("#paging_1")
    console.log(table)
    // const [header] = table.tHead.rows
    // const props = [...header.cells].map(h => h.textContent)
    const rows = [...table.rows].map(r => {
        const entries = [...r.cells].map((c, i) => {
            return [Number(i), c.textContent.replace(/(\r\n|\n|\r|\t)/gm, "")]
        })
        return Object.fromEntries(entries)
    })
    return [rows, doc]
}

function reg(htmlstr) {
    let re = new RegExp(/'\d','\w+','\w.+'/gm)
    // htmlstr.matchAll(/'\d','\w.+'/gm)
    return htmlstr.match(re)
}

new fs.readFile("/Users/frozo/PycharmProjects/findmycargofastapi/container_tracking/kmtc.html", (err, data) => {
    let info = reg(String(data))
    let dupl = info
    for (let i in dupl) {
        let element = dupl[i].split(',')
        for (let item of element) {
            element[element.indexOf(item)] = item.replaceAll(/'/g, "")
        }
        dupl[i] = element.join(",")
    }
    console.log(dupl)
})