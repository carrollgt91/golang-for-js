class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }
}

const p = new Person('Geraldine', 23);
console.log(p);
console.log(p.name);

// age gets default value of undefined here
const p2 = new Person("George");
console.log(p2);
// we would need to manually specify default
// values to achieve the same functionality
const p3 = Person("", 50);
console.log(p3);
