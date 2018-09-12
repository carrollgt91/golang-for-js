const values = ["a", "b", "c", "d", "e"];

values.push("z");

values.forEach((val, index) => {
  console.log(`index ${index}: ${val}`);
});
