//=====================================
// Spread operator
//=====================================
let A = ["a","b"]
let B = ["c", ...A] 
let C = ["c", A] 

console.log(B); // ["c", "a", "b"]
console.log(C); //["c", ["a", "b"]]

if (!undefined) {
  console.log("print");
}


//=====================================
// Number
//=====================================
// Conversion from String to Number
let n1 = parseInt('102', 10); // base 10, 2nd parameter
let n2 = + '102';
let n3 = parseFloat('10.12', 10);
console.log(n1); // 102
console.log(n2) // 102
console.log(n3) // 10.12

n1 = parseInt('102abc2', 10);
n2 = + '102abc2';
console.log(n1); // 102
console.log(n2); // NaN


// Not a number (a special type of number)
let bool = isNaN(NaN + 2);
console.log(bool) //true

// Inf and -inf
let inf = 1/0;
let nInf = -1/0;
console.log(inf); // Infinity
console.log(nInf); // -Infinity
console.log(isFinite(inf)); // false


//=====================================
// String
//=====================================
// Operations
let len = 'hello'.length;
let c1 = 'hello'.charAt(1);
let c2 = 'hello, world'.replace('world', 'panda');
let c3 = 'hello'.toUpperCase();
console.log(`${len} ${c1} ${c2} ${c3}`); // 5 e hello, panda HELLO


//=====================================
// Boolean
//=====================================
console.log(Boolean("" || NaN || undefined || null || 0)); // false
console.log(Boolean([] && {} && 1)); // true


//=====================================
// Block-scoped variables
//=====================================
if (true) {
  // start of block
  let v1 = 10;
  const v2 = 20;
  var v3 = 30;
  // end of block
}

// console.log(v1); // Error as v1 is not defined
// console.log(v2); // Error as v2 is not defined
console.log(v3); // 30


//=====================================
// Operator
//=====================================
// Number operators
console.log(-10 % 9); // -1, negative modulo
let x = 1;
x += 1;
x++;
console.log(x); // 3

// String operators
console.log('hello' + 'world');
console.log('3' + 4 + 5); // "345", number after String is converted to String
console.log(3 + 4 + '5' + 6); // "756", (3 + 4) is evaluated as number before converting to String

// Comparison operators
console.log(123 == '123'); // true, "type cocercion" converts to common type and match
console.log(10 >= '1') // true
console.log(undefined == false); // false, cannot convert undefined/NaN to boolean

console.log(123 === '123'); // false, checked both type and value
// console.log(10 >== '1'); // Error, all relational (inequality) operators perform type-conversion

let o1 = {key: "123"};
let o2 = {key: "123"};
console.log(o1 == o2); // false

// Type conversion rules:
// String/Boolean == Number - Convert to Number (For boolean true -> 1, false -> +0)
// Object == Number/String - Convert to Number/String respectively using valueOf/toString
// Object == Object - True, only when point to same object


//=====================================
// control Structures
//=====================================
// Ternary
let age = 20;
let group = (age > 18) ? 'old' : 'young';
console.log(group); // 'old'
