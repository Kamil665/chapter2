let num = 487;

let n1 = Math.floor(num / 100);
let n2 = Math.floor((num % 100) / 10);
let n3 = Math.floor((num % 100) % 10);

console.log("Трёхзначное число: ", num);
console.log("После перестановки: ", n1, n3, n2);
