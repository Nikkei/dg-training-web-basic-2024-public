const { greet } = await import("/lib.js");

greet("Everyone", document.querySelector("#content"));
