//=====================================
// Object
//=====================================

// Creation of Object
//=====================================
var obj = new Object();
obj = {};

// Attribute access
//=====================================
obj = {
  name: 'Carrot',
  details: {
    'color': 'orange', // key can be written as String, though not recommended
    size: 12
  }
}

console.log(obj.details.color); // orange
console.log(obj['details']['color']); // orange
let key = 'color';
console.log(obj['details'][key]); // oranage, variable "key" resolved at runtime to 'color'
console.log(obj.details.key); // undefined, does not work for dot operator, attribute key does not exist

// Iterate through properties in object
//=====================================
for (let k in obj) { // cannot use "of"
  console.log(k); // key of object
  console.log(obj[k]); // accessed value of object, cannot use "obj.k"
}

// Runtime creation
//=====================================
let type = 'cat1';
obj = {type : 123}; // {type: 123}
obj = {[type] : 123}; // {cat1: 123}, dynamic key

// Object access via [] prevents optimizations of some js engine, use dot operator if possible.


//=====================================
// Prototype/Classes
//=====================================

// Creation of object prototype
//=====================================
function Person(name, age) {
  this.name = name;
  this.age = age;
  this.nameAge = function() {
    return this.name + ' ' + this.age;
  }
}

// Using ES2015 class syntactic sugaring
class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  // method
  nameAge() {
    return this.name + ' ' + this.age;
  }
}


// Use of new and this declaration, hence Person cannot be created using a arrow function
let p1 = new Person('James', 12);
console.log(p1.nameAge()); // James 12

// To add new method/function to prototype
Person.prototype.ageName = function() {
  return this.age + ' ' + this.name;
}

console.log(p1.ageName()); // 12 James
let fn = p1.ageName;
console.log(fn());

// Object Creation (bad way)
//=====================================
// This is similar to that of this
function Person2(name, age) {
  let obj = {};
  obj.name = name;
  obj.age = age;
  obj.nameAge = function() {
    return obj.name + ' ' + obj.age;
  }
}

function makePerson(name, age) {
  return {
    name: name,
    age: age,
    nameAge: function() {
      return this.name + ' ' + this.age; // 'this' keyword is not supported in => function
    }
  }
}

let p2 = makePerson('James', 20);
console.log(p2.nameAge()); // "James 20"
let getFn = p2.nameAge;
console.log(getFn()); // undefined undefined, because it is running "this.name and this.age" here i.e. this is not binded to obj. In this scope, name and age is not defined.


// 'this' keyword and binding
//=====================================
function print(a, b) {
  return this.name + ' ' + a + ' ' + b;
}

// this access global scope and name is not defined
// Because 'this' is evaluated at runtime and is not dependent on where the method is declared,
// but rather on what object is "before the dot operator".
// For 'this' to be evaluated at point of declaration, use the arrow function.
console.log(print(1,2)); // undefined 1 2

// Pass 'this' value to function
// apply takes in an array [] to be passed to the function as parameters, while call takes in a rest arguments.
console.log(print.apply({name : 'anny'}, [1,2])); // anny 1 2
console.log(print.call({name: 'anny'}, 1, 2)); // anny 1 2

// Bind 'this' value to function
let bindedPrint = print.bind({name: 'anny'});
console.log(bindedPrint(1,2)); // anny 1 2

// Arrow function bind automatically to its local scope at creation, cannot change using bind, apply and call.




