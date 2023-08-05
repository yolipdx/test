"use strict";

const timers = []
// for (var i = 0; i < 100; i++) {
for (let i = 0; i < 100; i++) {
    timers.push(setTimeout(() => {
        console.log("i:", i);
        if (i !== 100) {
            console.error("something is wrong");
        }
    }, 10000));
}
console.log("waiting...\n\n");