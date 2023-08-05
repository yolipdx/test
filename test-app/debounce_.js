// f如果连续被call，只处理一次。比如最后一次f后冷却个ms
// 比如在input 框里输入，当停止输入1秒后我们认为输入完成，于是开始处理

function debounce(f, ms) {
    let timerId = null;

    return function wrapper(...args) {
        if (timerId) {
            clearTimeout(timerId);
            timerId = null;
        }

        timerId = setTimeout(() => {
            // ms后会执行我们的function，如果在timeout发生之前没有被cancel的话
            f.apply(this, args);
        }, ms);
    }
}

// refer: https://javascript.info/task/debounce