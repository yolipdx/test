let p = Promise.reject(new Error("promise failed"));
setTimeout(() => {
    p.catch(e => console.error("caught: ", e))
}, 1000);

globalThis.addEventListener("unhandledrejection", e => console.error("unhandledrejection: ", e.reason));