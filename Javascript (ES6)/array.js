//=====================================
// Array
//=====================================
let A = [1,6,3,9];

// for...of
//=====================================
for (let v of A) {
  console.log(v); // print value
}

// for...in
//=====================================
for (let v in A) {
  console.log(v); // print index/key
}

A[30] = 30;
console.log(A.length); // 31, last index + 1 and not size of array
console.log(A[10]); // undefined

// ForEach iterator
//=====================================
A.forEach((element, index, array) => {
  element = 1; // Pass-by-value for primative
  // array[index] = 1; // use this to update the value
})
console.log(A); // [ 1, 6, 3, 9, <26 empty items>, 30 ]

let B = [{key:1}, {key:2}];
B.forEach((element, index, array) => {
  element.key = 3; // Pass-by-reference for object
})
console.log(B); // [ { key: 3 }, { key: 3 } ]

// Add/remove (singular)
//=====================================
// Modify tail
A.push(20); // [ 1, 6, 3, 9, <26 empty items>, 30, 20 ], add to end
console.log(A);
A.pop(); // [ 1, 6, 3, 9, <26 empty items>, 30], remove from end
console.log(A);

// Modify head
A.unshift(20); // [ 20, 1, 6, 3, 9, <26 empty items>, 30 ]
console.log(A);
A.shift();
console.log(A); // [ 1, 6, 3, 9, <26 empty items>, 30 ]
// For stack, use push and pop
// For queue, use push and shift

// Add/remove (mass)
//=====================================
A.splice(2, 2, 7, 8, 9); // (start index, delete count, ...element(s) to add), method returns deleted elements in an Array
console.log(A); // [ 1, 6, 7, 8, 9, <26 empty items>, 30 ], in-place
A.splice(3); // Delete all elements after start index
console.log(A); // [1, 6, 7]

// Copy array (shallow-copy)
//=====================================
console.log(A.slice(1)); // Copy whole array from start index
console.log(A.slice(1,2)); // Range copy from start to end (non-inclusive)

// Find element
//=====================================
console.log(A.indexOf(6)); // 1
console.log(A.indexOf(10)); // -1

