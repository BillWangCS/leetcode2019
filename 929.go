/*Every email consists of a local name and a domain name, separated by the @ sign.

For example, in alice@leetcode.com, alice is the local name, and leetcode.com is the domain name.

Besides lowercase letters, these emails may contain '.'s or '+'s.

If you add periods ('.') between some characters in the local name part of an email address,
mail sent there will be forwarded to the same address without dots in the local name.
For example, "alice.z@leetcode.com" and "alicez@leetcode.com" forward to the same email address.
(Note that this rule does not apply for domain names.)

If you add a plus ('+') in the local name, everything after the first plus sign will be ignored.
This allows certain emails to be filtered, for example m.y+name@email.com will be forwarded to my@email.com.
(Again, this rule does not apply for domain names.)

It is possible to use both of these rules at the same time.

Given a list of emails, we send one email to each address in the list.
How many different addresses actually receive mails?



Example 1:

Input: ["test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"]
Output: 2
Explanation: "testemail@leetcode.com" and "testemail@lee.tcode.com" actually receive mails


Note:

1 <= emails[i].length <= 100
1 <= emails.length <= 100
Each emails[i] contains exactly one '@' character.*/

package main

import (
	"fmt"
	"strings"
)

func trimEmail(email string) string {
	tail := strings.Split(email, "@")[1]
	s := strings.Split(email, "+")[0]
	head := strings.Replace(s, ".", "", -1)
	return head+ "@" + tail
}

func numUniqueEmails(emails []string) int {
	emailMap := make(map[string]bool)
	for _, email := range emails {
		emailMap[trimEmail(email)] = true
	}
	return len(emailMap)
}

func main() {
	fmt.Println(trimEmail("test.email+alex@leetcode.com"))
	a := []string{"test.email+alex@leetcode.com","test.e.mail+bob.cathy@leetcode.com","testemail+david@lee.tcode.com"}
	fmt.Println(numUniqueEmails(a))
}