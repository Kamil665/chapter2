let num = 234;

let n1 = Math.floor(num / 100);
let n2 = Math.floor((num % 100) / 10);
let n3 = (num % 100) % 10;

console.log("Трёхзначное число: ", num);
console.log("После исправления: ", n2, n3, n1);
