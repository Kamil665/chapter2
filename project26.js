let n = 564
let op1 = Math.floor((n / 100) % 10)
let op2 = Math.floor(n / 10)
let op3 = Math.floor(op1 * 100)
let x = Math.floor(op3 + op2) - 100

console.log("n = ", n)
console.log("x = ", x)
