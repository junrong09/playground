//=====================================
// Function
//=====================================

// Basic function/arguments
//=====================================
let add1 = (x,y) => {
  return x + y;
}

console.log(add1(1, 2)); // 3
console.log(add1()) // NaN, undefined + undefined were converted to Number (NaN)

// Access arguments (without explicit declaration)
// Does not work for arrow => function declaration
//=====================================
function add2() {
  let x = 0;
  console.log(arguments); // { '0': 1, '1': 2 }, an object that is array-like
  for (let i = 0; i < arguments.length; i++) {
    x += arguments[i]; 
  }
  return x;
}

console.log(add2(1,2));


// Rest parameter operator
//=====================================
let add3 = (init, ...args) => {
  let x = init;
  console.log(args); // [2,3], an Array
  for (let v of args) {
    x += v;
  }
  return x;
}

console.log(add3(1,2,3));

