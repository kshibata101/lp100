package main

import (
    "os"
    "bufio"
    "io/ioutil"
    "regexp"
    "fmt"
    "strings"
    "net/http"
    "net/url"
    "encoding/json"
)

func main() {
    texts := []string{}

    r := regexp.MustCompile(`^\|(.*)`)
    rs := regexp.MustCompile(`^\{\{基礎情報`)
    re := regexp.MustCompile(`^\}\}`)
    
    filepath := "jawiki-uk.txt"
    f, _ := os.Open(filepath)
    defer f.Close()

    basicInfo := false
    s := bufio.NewScanner(f)
    for s.Scan() {
        l := s.Text()
        
        if re.Match([]byte(l)) {
            basicInfo = false
        }
        if basicInfo {
            res := r.FindStringSubmatch(l)
            if res == nil {
                texts[len(texts)-1] += l
            } else {
                texts = append(texts, res[1])
            }
        }
        if rs.Match([]byte(l)) {
            basicInfo = true
        }
    }

    rst := regexp.MustCompile(`'{2,}`)
    rmrk := regexp.MustCompile(`\[\[(([^\[]*)\|)?([^\[]*?)\]\]`)
    rref := regexp.MustCompile(`<ref(.*?)</ref>`)
    rref2 := regexp.MustCompile(`<ref(.*?)/>`)
    rhttp := regexp.MustCompile(`\[(.*?)\]`)
    rtpl := regexp.MustCompile(`\{\{([^\}]*\|)?([^\}]*?)\}\}`)
    rstar := regexp.MustCompile(`\*`)
    
    m := make(map[string]string)
    for _, text := range texts {
        res := strings.Split(text, " = ")
        if (len(res) == 2) {
            word := rst.ReplaceAllString(res[1], "")
            word = rmrk.ReplaceAllString(word, "$3")
            word = rref.ReplaceAllString(word, "")
            word = rref2.ReplaceAllString(word, "")
            word = rhttp.ReplaceAllString(word, "$1")
            word = rtpl.ReplaceAllString(word, "$2")
            word = rstar.ReplaceAllString(word, "")
            m[res[0]] = word
        }
    }

    flagPath := m["国旗画像"]
    flagPath = url.PathEscape(flagPath)

    flagUrl := "https://en.wikipedia.org/w/api.php?action=query&format=json&prop=imageinfo&iiprop=url&titles=File:" + flagPath
    
    response, _ := http.Get(flagUrl)
    defer response.Body.Close()

    bytes, _ := ioutil.ReadAll(response.Body)
    
    var resp map[string]interface{}
    err := json.Unmarshal(bytes, &resp)
    if err != nil {
        fmt.Println(err)
        fmt.Println(string(bytes))
        return
    }

    imageUrl := ""
    pages := resp["query"].(map[string]interface{})["pages"].(map[string]interface{})
    for _, v := range pages {
        imageUrl = v.(map[string]interface{})["imageinfo"].([]interface{})[0].(map[string]interface{})["url"].(string)
    }
    fmt.Println(imageUrl)
}
