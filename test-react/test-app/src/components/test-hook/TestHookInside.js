const hooks = [];
let hookIndex = 0;

function prepairHook() {
    hookIndex = 0;
    console.log('hook prepair done with index: ', hookIndex);
}

function updateHookIndex() {
    hookIndex ++;
}

function useState(initialState) {
    let statePair = hooks[hookIndex];
    if (!statePair) {
        hooks[hookIndex] = [initialState, setState];
    }

    function setState(nextState) {
        hooks[hookIndex][0] = nextState;
        updateDOM();
    };

    updateHookIndex();
    return hooks[hookIndex - 1];
}



function updateDOM() {

}