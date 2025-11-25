let num = 123;

let one = (num % 100) % 10;
let two = Math.floor((num % 100) / 10);
let three = Math.floor(num / 100);

console.log("Трёхзначное число: ", num);
console.log("После инверсии: ", one, two, three);
