package akkn

import "github.com/PuerkitoBio/goquery"

func checkNumberBelongsLine(doc *goquery.Document) bool {
	elem := doc.Find("#GridView1 > tbody > tr:nth-child(2) > td:nth-child(3))")
	return elem != nil
}
