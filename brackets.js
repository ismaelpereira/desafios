const data = "[]{}()[][][]";
let result = true;
let brackets = [];
for (let i = 0; i < data.length; i++) {
  if (data[i] === "[" || data[i] === "(" || data[i] === "{") {
    brackets.push(data[i]);
    continue;
  }

  if (
    data[i] === "]" &&
    brackets.length !== 0 &&
    brackets[brackets.length - 1] === "[" &&
    i !== data.length
  ) {
    brackets.pop();
  } else if (
    data[i] === ")" &&
    brackets.length !== 0 &&
    brackets[brackets.length - 1] === "(" &&
    i !== data.length
  ) {
    brackets.pop();
  } else if (
    data[i] === "}" &&
    brackets.length !== 0 &&
    brackets[brackets.length - 1] === "{" &&
    i !== data.length
  ) {
    brackets.pop();
  } else {
    result = false;
    break;
  }
}

if (brackets.length !== 0) {
  result = false;
}
console.log(result);

/*
[[[((({{{}}})))]]] TRUE
[()]{}[()()]() TRUE
[]()()((([]))) TRUE
[]()()(((([]))) FALSE
[(]){}{[()()]()} FALSE
[()]{}{[()()]()] FALSE
( FALSE
[ FALSE
{ FALSE
) FALSE
] FALSE
} FALSE
[(]) FALSE
[]()()(((([]))) FALSE
[](){{{[]}}} TRUE
[]{}()[][][] TRUE
([()][][{}]) TRUE
[([])()({})] TRUE
*/
