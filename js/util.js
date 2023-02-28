function getRandomPositiveInteger(a, b) {
    const lower = Math.ceil(Math.min(Math.abs(a), Math.abs(b)));
    const upper = Math.floor(Math.max(Math.abs(a), Math.abs(b)));
    const result = Math.random() * (upper - lower + 1) + lower;
    return Math.floor(result);
}

function checkStringLength(string, length) {
    return string.length <= length;
}

function createIdGenerator() {
    let lastGeneratedId = 0
    return function () {
        lastGeneratedId++;
        return lastGeneratedId;
    }
}

function createIdGeneratorFromRange(min, max) {
    const previousValues = [];
    return function () {
        let currentValue = getRandomPositiveInteger(min, max);
        if (previousValues.length >= (max - min + 1)) {
            console.error('Перебраны все числа из диапазона от ' + min + ' до ' + max);
            return null;
        }
        while (previousValues.includes(currentValue)) {
            currentValue = getRandomPositiveInteger(min, max);
        }
        previousValues.push(currentValue);
        return currentValue;
    }
}

export { getRandomPositiveInteger, createIdGenerator }