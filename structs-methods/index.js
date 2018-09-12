class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  setName(name) {
    this.name = name;
  }

  getName() {
    return this.name;
  }
}

const p = new Person('Geraldine', 23);
console.log(p.getName());
p.setName('Gerald');
console.log(p.getName());
