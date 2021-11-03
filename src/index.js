import './styles/scss/main.scss';

const obj = {
    a: "alpha",
    b: "bravo",
    d: "delta"
}

const objNew = {
    ...obj,
}

console.log('Hello')

console.log(obj, objNew)