const b = false;

if (b) {
  console.log('We need parens');
} else {
  console.log('Not much to see here.');
}

const str = "";

if (str) {
  console.log('Implicit type conversions');
} else {
  console.log('Empty string is false-y');
}
