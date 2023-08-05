console.log("main in");
let i = 0;

let timerId = setTimeout(function tick() {
    console.log(`tick ${i}`);
    if (i == 5) {
        return;
    }

    i ++;
    timerId = setTimeout(tick, 2000);
}, 2000);




console.log("main out");