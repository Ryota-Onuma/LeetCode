var createHelloWorld = function() {
    const helloWorld = "Hello World"
    return function(...args) {
        return helloWorld
    }
};

/**
 * const f = createHelloWorld();
 * f(); // "Hello World"
 */