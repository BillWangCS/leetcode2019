/*
A query word matches a given pattern if we can insert lowercase letters to the pattern word so that it equals the query.
(We may insert each character at any position, and may insert 0 characters.)

Given a list of queries, and a pattern, return an answer list of booleans,
where answer[i] is true if and only if queries[i] matches the pattern.



Example 1:

Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FB"
Output: [true,false,true,true,false]
Explanation:
"FooBar" can be generated like this "F" + "oo" + "B" + "ar".
"FootBall" can be generated like this "F" + "oot" + "B" + "all".
"FrameBuffer" can be generated like this "F" + "rame" + "B" + "uffer".
Example 2:

Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBa"
Output: [true,false,true,false,false]
Explanation:
"FooBar" can be generated like this "Fo" + "o" + "Ba" + "r".
"FootBall" can be generated like this "Fo" + "ot" + "Ba" + "ll".
Example 3:

Input: queries = ["FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"], pattern = "FoBaT"
Output: [false,true,false,false,false]
Explanation:
"FooBarTest" can be generated like this "Fo" + "o" + "Ba" + "r" + "T" + "est".


Note:

1 <= queries.length <= 100
1 <= queries[i].length <= 100
1 <= pattern.length <= 100
All strings consists only of lower and upper case English letters.
 */

package main

import (
	"fmt"
)


//这段儿想的有点儿复杂，可以用一维的map记录最大的q中的某字母的下标就可以了，主要是已经写了，犯懒没做优化
func canConstruct(q, p string) bool {
	dictQ := make(map[string]map[int]int)
	cntQ := make(map[string]int)
	for i, v := range q {
		if dictQ[string(v)] == nil {
			indexMap := make(map[int]int)
			indexMap[cntQ[string(v)]] = i
			dictQ[string(v)] = indexMap
		} else {
			dictQ[string(v)][cntQ[string(v)]] = i
		}
		cntQ[string(v)]++
	}
	cntP := make(map[string]int)
	fmt.Println(dictQ, cntP)
	last := -1
	for _, v := range p {
		if dictQ[string(v)] == nil {
			return false
		}
		flag := false
		for _, v := range dictQ[string(v)] {
			if v > last {
				flag = true
			}
		}
		if !flag {
			return false
		}
		last = dictQ[string(v)][cntP[string(v)]]
		cntP[string(v)]++
	}
	return true
}

func isCamelMatched(q string, p string) bool {
	upperQ := ""
	upperP := ""
	strsQ := make([]string, 0)
	strsP := make([]string, 0)
	lastq := 0
	lastp := 0
	for i, v := range q {
		if !(byte(v) >= 'a' && byte(v) <= 'z') {
			upperQ += string(v)
			strsQ = append(strsQ, q[lastq:i])
			lastq = i
		}
	}
	strsQ = append(strsQ, q[lastq:])
	for i, v := range p {
		if !(byte(v) >= 'a' && byte(v) <= 'z') {
			upperP += string(v)
			strsP = append(strsP, p[lastp:i])
			lastp = i
		}
	}
	strsP = append(strsP, p[lastp:])
	if upperP != upperQ {
		return false
	}
	fmt.Println(strsQ, strsP)
	for i := 0; i < len(strsP); i++ {
		if !canConstruct(strsQ[i], strsP[i]) {
			return false
		}
	}
	return true
}

func camelMatch(queries []string, pattern string) []bool {
	res := make([]bool, 0)
	for _, v := range queries {
		res = append(res, isCamelMatched(v, pattern))
	}
	return res
}

func main() {
	fmt.Println(camelMatch([]string{"FooBar","FooBarTest","FootBall","FrameBuffer","ForceFeedBack"}, "FoBa"))
	//fmt.Println(canConstruct("nfkznkhfuIuiixcccNip", "nznkhfuIixccNip"))
}
/*
["nsrzmnskhfuIiyxccNip","nfkznkhfuIuiixcccNip","naznkghfuewIixccNmip","nznkhfuqIixdcicNiygp","nznukhifuIixccNifbwp","nznkhfuIsixocjcNidmp"]
"nznkhfuIixccNip"
["CompetitiveProgramming","CounterPick","ControlPanel"]
"CooP"
 */