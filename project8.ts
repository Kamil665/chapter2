function dayOfWeek(k: number, d: number): number {
    let n = (d - 1 + k - 1) % 7 + 1;
    return n % 7;
}

//а)
console.log("1 января - понедельник:");
for (let k of [1, 2, 7]) {
    console.log(`k=${k} -> n=${dayOfWeek(k, 1)}`);
}

//б)
console.log("\n1 января - вторник:");
for (let k of [1, 2, 7]) {
    console.log(`k=${k} -> n=${dayOfWeek(k, 2)}`);
}

//в)
console.log("\n1 января - воскресенье (d = 7):");
for (let k of [1, 2, 7]) {
    console.log(`k=${k} -> n=${dayOfWeek(k, 7)}`);
}
