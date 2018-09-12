function split(str) {
	const firstHalf = str.substring(0, str.length/2);
	const secondHalf = str.substring(str.length/2, string.length);
	return [firstHalf, secondHalf];
}

const [s1, s2] = split("YELLOW_SUBMARINE");
console.log(s1);
// =>  YELLOW_S
console.log(s2);
// =>  UBMARINE
