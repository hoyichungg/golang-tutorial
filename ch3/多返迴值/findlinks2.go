package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

// 循環遍歷命令列參數。 這意味著你可以執行程式並提供多個URL，如：./program http://example.com http://anotherexample.com。
// 對於每一個URL，findLinks函數被呼叫。
// 如果回傳錯誤，它將錯誤列印到標準錯誤輸出並繼續下一個URL。
// 對於每一個找到的鏈接，它將其列印到標準輸出。

func main() {
	for _, url := range os.Args[1:] {
		links, err := findLinks(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "findlinks2: %v\n", err)
			continue
		}
		for _, link := range links {
			fmt.Println(link)
		}
	}
}

// 使用http.Get函數嘗試取得指定URL的內容。
// 如果HTTP回應的狀態碼不是200（即OK），則關閉回應體並傳回錯誤。
// 使用html.Parse函數嘗試解析響應體中的HTML內容。
// 關閉響應體。
// 使用visit函數遍歷解析後的HTML並收集所有的連結。
// 返回找到的連結。

// 調用多返迴值函數時，返迴給調用者的是一組值，調用者必鬚顯式的將這些值分配給變量:
// links, err := findLinks(url)

// 如果某個值不被使用，可以將其分配給blank identifier:
// links, _ := findLinks(url) // errors ignored

func findLinks(url string) ([]string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, resp.Status)
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML: %v", url, err)
	}
	return visit(nil, doc), nil
}

// 這是一個遞歸函數，用來遍歷HTML的節點樹。
// 如果目前節點是一個<a>標籤（即一個連結），它將尋找其href屬性並將該屬性的值加入連結清單。
// 無論當前節點是否是一個鏈接，它都會遞歸地檢查其所有子節點。
// 返回收集的連結列表。

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}

	// 	初始化 (c := n.FirstChild): 在迴圈的開始，我們初始化一個變數c為節點n的第一個子節點。 這基本上是說：「開始於n的第一個子節點」。

	// 條件 (c != nil): 這是一個繼續循環的條件。 只要c不為nil（即，只要還有要處理的節點），循環就會繼續。 一旦c成為nil（意味著我們已經處理完所有的兄弟節點），循環就會停止。

	// 後置操作 (c = c.NextSibling): 在每次循環迭代結束後，我們將c設定為它的下一個兄弟節點。 這樣，我們能夠遍歷所有的子節點。

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}
