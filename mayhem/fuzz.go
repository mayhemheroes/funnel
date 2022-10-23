package fuzz

import "strconv"
import "github.com/ohsu-comp-bio/funnel/tes"
import "github.com/ohsu-comp-bio/funnel/config"

func mayhemit(bytes []byte) int {

    var num int
    if len(bytes) < 1 {
        num = 0
    } else {
        num, _ = strconv.Atoi(string(bytes[0]))
    }

    switch num {
    
    case 0:
        content := string(bytes)
        tes.NewClient(content)
        return 0

    case 1:
        var test config.Config
        config.Parse(bytes, &test)
        return 0

    default:
        var test config.Config
        content := string(bytes)
        config.ParseFile(content, &test)
        return 0
    }
}

func Fuzz(data []byte) int {
    _ = mayhemit(data)
    return 0
}