export function greet(name = "World", target) {
  target.appendChild(document.createTextNode(`Hello, ${name}!`));
}
