/* 
* Two-for-one is a way of saying that if you buy one, you also get one for free.
* So the phrase "two-fer" often implies a two-for-one offer.
*/
package twofer

/*
* ShareWith generates a phrase "One for {name}, one for me."
* If no name is provided, it defaults to "you".
*/
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
