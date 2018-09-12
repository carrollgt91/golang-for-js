class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  int() {
    return this.age;
  }
}

class MyFloat {
  constructor(fl) {
    this.fl = fl;
  }

  int() {
    return Math.floor(this.fl)
  }
}

const p = new Person("Gerald", 25);
console.log(p.int())

const f = new MyFloat(25.5)
console.log(f.int());

const i = 20;
console.log(i.int()) // runtime error
