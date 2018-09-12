
// basic for loop
let fact = 1;
for (var i = 2; i < 8; i++) {
  fact = fact * i
}
console.log("fact", fact);
// fact 5040

let j = 2;
let res = 1;
let found = false;
while(!found) {
  res = res * j;
  j = j + 1;
  found = (res == fact);
}

console.log("found factorial", res);
// found factorial 5040

while(true) {
  // use `break` to escape the loop
}
