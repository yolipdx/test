
// throttle
// 限制一定时间内可以发送的请求次数
// idea就是设置 cool down，其实就是技能有CD呀

function throttle0(f, cd) {
    let inCooldown = false;
    return function wrapped(...args) {
        if (inCooldown) {
            console.log("ignore: ", args);
            return;
        }

        inCooldown = true;
        f.apply(this, args);
        setTimeout(() => {
            inCooldown = false;
        }, cd);
    };
}

function test(...args) {
    console.log(Date.now(), args);
}

const t = throttle0(test, 2000);
t(1);
t(2);
t(3);
t(4);
t(5);
console.log("done");


// q: 需要save this和args么
// 可以：比如如果这cd期间有call这个function，可以保留最后一次的情况，在cd完后call一次
// refer: https://javascript.info/task/throttle

// ex:
// In a browser we can setup a function to run at every mouse movement and get the pointer location as it moves. 
// During an active mouse usage, this function usually runs very frequently, 
// can be something like 100 times per second (every 10 ms). We’d like to update some information on the web-page when the pointer moves.

function throttle0(f, cd) {
    let inCooldown = false;
    let savedThis = null;
    let savedArgs = null;

    return function wrapped(...args) {
        if (inCooldown) {
            console.log("ignore: ", args);
            savedThis = this;
            savedArgs = args;
            return;
        }

        inCooldown = true;
        f.apply(this, args);
        setTimeout(() => {
            inCooldown = false;

            if (savedThis) {
                wrapped.apply(savedThis, savedArgs);
                savedThis = null;
                savedArgs = null;
            }

        }, cd);
    };
}