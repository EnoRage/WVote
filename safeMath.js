function pow8(number) {
    return Number((Number(number) * Math.pow(10, 8)).toFixed(0));
}

function powMinus8(number) {
    return Number((Number(number) * Math.pow(10, -8)).toFixed(8));
}

module.exports.pow8 = pow8;
module.exports.powMinus8 = powMinus8;