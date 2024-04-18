interface CustomIterator<T> {
  getNext(): T;
  hasMore(): boolean;
}

interface Agregator {
  getIterator(): CustomIterator<string>;
}

class AlphabetIterator implements CustomIterator<string> {
  #alphabet: Alphabet;
  #offset: number;
  constructor(alphabet: Alphabet) {
    this.#alphabet = alphabet;
    this.#offset = 0;
  }
  getNext(): string {
    const item = this.#alphabet.getChar()[this.#offset];
    this.#offset++;
    return item;
  }
  hasMore(): boolean {
    return this.#offset < this.#alphabet.getChar().length;
  }
}

class ReverseAlphabetIterator implements CustomIterator<string> {
  #alphabet: Alphabet;
  #offset: number;
  constructor(alphabet: Alphabet) {
    this.#alphabet = alphabet;
    this.#offset = alphabet.getChar().length - 1;
  }
  getNext(): string {
    const item = this.#alphabet.getChar()[this.#offset];
    this.#offset--;
    return item;
  }
  hasMore(): boolean {
    return this.#offset >= 0;
  }
}

class Alphabet implements Agregator {
  #char: string[];
  constructor(char: string[]) {
    this.#char = char;
  }
  getChar(): string[] {
    return this.#char;
  }
  getIterator(): CustomIterator<string> {
    return new AlphabetIterator(this);
  }
  getReverseIterator(): CustomIterator<string> {
    return new ReverseAlphabetIterator(this);
  }
}

const concreateItems = new Alphabet([
  "a",
  "b",
  "c",
  "d",
  "e",
  "f",
  "g",
  "h",
  "i",
  "j",
  "k",
  "l",
  "m",
  "n",
  "o",
  "p",
  "q",
  "r",
  "s",
  "t",
  "u",
  "v",
  "w",
  "x",
  "y",
  "z",
]);
const iterator = concreateItems.getIterator();
const reverseIterator = concreateItems.getReverseIterator();

function clientCode(iterator: CustomIterator<string>): string {
  let str = "";
  while (iterator.hasMore()) {
    str += iterator.getNext();
  }
  return str;
}

console.log("> normal order");
console.log(clientCode(iterator));
console.log("> reverse order");
console.log(clientCode(reverseIterator));
